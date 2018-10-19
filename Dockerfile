FROM golang:1.11-alpine as builder
RUN apk update; apk upgrade
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep
WORKDIR /go/src/github.com/phaicom/maqe-template-styling
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go


FROM alpine:3.8
WORKDIR /go/src
COPY --from=builder /go/src/github.com/phaicom/maqe-template-styling/ ./
CMD ["./main"]