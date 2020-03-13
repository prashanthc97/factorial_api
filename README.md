# factorial_api
Takes json with following structure {"a":int,"b":int} and calculates factorial of a and b using goroutines
## INSTRUCTIONS TO RUN DOCKER
1. Building the image
>docker build -t go-docker-Meant4 .
2. Running the Docker image
>docker run -d -p 8989:8989 go-docker-Meant4
3. Testing the API
>curl -H "Content-Type: application/json" -X POST -d {\"a\":6\,\"b\":12} http://localhost:8989/calculate
