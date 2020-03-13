FROM golang:latest

LABEL maintainer="prashc <prashanthc97@gmail.com>"

WORKDIR /app

# Copying everything from the current directory to the Working Directory inside the container
COPY . .

# Installing Dependancy
RUN go get github.com/julienschmidt/httprouter

# Build the Go app
RUN go build -o factorial .

# Expose port 8989 to the outside world
EXPOSE 8989

# Command to run the executable
CMD ["./factorial"]