FROM golang:1.13.8-alpine3.11
WORKDIR '/app'
COPY . .

EXPOSE 80
EXPOSE 443

CMD ['go','run','main.go']