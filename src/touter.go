//-*- mode: go -*-

package main

import (
    "fmt"
    // "net"
    "os"
    // "time"
    "flag"
    "log"
)

import "code.google.com/p/gcfg"

var server string
var port int
var config string
var repo_root string
var depth int
var profile string

type profiles struct {
    Profile map[string]*struct {
        Exclude []string
        Description string
    }
}

func init(){
    flag.StringVar(&server, "server", "localhost", "Server to send to")
    flag.StringVar(&server, "s", "localhost", "Server to send to (Shorthand)")

    flag.IntVar(&port, "port", 2002, "Port to connect on")
    flag.IntVar(&port, "p", 2002, "Port to connect on (Shorthand)")

    flag.StringVar(&config, "file", "/etc/touter.ini", "Config file for touter")
    flag.StringVar(&config, "f", "/etc/touter.ini", "Config file for touter (Shorthand)")

    flag.StringVar(&repo_root, "repo_root", "/path", "Repo root to start from")
    flag.StringVar(&repo_root, "r", "/path", "Repo root to start from (Shorthand)")

    flag.IntVar(&depth, "depth", 2, "Depth to walk through repo_root")
    flag.IntVar(&depth, "d", 2, "Depth to walk through repo_root (Shorthand)")

    flag.StringVar(&profile, "profile", "rails", "Profile from config file to use")
    flag.StringVar(&profile, "pr", "rails", "Profile from config file to use (Shorthand)")
}

func show_init_settings() string{
    return fmt.Sprintf("%s running with:\n\tServer: %s\n\tPort: %d\n\tConfig: %s\n\tRepo root: %s\n\tDepth: %d\n\tProfile: %s",
        os.Args[0],
        server,
        port,
        config,
        repo_root,
        depth,
        profile)
}

func load_profiles() []string{
    var p profiles
    err := gcfg.ReadFileInto(&p, config)
    if err != nil {
        log.Fatalf("Failed to parse gcfg data: %s", err)
    }
    fmt.Printf("\n%s :\n\t%s\nExcluding: %s\n",
        profile,
        p.Profile[profile].Description,
        p.Profile[profile].Exclude)
    return p.Profile[profile].Exclude
}

func main(){
    flag.Parse()
    fmt.Println( show_init_settings() )
    load_profiles()

    // Loop through repos
}
