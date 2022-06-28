FROM golang:1.18-alpine

WORKDIR /github.com/threeengineers/hands-on-lab/02-unit-testing/feature

COPY . .

RUN go mod vendor
RUN go build -o calculator

EXPOSE 1323 8080 80

CMD ["./calculator"]