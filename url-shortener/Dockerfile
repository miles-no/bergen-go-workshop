FROM golang
WORKDIR /src
ADD . .
RUN go build -o /bin/url-shortener
ENV HTTP_ADDR :8080
EXPOSE 8080
ENTRYPOINT /bin/url-shortener
