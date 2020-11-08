FROM golang:1.14.3-alpine
ADD . /go/src/counter
RUN go install counter

FROM alpine:latest
COPY --from=0 /go/bin/counter .
ENV PORT 8080
CMD ["./counter"]