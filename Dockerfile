FROM golang:alpine

WORKDIR /go-cart/api/

COPY . .

ENV HOST=${HOST}
ENV USER=${USER}
ENV PASS=${PASS}
ENV NAME=${NAME}
ENV PORT = ${PORT}

RUN go mod tidy
RUN go build -o start
RUN /go-cart/api/start migrate -u

CMD [ "/go-cart/api/start serve" ]