FROM golang:1.16-buster AS build-layer

RUN apt-get update -qq && \
    apt-get install -y flex cmake bison

ARG pcap_version=1.5.3
RUN cd /tmp/ && curl -LO https://github.com/the-tcpdump-group/libpcap/archive/libpcap-${pcap_version}.tar.gz && \
    tar xzf libpcap-${pcap_version}.tar.gz && \
    cd libpcap-libpcap-${pcap_version}/ && \
    ./configure && \
    make && make install
# to help unit test can link libpcap, extend LD_LIBRARY_PATH
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/lib
COPY . /apps
WORKDIR /apps
RUN CGO_ENABLE=false go build -o inject


FROM debian:buster AS finale-layer

RUN apt-get update && apt-get install -y libpcap0.8-dev
WORKDIR /apps
COPY --from=build-layer /apps/inject /usr/local/bin/

CMD /bin/bash
