FROM golang:1.11 AS go-builder

WORKDIR /go/src/github.com/Evertras/fbr

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	  go build -a -tags netgo \
	  -ldflags '-w -extldflags "-static"' \
	  -o /fbr \
	  ./cmd/fbr/main.go

FROM scratch

COPY --from=go-builder /fbr /fbr

ENTRYPOINT ["/fbr"]
