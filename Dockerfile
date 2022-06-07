FROM hub.mrj.com:30080/base/golang:1.67.7-alpine.314-build
COPY ./app /app

CMD ["/app"]
