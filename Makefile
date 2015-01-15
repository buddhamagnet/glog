default: install

deploy:
				git push origin master && git push heroku master

install:get
				 GOPATH=~/glog go install

fmt:
				 go fmt *.go

get:
				 GOPATH=~/glog go get