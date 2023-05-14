From golang:latest 
WORKDIR /go/src

COPY .  ./ 
RUN go mod tidy
EXPOSE 8089

CMD ["tail", "-f", "/dev/null"]
