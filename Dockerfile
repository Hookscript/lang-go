FROM golang:1.4.2-wheezy

ADD main.go /gopath/

ADD compile /bin/
ADD run /bin/
