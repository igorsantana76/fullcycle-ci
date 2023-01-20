FROM golang:latest

WORKDIR /app

COPY . .

RUN cd src && go build -o math

CMD ["./src/math"]\