FROM golang:1.18.1-alpine3.15 as builder

ENV GO111MODULE=on

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create mmonitoring user
ENV USER=mmonitoring
ENV UID=10001

RUN adduser \
	--disabled-password \
	--gecos "" \
	--home "/nonexistent" \
	--shell "/sbin/nologin" \
	--no-create-home \
	--uid "${UID}" \
	"${USER}"

WORKDIR $GOPATH/mmonitoring
COPY . .
RUN go mod vendor
RUN ls
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/mmonitoring -mod vendor main.go

# Container
FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/mmonitoring /go/bin/mmonitoring

USER mmonitoring:mmonitoring
ENTRYPOINT ["/go/bin/mmonitoring"]
EXPOSE 8080