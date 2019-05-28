FROM golang:1.11.5-alpine3.8 as builder

# we want a static binary
ENV CGO_ENABLED=0

ENV GO_BUILD_TAGS=mockvpp

RUN apk add --update git make

COPY . /go/src/github.com/contiv/vpp

WORKDIR /go/src/github.com/contiv/vpp

RUN make contiv-bgp

FROM scratch

COPY --from=builder /go/src/github.com/contiv/vpp/cmd/contiv-bgp/contiv-bgp /contiv-bgp

ENTRYPOINT ["/contiv-bgp"]