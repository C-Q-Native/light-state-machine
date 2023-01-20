FROM golang:1.19-alpine

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download

COPY light-state-machine ./

#RUN go build -o /light-state-machine

EXPOSE 80

CMD [ "/light-state-machine" ]