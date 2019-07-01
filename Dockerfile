FROM golang:1.12

RUN apt-get update && apt-get install -y libpcap0.8-dev
COPY . /apps
WORKDIR /apps
RUN go build

CMD /bin/bash
