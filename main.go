package main

import (
	"bytes"
	"code2pdf/producer"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

var prod, sig = producer.StartProducer()

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

	for i, _ := range urls{
		urls[i] = strings.Trim(urls[i], " ")
		map1 := map[string]string{
			"url": urls[i],
			"usn": "u1",
		}
		jsonStr,_ := json.Marshal(map1)
		producer.ProduceMsg(prod,sig, "url", string(jsonStr), "u1")
	}
    // fmt.Print("Yo : " + r.MultipartForm.Value["urls"][0] + "\n")


	// File messaging
    files := r.MultipartForm.File["myFile"]
    
    for i, _ := range files {
        file, _ := files[i].Open()
        io.Copy(&buf, file)
		err := ioutil.WriteFile("volume/" + files[i].Filename, buf.Bytes(), 0644)
		map1 := map[string]string{
			"file": files[i].Filename,
			"usn": "u1",
		}
		jsonStr,_ := json.Marshal(map1)
		producer.ProduceMsg(prod,sig, "file", string(jsonStr), "u1")
		if err != nil {
			panic(err)
		}
		buf.Reset()
    }
	// Download File
    w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("volume/" + files[0].Filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeFile(w, r, "volume/" + files[0].Filename)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		display(w, "ft", nil)
	case "POST":
		uploadFile(w, r)
	}
}

func url(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	var u urlInput
	err := json.NewDecoder(req.Body).Decode(&u)

	

	// log.Printf(string(data))
	if err == nil{
		for i := 0; i < len(u.Urls); i++{
			producer.ProduceUrl(prod,sig, "url", u.Urls[i])
		}
		fmt.Fprintf(w, "Message Added")
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

	// Flags
	// consumerFlagValue := flag.Bool("c", false, "    Use this flag to start a Kafka Consumer")
	// producerFlagValue := flag.Bool("p", false, "    Use this flag to start a Kafka Producer")
	// stringFlagValue := flag.String("a", "", "    Use this flag with either \"consumer\" or \"producer\"")

	// Flag Processing
	// flag.Parse()

	

	

	// Decision Time
	// if *producerFlagValue {
	// 	producer.StartProducer()
	// } else if *consumerFlagValue  {
	// 	consumer.StartConsumer()
	// } else if *stringFlagValue == "consumer" {
	// 	consumer.StartConsumer()
	// } else if *stringFlagValue == "producer" {
	// 	producer.StartProducer()
	// } else {
	// 	fmt.Print("Usage: \n -c     Use this flag to start a Kafka Consumer\n -p     Use this flag to start a Kafka Producer\n -a     Use this flag with either \"consumer\" or \"producer\"\n")
	// }
	http.HandleFunc("/url", url)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/upload",uploadHandler)
	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	http.ListenAndServe(":8090", nil)
}
