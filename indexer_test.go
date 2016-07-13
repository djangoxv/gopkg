// gopkg/indexer_test.go
package main

import "testing"

// tests that a package addition and add of same package returns OK
func TestPkgIndex(t *testing.T) {

        alphaPkg := PkgCreate("pkgtest")
        betaPkg := PkgCreate("pkgtest")
        if alphaPkg != betaPkg {
                t.Error("two packages created differently")
        }
}
