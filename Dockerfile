FROM golang
WORKDIR /app
COPY lola lola
RUN go env -w GO111MODULE=auto
RUN go mod init lola
RUN go get github.com/go-redis/redis

# RUN go install github.com/go-redis/redis

# RUN go mod download
# RUN ls 
CMD [ "go", "run", "lola/lola.go" ] 