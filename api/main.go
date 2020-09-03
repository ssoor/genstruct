package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

func main() {
	port := flag.String("addr", ":8989", "addr ip:port")
	flag.Parse()

	c := cors.AllowAll()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("request body read error \n%s", err)))
			return
		}

		p := &struct {
			Table string   `json:"table"`
			Tags  []string `json:"tags"`
		}{}

		err = json.Unmarshal(body, p)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("request body json Unmarshal error \n%s", err)))
			return
		}

		st, err := ConvertSQL(p.Table, p.Tags)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("generate struct error \n%s", err)))
			return
		}
		w.Write([]byte(st))
	})

	http.Handle("/", http.FileServer(http.Dir("./dist")))
	http.Handle("/genapi/struct/gen", c.Handler(handler))

	fmt.Printf("server: http://%v\n", *port)
	err := http.ListenAndServe(*port, nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
