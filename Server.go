package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type TestResponse struct {
	Header  map[string]string
	Content string
}

func handler(w http.ResponseWriter, request *http.Request) {
	fmt.Print("\n\n==============================RequestStart==============================\n")
	fmt.Println("Method:" + request.Method)
	fmt.Println("URL:" + request.URL.String())
	fmt.Println("Host:" + request.Host)
	fmt.Println("RemoteAddr:" + request.RemoteAddr)

	fmt.Println("==============================Header==============================")
	for k, vs := range request.Header {
		for _, v := range vs {
			fmt.Printf("%s=%s\n", k, v)
		}
	}

	fmt.Println("==============================Body==============================")
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", b)

	fmt.Print("\n\n==============================RequestEnd==============================\n\n")

	bytes, err := ioutil.ReadFile("response.json")

	var testResponse TestResponse
	json.Unmarshal(bytes, &testResponse)

	for k, v := range testResponse.Header {
		w.Header().Set(k, v)
	}

	w.WriteHeader(200)

	fmt.Fprint(w, testResponse.Content)
}

func main() {
	port := "9993"
	sslPort := "9994"
	args := os.Args
	argsLen := len(args)
	if argsLen > 1 {
		port = args[1]
	}
	if argsLen > 2 {
		sslPort = args[2]
	}
	fmt.Println("Server User Port:" + port + ",sslPort" + sslPort)
	http.HandleFunc("/", handler)
	go http.ListenAndServe(":"+port, nil)
	http.ListenAndServeTLS(":"+sslPort, "server.crt", "server.key", nil)
}
