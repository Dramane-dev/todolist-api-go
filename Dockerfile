FROM golang:1.18 AS builder
WORKDIR /simplytodo-go
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-api .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /simplytodo-go/go-api ./
CMD ["./go-api"]