FROM golang:1.5.1-wheezy

ADD main.go /gopath/

ADD compile /bin/
ADD run /bin/
