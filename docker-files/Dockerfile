FROM golang:1.11-alpine AS build
WORKDIR /src/
COPY src/main.go  /src/
RUN CGO_ENABLED=0 go build -o /bin/api-demo

FROM scratch
COPY --from=build /bin/api-demo /bin/api-demo
ENTRYPOINT ["/bin/api-demo"]
