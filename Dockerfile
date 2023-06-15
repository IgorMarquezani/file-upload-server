FROM golang:1.19-alpine

WORKDIR /src/app

COPY . ./

RUN go mod download 
RUN go mod verify
RUN go build main.go

EXPOSE 3300

CMD [ "./main" ]
