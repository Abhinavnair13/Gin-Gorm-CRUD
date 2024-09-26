# Step 1: Use the official Golang image as a base
FROM golang:1.23-alpine AS builder

# Step 2: Set the working directory inside the container
WORKDIR /gin-gorm-crud

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

# Step 8: Set environment variables for your application (if needed)
# These values will be overwritten by the ones in the .env file
ENV PORT=8080
# Set environment variables for the database
ENV DB_URL="host=localhost user=postgres password=password@pg dbname=pg_db port=5432 sslmode=disable"


# Step 9: Expose the port the application will run on
EXPOSE 8080

# Step 10: Command to run the application
ENTRYPOINT [ "./main.go" ] 
