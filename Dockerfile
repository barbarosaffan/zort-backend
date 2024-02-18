# Use the official Go image as a parent image
FROM golang:1.22 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum to download dependencies
COPY go.* ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -v -o zort-backend ./cmd/zort-backend

# Use a lightweight Docker image (scratch) for the runtime environment
FROM scratch

# Copy the compiled application from the builder stage
COPY --from=builder /app/zort-backend /zort-backend

# Command to run the application
CMD ["/zort-backend"]
