FROM golang:1.26 as build
WORKDIR /build
COPY . .
ENV CGO_ENABLED=0
RUN go build

FROM scratch
COPY --from=build /build/probes_example_server /
CMD ["/probes_example_server", "server"]
EXPOSE 8000
