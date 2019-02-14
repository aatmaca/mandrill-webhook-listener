package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Mandrill feed implementation

// NewMandrillFeed creates Mandrill Feed
func NewMandrillFeed(mandrillAddr string) Feed {
	return &mandrillFeed{
		mandrillAddr: mandrillAddr,
		out:          make(chan interface{}, 0),
	}
}

type mandrillFeed struct {
	mandrillAddr string
	out          chan interface{}
}

func (m *mandrillFeed) Out() <-chan interface{} {
	return (<-chan interface{})(m.out)
}

func (m *mandrillFeed) Start() {
	go func() {
		err := http.ListenAndServe(m.mandrillAddr, m)
		if err != nil {
			log.Printf("ListenAndServe: %s", err)
		}
	}()
}

func (m *mandrillFeed) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	//log.Printf("Req: %s", string(r.URL.Path)+"...."+r.URL.RequestURI())

	data, err := ioutil.ReadAll(r.Body)

	if len(data) < 17 {
		return
	}

	if err != nil {
		log.Printf("error reading request body: %s", err)
		return
	}

	// TO DO should be removed after test
	log.Printf("Body received: %s", string(data))

	// Every Mandrill webhook uses the same general data format, regardless of the event type.
	// The webhook request is a standard POST request with a single parameter
	// Currently this parameter is mandrill_events. So we get slice [16:] of body.
	unescaped, unescapedErr := url.QueryUnescape(string(data)[16:])
	if unescapedErr != nil {
		log.Printf("error in QueryUnescape: %s", unescapedErr)
	}
	//TRACE.Println(string(unescaped))

	var msg []interface{}
	err = json.Unmarshal([]byte(unescaped), &msg)
	if err != nil {
		log.Printf("error unmarshalling into json: %s", err)
	}

	for _, payload := range msg {
		m.out <- payload
	}
}
