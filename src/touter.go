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
    flag.StringVar(&server, "s", "localhost", "Server to send to (Shorthand)")

    flag.IntVar(&port, "port", 2002, "Port to connect on")
    flag.IntVar(&port, "p", 2002, "Port to connect on (Shorthand)")

    flag.StringVar(&config, "file", "/etc/touter.ini", "Config file for touter")
    flag.StringVar(&config, "f", "/etc/touter.ini", "Config file for touter (Shorthand)")
}

func show_config() string{
    return fmt.Sprintf("%s running with:\n\tServer: %s\n\tPort: %d\n\tConfig: %s",
        os.Args[0],
        server,
        port,
        config)
}

func main(){
    flag.Parse()
    fmt.Println( show_config() )
}
