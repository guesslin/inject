FROM golang:1.12 AS build-layer

RUN apt-get update && apt-get install -y libpcap0.8-dev
COPY . /apps
WORKDIR /apps
RUN CGO_ENABLE=false go build -o inject


FROM debian AS finale-layer

RUN mkdir -p /apps
WORKDIR /apps
COPY --from=build-layer /apps/inject /apps

CMD /bin/bash
