FROM golang:1.18

# set working directory
WORKDIR /go/src/app

# Copy the source code
COPY . .

# EXPOSE the port
EXPOSE 8080

#Build the GO app
RUN go build -o main cmd/main.go

# Run the executable
CMD ["./main"]