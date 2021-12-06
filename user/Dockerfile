FROM alpine

RUN go build -o user *.go

ADD user /user

ENTRYPOINT ["/user"]

