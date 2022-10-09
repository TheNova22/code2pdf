package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"goServer/producer"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

const(
	flaskUrl string = "http://flask:5000/makePdf"

)

type pdfRequest struct {
    UrlLen int `json:"urlLen"`
	FileLen int `json:"fileLen"`
    Urls []string `json:"urls"`
	Files []string `json:"files"`
	Usn string `json:"usn"`
}

var prod, sig = producer.StartProducer()

type pdfResponse struct{
	Pdf string `json:"pdf"`
}

type urlInput struct{
	Urls []string `json:"urls"`
}

var templates = template.Must(template.ParseFiles("ft.html"))

// Display the named template
func display(w http.ResponseWriter, page string, data interface{}) {
	templates.ExecuteTemplate(w, page+".html", data)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
    // fmt.Print(r.FormValue("urls"))
    // nr := r.Clone(r.Context())
    r.ParseMultipartForm(20 << 20)

	// URL messaging
	urls := strings.Split(r.MultipartForm.Value["urls"][0], ",")
	urlSlice := make([]string, len(urls))
	for i, _ := range urls{
		urls[i] = strings.Trim(urls[i], " ")
		urlSlice[i] = urls[i]
		map1 := map[string]string{
			"url": urls[i],
			"usn": "u1",
		}
		jsonStr,_ := json.Marshal(map1)
		defer producer.ProduceMsg(prod,sig, "url", string(jsonStr), "u1")
	}
    // fmt.Print("Yo : " + r.MultipartForm.Value["urls"][0] + "\n")


	// File messaging
    files := r.MultipartForm.File["myFile"]
    fileSlice := make([]string, len(files))
    for i, _ := range files {
		fileSlice[i] = files[i].Filename
        file, _ := files[i].Open()
        io.Copy(&buf, file)
		err := ioutil.WriteFile("/volume/" + files[i].Filename, buf.Bytes(), 0644)
		map1 := map[string]string{
			"file": files[i].Filename,
			"usn": "u1",
		}
		jsonStr,_ := json.Marshal(map1)
		defer producer.ProduceMsg(prod,sig, "file", string(jsonStr), "u1")
		if err != nil {
			panic(err)
		}
		buf.Reset()
    }

	pdfReq := pdfRequest{UrlLen: len(urls), FileLen: len(files), Urls: urlSlice, Files: fileSlice, Usn: "1MS19IS051"}

	jsonStr,_ := json.Marshal(pdfReq)

	responseBody := bytes.NewBuffer(jsonStr)

   	resp, err := http.Post(flaskUrl, "application/json", responseBody)


	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	var p pdfResponse
	x := json.NewDecoder(resp.Body).Decode(&p)

	

	// log.Printf(string(data))
	if x == nil{
		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(string(p.Pdf)))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r,  "/volume/" + string(p.Pdf))
	} else{
		fmt.Printf("Error Occured")
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "ft", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/",uploadHandler)
	http.ListenAndServe(":8090", nil)
}
