FROM golang:alpine
WORKDIR /go/src/power/ch1/httpserver
COPY . ./
RUN go build -o httpserver .
CMD ["/go/bin/httpserver"]

#docker build . -t httpserver
#docker push httpserver // 推送到远程镜像服务器
#docker run -p 8080:8080 httpserver
