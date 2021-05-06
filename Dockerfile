#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d
RUN go build -o /go/bin/app 

#final stage
FROM alpine:latest
COPY --from=builder /go/bin/app /app
ENV PORT=44433
ENV MONGODB_URL=mongodb://mongo
ENV SECRET_KEY=KSOR43HOGILVCNFNPBZGQMHMAUZXAYS9HMTYNOT65N3HYN20C1WLMOMP7GHVIISX
ENTRYPOINT /app
LABEL Name=authenticationservice Version=0.0.1
EXPOSE 44433
