FROM golang:1.17.1-alpine3.14
WORKDIR /home/mcstk/kolo-assignment/
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-starter .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /home/mcstk/kolo-assignment/go-starter .
ENV APP_ENV=dev
EXPOSE 80
CMD ["./go-starter"]
