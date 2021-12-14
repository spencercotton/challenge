FROM golang:1.16

WORKDIR /opt/app

# Copy app files over
COPY . .

# Disable cgo for a static binary
ENV CGO_ENABLED=0

# Ensure all dependencies are downloaded
RUN go get -d -v ./...

# Build the application
RUN go build -o coupons ./cmd/coupons/main.go

# Run the application
CMD ["./coupons"]
