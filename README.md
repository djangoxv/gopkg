## gopkg: A Package Indexer ##

  Written in the Go Programming Language as a sample project for DigitalOcean



### To run test provided by DigitalOcean in Linux ###
* Download latest binary & make executable

  curl http://52.40.15.68/artifactory/gopkg/aws-build/gopkg -o ./gopkg  -s;chmod +x gopkg


* Start gopkg in background mode (-logfile /tmp/gopkg/log -port 8080)

   ./gopkg &


*  Download DigitalOcean test provided 

   curl http://52.40.15.68/artifactory/gopkg/test/do-package-tree_linux -o /tmp/do-package-tree_linux -s;chmod +x /tmp/do-package-tree_linux


* Run test (with concurrency 100)

   /tmp/do-package-tree_linux -concurrency 100
  

 
### Binary Distribution Notes ###

  See https://golang.org/doc/install or [gopkg/INSTRUCTIONS.md](https://github.com/djangoxv/gopkg/blob/master/INSTRUCTIONS.md) for further information about building and running from go or docker 
