package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var (
	TRACE = log.New(os.Stdout, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// handle post event
func MandrillPost(w http.ResponseWriter, req *http.Request) {
	TRACE.Println("in MandrillPost")
	TRACE.Println("Req Method: " + req.Method)
	w.Write([]byte("OK")) // send 200 response

	// decode JSON
	body, err := ioutil.ReadAll(req.Body)
	TRACE.Println("Req Body: " + string(body))
	if err != nil {
		ERROR.Printf("error in ReadAll: %s", err)
	}

	// Every Mandrill webhook uses the same general data format, regardless of the event type.
    // The webhook request is a standard POST request with a single parameter
    // Currently this parameter is mandrill_events. So we get slice [16:] of body.
	unescaped, unescapedErr := url.QueryUnescape(string(body)[16:])
	if unescapedErr != nil {
		ERROR.Printf("error in QueryUnescape %s", unescapedErr)
	}
	TRACE.Println(string(unescaped))

	var payload interface{}
	err = json.Unmarshal([]byte(unescaped), &payload)
	TRACE.Println(payload)

	// Mandrill Web Hook
	var msg MandrillWebHook
	err = json.Unmarshal([]byte(unescaped), &msg)
	if err != nil {
		ERROR.Printf("json.Unmarshal error: %s", err)
	}
	TRACE.Println(msg)
}

func main2() {
	http.HandleFunc("/", MandrillPost)
	TRACE.Println("calling ListenAndServe:8118")
	err := http.ListenAndServe(":8118", nil)
	if err != nil {
		ERROR.Printf("ListenAndServe: %s", err)
	}
}
