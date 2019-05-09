FROM golang:1.12

COPY . /apps
RUN apt-get update && apt-get install -y libpcap0.8-dev
RUN cd /apps/cmd && go build

CMD /bin/bash