##  Installation of gopkg Ubuntu 14.04

# Method 1: run as daemon on source 

1. Install golang

# update certificates used for downloading 
  sudo apt-get install apt-transport-https ca-certificates

# downloads and unpacks golang 1.6.2 to /usr/local/go
  sudo curl -s https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz | sudo tar -v -C /usr/local -xz  

# sets up local home directory to run go in
  mkdir go
  echo  "export PATH=/usr/local/go/bin/:$PATH" >> ~/.profile
  echo  "export GOPATH=/home/ubuntu/work" >> ~/.profile
  source ~/.profile

2. Run gopkg for first time

# download and install gopkg latest from github
  go get github.com/djangoxv/gopkg
  go install github.com/djangoxv/gopkg

# Run in daemon mode
# default flags are "-logfile /tmp/gopkg.log"  and "-port 8080" 
  /home/ubuntu/go/bin/gopkg &

# Method 2: Run docker, See https://docs.docker.com/engine/installation/linux/ubuntulinux/

1. Install and setup docker daemon for user

#  become root temporarily to create apt-get source file
   sudo su -
   echo "deb https://apt.dockerproject.org/repo ubuntu-trusty main" > /etc/apt/sources/list.d/docker.list
   exit 

# Add apt key for dockerproject apt repository
   sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D

# Install latest docker-engine
   sudo apt-get update
   sudo apt-get install docker-engine   

# Add current user (e.g. "ubuntu") to docker group
   sudo usermod -aG docker ubuntu

# Log OUT & Log IN
   in practice

4. Build and gopkg container
   docker build -f Dockerfile.ubuntu -t gopkg github.com/djangoxv/gopkg

5. Start gopkg container
   docker run -d -p 8080:8080 gopkg

## Testing

# download DigitalOcean test provided
  curl http://52.40.15.68/artifactory/gopkg/test/do-package-tree_linux -o /tmp/do-package-tree_linux -s;chmod +x /tmp/do-package-tree_linux

# run test (with concurrency 100)
   /tmp/do-package-tree_linux   
