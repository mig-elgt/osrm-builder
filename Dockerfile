FROM golang:1.19 AS builder
WORKDIR /go/src/github.com/mig-elgt/osrm-builder
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o /bin/app cmd/osrm-builder/main.go

FROM osrm/osrm-backend:v5.22.0
MAINTAINER Miguel Angel Galicia
RUN apt-get update && apt-get install -y curl

COPY --from=builder /bin/app /usr/local/bin/app
RUN chmod +x /usr/local/bin/app
CMD ["app"]
