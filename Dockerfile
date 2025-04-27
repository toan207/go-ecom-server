FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ecom.pawtopia.vn ./cmd/server

FROM scratch
COPY ./configs /configs
COPY --from=builder /build/ecom.pawtopia.vn /

ENTRYPOINT [ "/ecom.pawtopia.vn", "configs/local.yaml" ]