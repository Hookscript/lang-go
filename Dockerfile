FROM golang:1.6.0-wheezy

ADD main.go /gopath/

ADD compile /bin/
ADD run /bin/
