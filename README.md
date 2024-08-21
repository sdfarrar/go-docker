# Golang auto reloading with Docker

### Build the image
```bash
docker build -t my-go-docker .
# podman build -t my-go-docker .
```

### Run the container
```bash
# Mount the air configuration file
docker run -p 8888:8888 --rm -v $(pwd):/app -v /app/tmp --name my-go-docker-air my-go-docker
# podman run -p 8888:8888 --rm -v $(pwd):/app -v /app/tmp --security-opt label=disable --name my-go-docker-air my-go-docker
```

#### Troubleshooting
If there's network issues when running `docker`, try adding `--network=host` to the command.
