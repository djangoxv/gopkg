// gopkg/handler.go
package main

import (
    "bufio"
    "net"
    "regexp"
    "time"
)

type ReturnCode string
const (
        OK = "OK"
        FAIL = "FAIL"
        ERROR = "ERROR"
)


func ParseRequest(msg string, pkgindexer *PkgIndex) ReturnCode {
    
    valid := regexp.MustCompile(`(^INDEX|REMOVE|QUERY)\|[A-Za-z0-9-_\+]+\|(\s|[A-Za-z0-9-_+,]+)`)
    if valid.MatchString(msg) {
        r := regexp.MustCompile(`\|`).Split(msg, -1)
        action, pkgname, pkgdeps := r[0], r[1], r[2]

        if action == "INDEX" {
            deplist := regexp.MustCompile(`,`).Split(pkgdeps, -1)
            if pkgindexer.PkgInvoke(pkgname, deplist) {
                return OK
            }
            return FAIL
            
        } else if action == "QUERY" {
            if pkgindexer.PkgQuery(pkgname) {
                return OK
            }
            return FAIL

        } else if action == "REMOVE" {
            if pkgindexer.PkgRemove(pkgname) {
                return OK
            }
            return FAIL
        } else {
            return ERROR
        }
    } else {
        return ERROR
    }
}

func PkgHandler(cx net.Conn, pkgindexer *PkgIndex) {
    cx.SetReadDeadline(time.Now().Add(time.Second * 20)) // 20 second timeout
    defer cx.Close() // close connection on exit

    for {
        // will listen for message to process ending in newline (\n)
        request, err := bufio.NewReader(cx).ReadString('\n')
        if err != nil {
            logError.Println(err)
            break
        }

        responseString := ParseRequest(request, pkgindexer)
        //logError.Println("RESPONSE: ", responseString)
        cx.Write([]byte(responseString + "\n"))

    }
}
