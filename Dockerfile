# The infrastructure shell: members are compiled in from their modules at the versions go.mod pins (no replace).
FROM golang:1.27.2-alpine AS build
WORKDIR /src
# git: when the public proxy answers 404 for a module released minutes ago (it caches a failed fetch for about half an
# hour), go falls back to the repository itself; go.sum checks what it gets (C10 retro R2)
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /shell .

FROM alpine:3.24.2
COPY --from=build /shell /shell
# the shell's own port (health, readiness, description, every member's metrics); members listen on their own ports
ENV PORT=8224
EXPOSE 8224
ENTRYPOINT ["/shell"]
