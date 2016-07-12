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

func ParseRequest(msg string) ReturnCode {
    valid := regexp.MustCompile(`(^INDEX|REMOVE|QUERY)\|[A-Za-z0-9-_\+]+\|(\s|[A-Za-z0-9-_+,]+)`)

    logError.Printf("MATCHED= %s MESSAGE= %s ", valid.MatchString(msg), msg )

    if valid.MatchString(msg) {
        r := regexp.MustCompile(`\|`).Split(msg, -1)
        action, pkgname, pkgdeps := r[0], r[1], r[3]
        // logError.Printf("%s PKG=%s DEPS=%s", action, pkgname, pkgdeps)
        if action == "INDEX" {
            logError.Printf("%s PKG=%s DEPS=%s", action, pkgname, pkgdeps)
            return OK
        } else if action == "QUERY" {
            return OK
        } else if action == "REMOVE" {
            logError.Printf("%s PKG=%s DEPS=%s", action, pkgname, pkgdeps)
            return OK
        } else {
            return ERROR
        }
    } else {
        return ERROR
    }
}

func PackageRequestHandler(cx net.Conn) {
    cx.SetReadDeadline(time.Now().Add(time.Second * 10)) // 10 second timeout
    defer cx.Close() // close connection on exit

    for {
        // will listen for message to process ending in newline (\n)
        request, err := bufio.NewReader(cx).ReadString('\n')
        if err != nil {
            logError.Println(err)
            break
        }

        responseString := ParseRequest(request)
        logError.Println("RESPONSE: ", responseString)
        cx.Write([]byte(responseString + "\n"))

    }
}
