FROM golang:alpine

WORKDIR /app

ENV GITHUB_BASE_URL=https://api.github.com
ENV GITHUB_APP_ID=127997
ENV GITHUB_APP_PRIVATE_KEY_PATH=/app/orange-stack-test.private-key.pem
ENV GITHUB_APP_WEBHOOK_SECRET=123

COPY . .

EXPOSE 8000

CMD ["go", "run", "main.go"]