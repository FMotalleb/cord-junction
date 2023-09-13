FROM golang:1.20.6-alpine3.18 AS builder
RUN mkdir /app
COPY ./ /app
WORKDIR /app
RUN go build -o cord . 

FROM gcr.io/distroless/base-debian11:nonroot AS distroless-runtime

WORKDIR /app
COPY --from=builder /app/cord /app/

EXPOSE 80
EXPOSE 80/udp

# watching is not supported in container
# ENV WATCH_CONFIG_FILE "false"

ENTRYPOINT [ "/app/cord" ]


FROM alpine:latest AS runtime
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/cord /app/
RUN chmod +x /app/cord
EXPOSE 80
EXPOSE 80/udp

RUN apk del apk-tools
ENV PATH "/app:$PATH"
ENV LOG_LEVEL info
ENV LOG_FILE "/app/cord.log"
ENV CONFIG_FILE "/app/config.yaml"

# watching is not supported in container
# ENV WATCH_CONFIG_FILE "false"

ENTRYPOINT [ "/app/cord" ]
