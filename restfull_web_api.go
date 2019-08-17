package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	addr = flag.String("addr", ":8000", "http service address")
	data map[string]string
)

func main() {
	flag.Parse()
	data = map[string]string{}
	r := httprouter.New()
	r.GET("/entry/:key", show)
	r.GET("/list", show)
	r.PUT("/entry/:key/:value", update)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func show(w http.ResponseWriter, r *http.Request, p httprouter.params) {
	k := p.ByName("key")
	if k == "" {
		fmt.Fprintf(w, "Read list: %v", data)
		return
	}
	fmt.Fprintf(w, "Read entry: data[%s] = %s", k, data[k])
}
func update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	k := p.ByName("key")
	v := p.ByName("value")

	data[k] = v

	fmt.Fprintf(w, "Updated: data[%s] = %s", k, data[k])
}
