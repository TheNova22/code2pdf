package main

import (
	"code2pdf/producer"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	
	r.Body = http.MaxBytesReader(w, r.Body, 2 * 1024 * 1024)

	rd, _:=r.MultipartReader()
	
	for {

		handler, err:=rd.NextPart()
		if err == io.EOF {
			break
		} else if err != nil{
			print(err.Error())
			break
		}
		// keys := make([]string, 0, len(handler.Header))
		
		byteContainer, _:=ioutil.ReadAll(handler)
		fmt.Printf("Uploaded File: %+v\n", handler.FileName())

		byteHex := hex.EncodeToString(byteContainer)

		producer.ProduceFile(prod,sig, "file", byteHex, handler.FileName())
		handler.Close()
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
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
