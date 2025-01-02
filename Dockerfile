FROM golang:1.23

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod tidy

COPY . .
COPY .env .

RUN go build -o e-wallet-ums
RUN chmod +x e-wallet-ums

EXPOSE 8080
EXPOSE 7000

CMD [ "./e-wallet-ums" ]