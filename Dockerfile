# Use slim alpine image for a smaller footprint
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .

# Download dependencies
RUN go mod download

# Build the Go binary
RUN go build -o database-clustermgt-ms .

# Switch to a clean alpine image for the final image
FROM golang:1.22-alpine
WORKDIR /app
RUN mkdir /app/conf
# Copy binary and conf files
COPY --from=builder /app/database-clustermgt-ms /app/database-clustermgt-ms
RUN chmod 777 /app/database-clustermgt-ms
COPY --from=builder /app/conf/database.toml /app/conf/database.toml



# Expose port
EXPOSE 3002


# Define the command to run the application
CMD ["/app/database-clustermgt-ms"]