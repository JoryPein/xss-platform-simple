package main

import (
	"flag"
	"xss/client"
	"xss/server"
)

var (
	mode   string
	port   int
	ip     string
	pass   string
	sslcrt string
	sslkey string
)

func main() {
	flag.StringVar(&mode, "mode", "", "run mode,server or client")
	flag.IntVar(&port, "port", 9999, "server port")
	flag.StringVar(&ip, "ip", "", "Listen ip")
	flag.StringVar(&mode, "pass", "", "Connect password")
	flag.StringVar(&sslcrt, "sslcrt", "./server.crt", "sslcrt")
	flag.StringVar(&sslkey, "sslkey", "./server.key", "sslkey")
	flag.Parse()
	if mode == "server" {
		server.Run(port, ip, pass, sslcrt, sslkey)
	} else if mode == "client" {
		client.Run(port, ip, pass)
	}

}
