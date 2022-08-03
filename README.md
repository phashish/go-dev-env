# go-dev-env
My GO development environment


## Login to the registry
```
docker login ghcr.io --username shifu137
```
## Pull the image
```
docker pull ghcr.io/shifu137/go-dev-env
```
## Run the container
```
docker run -it --restart unless-stopped --name ubu02 -h ubu02 -u ashish ghcr.io/shifu137/go-dev-env:latest /bin/bash
```
## Add this to ~/.bashrc
```
echo "alias ubu02='docker exec -it ubu02 /bin/bash'" >> ~/.bashrc
. ~/.bashrc
```
