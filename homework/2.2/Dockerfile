FROM golang AS build
COPY  . .
ENV CGO_ENABLED=0 
ENV GO111MODULE=on 
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build  -o /httpsrv1 main.go

FROM alpine
COPY --from=build /httpsrv1 /bin/
EXPOSE 80
ENTRYPOINT ["/bin/httpsrv1"]
