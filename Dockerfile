# Mulai dari golang base image
FROM golang:alpine as builder

# Menambahkan informasi maintainer
LABEL maintainer="Fadhli Mulyana <baba.fadhli@gmail.com>"

# Install git
RUN apk update && apk add --no-cache ca-certificates git

# Mengatur working directory dalam container
WORKDIR /app

# Meng-copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download semua depedency
RUN go mod download

# Copy semua source code dari current directory ke current working directory
COPY . .

# Build aplikasi golang
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Memulai stage baru dari scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/database/migrations /root/database/migrations
COPY --from=builder /app/internal/config/casbin /root/internal/config/casbin 
COPY --from=builder /app/internal/template /root/internal/template
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port 80 to the outside world
# EXPOSE 443
EXPOSE 3000

# VOLUME ["/cert-cache"]

#Command to run the executable
CMD ["./main", "serve"]
