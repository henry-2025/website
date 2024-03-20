FROM golang:1.22 AS build

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o ./site ./cmd/site

FROM scratch AS runtime
WORKDIR /usr/bin/app
COPY --from=build /usr/src/app/site ./
COPY static ./static

EXPOSE 8080

CMD ["./site"]
