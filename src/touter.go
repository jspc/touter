//-*- mode: go -*-

package main

import (
    "fmt"
    "strings"
    "bytes"
    "net"
    "os"
    "os/exec"
    "os/signal"
    "syscall"
    "path/filepath"
    "flag"
    "log"
    "encoding/json"
    "time"
)

import "code.google.com/p/gcfg"

var server string
var port int
var config string
var repo_root string
var depth int
var wait int
var profile string
var excludes []string
var hostname string
var err error

type Profiles struct {
    Profile map[string]*struct {
        Exclude []string
        Description string
    }
}

type Project struct {
    Path string
    Sha string
    Branch string
}
var projects []Project

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

    flag.IntVar(&wait, "wait", 1, "Minutes to wait between updates")
    flag.IntVar(&wait, "w", 1, "Minutes to wait between updates (Shorthand)")
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

func load_profile() []string{
    var p Profiles
    err := gcfg.ReadFileInto(&p, config)
    if err != nil {
        log.Fatalf("Failed to parse gcfg data: %s", err)
    }
    log.Printf("Profile: %s: %s", profile, p.Profile[profile].Description)
    log.Printf("Excluding: %s\n", p.Profile[profile].Exclude)

    return p.Profile[profile].Exclude
}

func git_info(fp string) (string, string){
    // Create new objects over and over or we end up appending
    var sha bytes.Buffer
    var branch bytes.Buffer

    branch_command := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
    branch_command.Dir = fp
    branch_command.Stdout = &branch
    branch_err := branch_command.Run()

    sha_command := exec.Command("git", "rev-parse", "HEAD")
    sha_command.Dir = fp
    sha_command.Stdout = &sha
    sha_err := sha_command.Run()

    if sha_err != nil || branch_err != nil {
        log.Fatalf("Error(s) in %s:\n\t%s %s", fp, sha_err, branch_err)
    }

    return strings.Trim(sha.String(), "\n"), strings.Trim(branch.String(), "\n")
}

func ignored(dir_parts []string) bool{
    for _, exclude := range excludes {
        for _, dir := range dir_parts {
            if dir == exclude {
                return true
            }
        }
    }
    return false
}

func walker(fp string, fi os.FileInfo, err error) error{
    if err != nil {
        return nil
    }
    relative_path := strings.Replace(fp, repo_root, "", 1)
    split_path    := strings.Split(relative_path, "/")
    filename      := split_path[len(split_path)-1]

    if !fi.IsDir() ||
       strings.Count(relative_path, "/") > depth ||
       filename != ".git" ||
       ignored(split_path) {
        return nil
    }
    project_name := strings.Replace(fp, "/.git", "", 1)

    var p Project
    p.Path = project_name
    p.Sha, p.Branch = git_info(project_name)
    projects = append(projects, p)

    return nil
}

func send_message(msg string) bool {
    destination := fmt.Sprintf("%s:%d", server, port)
    conn, err := net.Dial("udp", destination)
    if err != nil {
        log.Fatalf("Could not send to %s: %s", destination, err)
    }
    fmt.Fprintf(conn, msg)
    return true
}

func worker(){
    projects = nil
    filepath.Walk( repo_root, walker)
    log.Println( "Projects: ", projects )

    message := make(map[string][]Project)
    message[hostname] = projects

    j, err := json.Marshal(message)
    if err != nil {
        log.Fatalf("Could not marshal projects data: %s\n", err)
    }

    if send_message(string(j)) {
        log.Println("Succesfully spaffed our repos")
    } else {
        log.Fatal("Something fucked up")
    }
}

func main(){
    flag.Parse()
    hostname, err = os.Hostname()
    if err != nil {
        log.Fatalf("Could not get hostname: %s", err)
    }

    log.Printf( "Hostname: %s\n", hostname )

    log.Println( show_init_settings() )
    excludes = load_profile()

    reload := make(chan os.Signal, 1)
    send   := make(chan os.Signal, 1)

    signal.Notify(reload, syscall.SIGHUP)
    signal.Notify(send, syscall.SIGUSR1)

    go func(){
        for sig := range reload {
            log.Printf("Received %s - reloading profile\n", sig)
            excludes = load_profile()
        }
    }()

    go func(){
        for sig := range send {
            log.Println("Received %s - sending repo data", sig)
            worker()
        }
    }()

    for {
        go worker()
        time.Sleep( time.Duration(wait) * time.Minute )
    }
}
