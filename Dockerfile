FROM golang:1.21-bookworm AS go-build

WORKDIR /go
ADD src/ghlatest ghlatest/

WORKDIR ghlatest
RUN go build -o ghlatest

WORKDIR /go
ADD src/html-to-pdf html-to-pdf/

WORKDIR html-to-pdf
RUN go build -o html-to-pdf

FROM debian:bookworm-slim

COPY --from=go-build /go/ghlatest/ghlatest /root/bin/ghlatest
COPY --from=go-build /go/html-to-pdf/html-to-pdf /usr/local/bin/html-to-pdf

ENV PATH=/opt/pandoc/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/root/bin

ADD bin/setup /root/bin/setup
RUN setup 

WORKDIR /cv
ENTRYPOINT ["make"]
