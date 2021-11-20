package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func removeBlackets(s string) string {
	s = strings.ReplaceAll(s, ":[\"", ":\"")
	s = strings.ReplaceAll(s, "\"]", "\"")
	return s
}

var db *sql.DB
var err error

func index(w http.ResponseWriter, r *http.Request) {
	ipPort := strings.Split(r.RemoteAddr, ":")
	ip := ipPort[0]
	port, _ := strconv.Atoi(ipPort[1])

	url := r.URL.String()

	headerJson, _ := json.Marshal(r.Header)
	header := removeBlackets(string(headerJson))

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	body := string(bodyBytes)

	fmt.Println("----------------------")
	fmt.Println("Ip: ", ip)
	fmt.Println("Port: ", port)
	fmt.Println("URL: ", url)
	fmt.Println("Header: ", header)
	fmt.Println("Body: ", body)

	err := insertData(db, cross{ip, port, url, header, body})
	if checkErr(err) {
		fmt.Println(err)
	}
}

func Run(port int, ip, pass, sslcrt, sslkey string) {
	db, err = sql.Open(dbDriverName, dbName)
	if checkErr(err) {
		fmt.Println(err)
	}
	err := createTable(db)
	if checkErr(err) {
		fmt.Println(err)
	}

	go func() {
		httpserver := http.NewServeMux()
		httpserver.HandleFunc("/", index)
		fmt.Println("xsspot running on 80")
		if err := http.ListenAndServe("0.0.0.0:80", httpserver); err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}()

	go func() {
		httpserver := http.NewServeMux()
		httpserver.HandleFunc("/", index)
		fmt.Println("xsspot running on 443")
		if sslcrt != "" && sslkey != "" {
			if err := http.ListenAndServeTLS("0.0.0.0:443", sslcrt, sslkey, httpserver); err != nil {
				log.Fatal("ListenAndServeTLS: ", err)
			}
		}
	}()
	select {}

}
