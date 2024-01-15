FROM golang:alpine as builder

ARG APP=magnetu
ARG VERSION=v1.0.0
ARG RELEASE_TAG=$(VERSION)

WORKDIR /go/src/${APP}
COPY . .
RUN apk add --no-cache gcc musl-dev
ENV GOPROXY="https://goproxy.cn"
RUN go build -ldflags "-w -s -X main.VERSION=${RELEASE_TAG}" -o ./${APP} .

FROM alpine
ARG APP=magnetu
WORKDIR /go/src/${APP}
COPY --from=builder /go/src/${APP}/${APP} /usr/bin/
COPY --from=builder /go/src/${APP}/build/configs /usr/bin/configs
COPY --from=builder /go/src/${APP}/build/dist /usr/bin/dist
ENTRYPOINT ["magnetu", "start", "-d", "/usr/bin/configs", "-c", "prod", "-s", "/usr/bin/dist"]
EXPOSE 8040
