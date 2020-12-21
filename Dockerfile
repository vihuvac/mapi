FROM golang:1.15

WORKDIR /app

COPY ./src /app/src

RUN go get github.com/gorilla/mux \
  gopkg.in/go-playground/validator.v9 \
  github.com/jinzhu/gorm \
  github.com/jinzhu/gorm/dialects/postgres \
  github.com/Altoros/gorm-goose/cmd/gorm-goose \
  github.com/satori/go.uuid \
  github.com/dgrijalva/jwt-go

RUN go build -o mapi ./src/main.go

RUN chmod +x mapi

EXPOSE 8080

CMD [ "sh", "mapi" ]
