//-*- mode: go -*-

package main

import (
    "fmt"
    // "net"
    "os"
    // "time"
    "flag"
)

var server string
var port int
var config string

func init(){
    flag.StringVar(&server, "server", "localhost", "Server to send to")
    flag.StringVar(&server, "s", "localhost", "Server to send to")

    flag.IntVar(&port, "port", 2002, "Port to connect on")
    flag.IntVar(&port, "p", 2002, "Port to connect on")

    flag.StringVar(&config, "file", "/etc/touter.ini", "Config file for touter")
    flag.StringVar(&config, "f", "/etc/touter.ini", "Config file for touter")
}

func show_config(){
    fmt.Println(os.Args[0], " running as per:")
    fmt.Println("Server: ", server)
    fmt.Println("Port: ", port)
    fmt.Println("Path: ", config)
}

func main(){
    flag.Parse()
    show_config()
}
