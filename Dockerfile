# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-game-api

EXPOSE 8080

CMD [ "/docker-game-api" ]

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /docker-game-api /docker-game-api

EXPOSE 5040

USER nonroot:nonroot

ENTRYPOINT ["/docker-game-api"]