package main

import (
	"flag"
	"xss/client"
	"xss/server"
)

var (
	mode   string
	port   string
	ip     string
	pass   string
	sslcrt string
	sslkey string
)

func main() {
	flag.StringVar(&mode, "mode", "", "run mode,server or client")
	flag.StringVar(&port, "port", "9999", "server port")
	flag.StringVar(&ip, "ip", "", "Listen ip")
	flag.StringVar(&mode, "pass", "", "Connect password")
	flag.StringVar(&mode, "sslcrt", "", "sslcrt")
	flag.StringVar(&mode, "sslkey", "", "sslkey")
	flag.Parse()
	if mode == "server" {
		server.Run()
	} else if mode == "client" {
		client.Run()
	}

}
