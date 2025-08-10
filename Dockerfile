FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o kubeapp

FROM gcr.io/distroless/base
COPY --from=builder /app/kubeapp /kubeapp
CMD ["/kubeapp"]

