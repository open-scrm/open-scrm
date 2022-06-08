FROM hub.mrj.com:30080/base/golang:1.67.7-alpine.314-build as builder
COPY . /app
WORKDIR /app
ENV GOPROXY https://goproxy.cn
RUN go build -o app cmd/root.go

FROM hub.mrj.com:30080/base/alpine
COPY --from=builder /app/app /app

CMD ["/app"]
