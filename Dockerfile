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

# Add Chrome as a user
RUN groupadd -r chrome && useradd -r -g chrome -G audio,video chrome \
	&& mkdir -p /home/chrome/reports && chown -R chrome:chrome /home/chrome

VOLUME /home/chrome/reports
WORKDIR /home/chrome/reports
USER chrome

ENTRYPOINT [ "mmonitoring" ]
