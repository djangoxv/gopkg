// gopkg/indexer.go
package main

// a Package consists of a name and dependencies
type Package struct {
    PkgName   string
    PkgDeps   []*Package
}

// AddDep adds dependencies to a Package
func (pkg *Package) AddDep(to *Package) {
    pkg.PkgDeps = append(pkg.PkgDeps, to)
}

// the PkgIndex is the total collection of Packages
// a new Package's dependencies must be indexed *before* a new package may be indexed
type PkgIndex struct {
    Packages []*Package
}

func (pkgindex *PkgIndex) PkgList() []string {
   pkglist := []string{}
   for _, p := range pkgindex.Packages {
       pkglist = append(pkglist, p.PkgName)
   }
   return pkglist
}

// PkgCreate returns an empty package
func PkgCreate(pkgname string) *Package {
    return &Package{
        PkgName: pkgname,
        PkgDeps: make([]*Package, 0),
    }
}

// PkgInvoke finds or creates a new package  
func (pkgindex *PkgIndex) PkgInvoke(pkgname string, pkgdeps []string) bool {
    var pkg *Package
    for _, p := range pkgindex.Packages {
        if p.PkgName == pkgname {
            return true
        }
    }

    if pkg == nil {
        // create empty package
        pkg = PkgCreate(pkgname)
        // add deps to package
        for _, d := range pkgdeps {
            for _, p := range pkgindex.Packages {
                if p.PkgName == d {
                    pkg.AddDep(p) 
                }
            }
        }       
        pkgindex.Packages = append(pkgindex.Packages, pkg)
        for _, p := range pkgindex.Packages {
            if p.PkgName == pkgname {
               logError.Println("INDEXED ", p.PkgName)
            }
        }
    }

    return true
}


// PkgQuery returns boolean for whether the Package shows up
func (pkgindex *PkgIndex) PkgQuery(pkgname string) bool {
    logError.Println("QUERIED ", pkgname)
    for _, p := range pkgindex.Packages {
        logError.Println("QUERIED ", p.PkgName)
        if p.PkgName == pkgname {
            return true
        }
    }
    return false
}

func (pkgindex *PkgIndex) PkgRow(pkgname string) (int, bool) {
    var i int
    for i := range pkgindex.Packages {
        if pkgindex.Packages[i].PkgName == pkgname {
            return i, true
        }
    }
    return i, false
}

func (pkgindex *PkgIndex) DelPkg(i int) {
    pkgindex.Packages = append(pkgindex.Packages[:i], pkgindex.Packages[i+1:]...)
    return
}

// PkgRemove removes a package from the list, returns boolean if removal is depended upon
func (pkgindex *PkgIndex) PkgRemove(pkgname string) bool {
    var exists bool
    for _,p := range pkgindex.Packages {

        for _, q := range p.PkgDeps {
            if q.PkgName == pkgname {
                return false
            }
        }
    }
    
    
    logError.Println("Removed ", pkgname)
    j, exists := pkgindex.PkgRow(pkgname)
    if exists {
        pkgindex.DelPkg(j)
    }
    return true    
}


