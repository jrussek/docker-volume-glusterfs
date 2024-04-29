FROM golang:1.22 as builder
COPY . /go/src/github.com/jrussek/docker-volume-glusterfs
WORKDIR /go/src/github.com/jrussek/docker-volume-glusterfs
RUN go mod vendor
RUN go install --ldflags '-extldflags "-static"'
CMD ["/go/bin/docker-volume-glusterfs"]

FROM ubuntu:24.04
RUN apt-get update \
  && apt-get install tini software-properties-common -y \
  && add-apt-repository ppa:gluster/glusterfs-11 \
  && apt-get update \
  && apt-get install glusterfs-client -y \
  && apt-get purge software-properties-common -y \
  && apt-get autoremove -y \
  && rm -rf /var/lib/apt/lists/*
COPY --from=builder /go/bin/docker-volume-glusterfs /bin/
CMD ["docker-volume-glusterfs"]

