From golang:alpine

#Création d'un ou plusieur label (utilisateur)
LABEL maintainer="Régis Antignac antignac.regis@gmail.com"

WORKDIR /go/src/api

ADD . /go/src/api

RUN apk add --no-cache git \
    && go get github.com/satori/go.uuid \
    && go get github.com/gorilla/mux \
    && go get github.com/globalsign/mgo \
    && go install api \
    && apk del git

ENTRYPOINT [ "/go/bin/api" ]

EXPOSE 3000