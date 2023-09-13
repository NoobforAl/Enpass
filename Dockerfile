FROM golang:1.20

WORKDIR /app

COPY . /app/

RUN go build -o /tmp/enpass ./cmd/main.go
RUN rm -rf /app/*   
RUN mv /tmp/enpass /app/enpass 

CMD [ "./enpass" ]