FROM golang:1.20

WORKDIR /app

COPY  . /app/

RUN go mod tidy

RUN go build main.go

RUN mv /app/main /tmp/main

RUN rm -rf /app/*   

RUN mv /tmp/main /app/main 

CMD [ "./main" ]