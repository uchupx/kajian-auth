FROM golang:1.20-bullseye as base

# RUN adduser \
#   --disabled-password \
#   --gecos "" \
#   --home "/nonexistent" \
#   --shell "/sbin/nologin" \
#   --no-create-home \
#   --uid 65532 \
#   small-user

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go mod vendor
RUN go mod verify

RUN go install github.com/securego/gosec/v2/cmd/gosec@latest
RUN go install golang.org/x/vuln/cmd/govulncheck@latest

# Security Check
RUN gosec ./..
# Depedencies check
RUN govulncheck -scan=package ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main .

FROM gcr.io/distroless/static-debian11

COPY --from=base /main .

# USER small-user:small-user

CMD ["./main"]