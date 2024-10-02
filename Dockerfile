# Step 1: Use the official Golang image as a base
FROM golang:1.23-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR ./

# Step 3: Copy the Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Step 4: Copy the rest of the application code
COPY . .

# Step 5: Build the Go application
RUN go build -o main ./main.go

# Step 6: Use a minimal base image to run the application
FROM alpine:latest
WORKDIR /root/

# Step 7: Copy the binary from the builder stage
COPY --from=builder /gin-gorm-crud/main .

# Step 8: Expose the port the application will run on
EXPOSE 8080

# Step 9: Command to run the application
ENTRYPOINT [ "./main" ]  # Run the compiled Go binary
