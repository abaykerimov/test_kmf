FROM alpine:3.18 as base

WORKDIR /app
COPY . .

RUN go build -o /test_kmf

FROM alpine

WORKDIR /

COPY --from=base /test_kmf /test_kmf

EXPOSE 8000

CMD ["/test_kmf", "serve"]