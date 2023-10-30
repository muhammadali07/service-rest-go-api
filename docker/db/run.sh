# Build Docker image
docker build -f Dockerfile -t restgo-db .

# Run Docker container
docker run -d --name restgo-db -p 5432:5432 restgo-db