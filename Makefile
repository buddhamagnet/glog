default: install

install:get
				 GOPATH=~/glog go install

fmt:
				 go fmt *.go

get:
				 GOPATH=~/glog go get