FROM alpine:3.5

COPY currencies /currencies
COPY config/*.yaml /

ENTRYPOINT ["/currencies"]