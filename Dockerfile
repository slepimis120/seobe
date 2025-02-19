FROM ubuntu:latest

WORKDIR /app

RUN apt-get update && apt-get install -y \
    curl \
    git \
    bash \
    ca-certificates \
    docker.io \
    libicu-dev \
    && rm -rf /var/lib/apt/lists/*

COPY .env.local /app/.env.local

RUN curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | tee /usr/share/keyrings/githubcli-archive-keyring.gpg > /dev/null \
    && echo "deb [signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee -a /etc/apt/sources.list.d/github-cli.list \
    && apt-get update \
    && apt-get install -y gh

ARG GITHUB_ACCESS_TOKEN
ENV GH_TOKEN=$GITHUB_ACCESS_TOKEN

RUN echo $GH_TOKEN | gh auth login --with-token

RUN gh extension install github/gh-actions-importer

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
