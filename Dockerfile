# Builder
FROM golang:1.19.1-alpine3.15 as builder

ENV GO111MODULE=on
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

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

WORKDIR $GOPATH/src/mmonitoring
COPY . .

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /go/bin/mmonitoring -mod vendor main.go


# Main Image
FROM debian:buster-slim

LABEL name="lighthouse" \
	maintainer="Leonardo Metzger <leonardo.metzger@outlook.com>" \
	version="1.0" 

# Install deps + add Chrome Stable + purge all the things
RUN apt-get update && apt-get install -y \
	apt-transport-https \
	ca-certificates \
	curl \
	gnupg \
	--no-install-recommends \
	&& curl -sSL https://deb.nodesource.com/setup_16.x | bash - \
	&& curl -sSL https://dl.google.com/linux/linux_signing_key.pub | apt-key add - \
	&& echo "deb https://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list \
	&& apt-get update && apt-get install -y \
	google-chrome-stable \
	nodejs \
	--no-install-recommends \
	&& apt-get purge --auto-remove -y curl gnupg \
	&& rm -rf /var/lib/apt/lists/*

ARG CACHEBUST=1
RUN npm install -g lighthouse

# Copy go application to container
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /go/bin/mmonitoring /usr/bin/mmonitoring
COPY configs /var/mmonitoring/configs

RUN echo 'kernel.unprivileged_userns_clone=1' > /etc/sysctl.d/userns.conf

# Add Chrome as a user
RUN groupadd -r chrome && useradd -r -g chrome -G audio,video chrome \
	&& mkdir -p /home/chrome/reports && chown -R chrome:chrome /home/chrome \
	&& mkdir -p /tmp/reports

VOLUME /home/chrome/reports
WORKDIR /home/chrome/reports
USER chrome

ENTRYPOINT [ "/usr/bin/mmonitoring" ]
