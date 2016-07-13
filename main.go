// gppkg/main.go
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

// checks for FATAL ERROR
func CheckError(err error) {
    if err != nil {
        logError.Fatalf("Fatal Error: %s", err.Error())
    }
}

// -logfile creates a log file
func Init(logfile string, debug bool) {
    
    // open a log file, if you can
    file, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Fatal Error: ", err)
    }
    
    if debug {
        logError = log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
    } else {
        logError = log.New(file, "INFO: ", log.Ldate|log.Ltime)
    }

    

}

func main() {
    // process flags passed from service startup
    // count     := flag.Int("count", 100, "Maximum number of concurrent requests")
    port      := flag.String("port", "8080", "Port to bind gopkg to")
    logfile   := flag.String("logfile", "/tmp/gopkg.log", "Log file path and name")
    debug     := flag.Bool("debug", false, "Log verbosity level for debugging gopkg")
    flag.Parse()

    // begin logging
    Init(*logfile, *debug)

    // binds to port
    tcpAddr, err := net.ResolveTCPAddr("tcp4", ":" + *port)
    CheckError(err)

    // binds to all addresses
    listener, err := net.ListenTCP("tcp", tcpAddr)
    CheckError(err)

    // starts a concurrent handler
    pkgindexer := &PkgIndex{}
    
    for {
        cx, err := listener.Accept()
        if err != nil {
            continue
        }
        go PkgHandler(cx, pkgindexer)
    }
}

