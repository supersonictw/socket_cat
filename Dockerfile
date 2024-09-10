# socket_cat
# (c) 2024 Shiranui (https://shiranui.xyz)

FROM golang:alpine AS builder
COPY . /workplace
WORKDIR /workplace
RUN apk add --no-cache make
RUN make && make clean-deps

FROM alpine:latest
COPY --from=builder /workplace/build/socket_cat /workplace/socket_cat
WORKDIR /workplace
ENTRYPOINT /workplace/socket_cat
EXPOSE 8000
