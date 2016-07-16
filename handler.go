// gopkg/handler.go
package main

import (
    "bufio"
    "net"
    "regexp"
    "time"
)

type ReturnCode string
// set of possible reponse strings
const (
    OK = "OK"
    FAIL = "FAIL"
    ERROR = "ERROR"
)

func PkgEncoder(pkgname string, pkgdeps []string) *Package {
    pkg := PkgCreate(pkgname, pkkgdeps)
    return pkg
}

// parses for a valid string format in received message
func ParseRequest(msg string, pkgindexer *PkgIndex) ReturnCode {
    // not great performance, distribute information > arduous regex
    valid := regexp.MustCompile(`(^INDEX|REMOVE|QUERY)\|[A-Za-z0-9-_\+]+\|(\s|[A-Za-z0-9-_+,]+)`)

    if valid.MatchString(msg) {
        r := regexp.MustCompile(`\|`).Split(msg, -1)
        action, pkgname, pkgdeps := r[0], r[1], r[2]
        deplist := regexp.MustCompile(`,`).Split(pkgdeps, -1)
        pkg := CreatePackage(pkgname, deplist)
        switch {
            case action == "INDEX":
                return PkgInvoke(pkgname, deplist)
            case action == "QUERY":
                return PkgQuery(pkgname)
            case action == "REMOVE":
                return pkgindexer.PkgRemove(pkgname)
        }
    }
    return ERROR
}

// handles the connection timeout and reading buffer
func PkgHandler(cx net.Conn, mq chan) {
    cx.SetReadDeadline(time.Now().Add(time.Second * 20)) // 20 second timeout
    defer cx.Close() // close connection on exit

    for {
        // will listen for message to process (\n)
        request, err := bufio.NewReader(cx).ReadString('\n')
        if err != nil {
            logError.Println(err)
            break
        }

        responseString <- mq)
        cx.Write([]byte(responseString + "\n"))

    }
}
