# Go Lang websocket server

#### Building an image
`$docker build . -f Dockerfile.multistage -t BUILD_IMAGE_NAME --build-arg port=5000`

#### Starting a container
`$docker run -it --name DOCKER_NAME -p 5000:5000 BUILD_IMAGE_NAME`

---
#### Setup
You can define port in image build, setting as parameter
`--build-arg port=PORT`

The code was ben adapted to consumes port setted from variable