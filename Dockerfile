FROM golang
WORKDIR /app
COPY lola /app/lola
RUN go env -w GO111MODULE=auto
RUN go mod init /app/lola
RUN go get github.com/go-redis/redis

# RUN go install github.com/go-redis/redis

# RUN go mod download
# RUN ls 
RUN go run lola/lola.go