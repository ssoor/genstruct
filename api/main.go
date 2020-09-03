package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

// 通过执行 go generate ./... 将静态资源带包到程序里
//go:generate gex go-bindata -fs -o assets_gen.go -prefix ../web/dist ../web/dist

func main() {
	addr := flag.String("addr", ":8989", "addr ip:port")
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

	http.Handle("/", http.FileServer(AssetFile()))
	http.Handle("/genapi/struct/gen", c.Handler(handler))

	ln, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatal("Listen", err)
	}
	fmt.Printf("server: http://%v\n", ln.Addr())

	err = http.Serve(ln, nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
