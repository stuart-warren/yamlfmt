FROM golang:1.13
WORKDIR /go/src/github.com/stuart-warren/yamlfmt/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o yamlfmt ./cmd/yamlfmt

FROM alpine:3.16  
RUN apk --no-cache add diffutils
WORKDIR /tmp
COPY --from=0 /go/src/github.com/stuart-warren/yamlfmt/yamlfmt /usr/local/bin/yamlfmt
ENTRYPOINT ["/usr/local/bin/yamlfmt"] 