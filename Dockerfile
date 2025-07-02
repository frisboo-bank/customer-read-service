FROM golang:alpine AS builder

WORKDIR /app

RUN apk --update --no-cache add ca-certificates make protoc

COPY Makefile go.mod go.sum ./
RUN go env -w GOPROXY=https://goproxy.io,direct/
RUN make init && go mod download

COPY .. ../
COPY . ./
RUN make build/release

FROM gcr.io/distroless/base-debian10

WORKDIR /app

ENV GOTRACEBACK=single

COPY --from=builder /app/bin/customers-service /app/

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/app/customers-service"]
