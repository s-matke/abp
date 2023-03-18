FROM golang:alpine as build_container
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o flight-app
#EXPOSE
# ENTRYPOINT [ "flight-app" ]

FROM alpine
COPY --from=build_container /app/flight-app /usr/bin
EXPOSE 8084
ENTRYPOINT [ "flight-app" ]