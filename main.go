// gopkg/main.go
package main

import (
    "flag"
    "log"
    "net"
    "os"
)

// creates a logger 
var (
    logError *log.Logger
)

// checks for FATAL ERROR on startup
func CheckError(err error) {
    if err != nil {
        logError.Fatalf("Fatal Error: %s", err.Error())
    }
}

// the -logfile flag creates a log file
func Init(logfile string) {
    
    // open a log file, if you can
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Fatal Error: ", err)
    }
    
    logError = log.New(file, "INFO: ", log.Ldate|log.Ltime)
}

func main() {
    // process flags passed from service startup
    port      := flag.String("port", "8080", "Port to bind gopkg to")
    logfile   := flag.String("logfile", "/tmp/gopkg.log", "Log file path and name")
    flag.Parse()

    // begin logging
    Init(*logfile)

    // binds to port
    tcpAddr, err := net.ResolveTCPAddr("tcp4", ":" + *port)
    CheckError(err)

    // binds to all addresses
    listener, err := net.ListenTCP("tcp", tcpAddr)
    CheckError(err)

    // starts a shared struct
    pkgindexer := &PkgIndex{}
    
    for {
        cx, err := listener.Accept()
        if err != nil {
            continue
        }
        // concurrency
        go PkgHandler(cx, pkgindexer)
    }
}

