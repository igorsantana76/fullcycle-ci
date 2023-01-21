FROM golang:1.15

WORKDIR /app

COPY . .

RUN cd src && go build -o math

CMD ["./src/math"]\