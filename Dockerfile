FROM golang:1.19
WORKDIR /go/src/hdwallet
COPY . .
RUN go build -o bin/server main.go

RUN apt-get update
RUN apt-get -y install python3
RUN apt-get -y install python3-pip
RUN pip3 install --upgrade pip
RUN pip3 install hdwallet click click_aliases tabulate

CMD ["./bin/server"]