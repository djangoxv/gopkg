// gopkg/handler_test.go
package main

import "testing"

func TestParseRequest(t *testing.T) {
    pkgindex := &PkgIndex{}
    testBadRequest := "I|fail\n"
    if ParseRequest(testBadRequest, pkgindex) != ERROR {
        t.Error("ReturnCode incorrect for a malformed REQUEST")
    }
}
