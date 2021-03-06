package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strings"
    "strconv"
    "database/sql"
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

func main() {
    db, err = sql.Open(dbDriverName, dbName)
    if checkErr(err) {
        fmt.Println(err)
    }
    err := createTable(db)
    if checkErr(err) {
        fmt.Println(err)
    }
    http.HandleFunc("/", index)
    fmt.Println("xsspot running on 80")
    if err := http.ListenAndServe("0.0.0.0:80", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}