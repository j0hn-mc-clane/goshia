# Still experimental, wont probably work when the script is outside of your Docker container

FROM golang:1.20

WORKDIR /app

COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /goshia

# Port for Gin
EXPOSE 8080


CMD [ "./goshia"]