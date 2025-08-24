# VS Code の Dev Container ベース（Go 1.24 / Debian bookworm）
FROM mcr.microsoft.com/devcontainers/go:1-1.24-bookworm

# 開発に便利なツール（psqlクライアント等）
RUN apt-get update \
    && apt-get install -y --no-install-recommends postgresql-client ca-certificates curl make gcc \
    && rm -rf /var/lib/apt/lists/*

# 任意: Air（ホットリロード）
RUN go install github.com/air-verse/air@v1.62.0
ENV PATH="$PATH:/go/bin"
