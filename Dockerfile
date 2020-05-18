FROM alpine

# CGO_ENABLED=0 go build -v .
ADD go-sitecheck /usr/bin/
RUN chmod +x /usr/bin/go-sitecheck 