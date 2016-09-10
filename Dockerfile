FROM golang:1.7.1
RUN go get -u github.com/kataras/iris/iris
RUN go get -u github.com/iris-contrib/template/html
EXPOSE 8080