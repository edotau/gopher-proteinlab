FROM golang:1.22.2-alpine


# Set up working directory and add all golang files to container
WORKDIR /src/github.com/edotau/gopherproteinlab
COPY go.mod go.sum ./

# Download all gopherproteinlab dependencies and modules
RUN go mod download && go mod verify

# Add gopherproteinlab repository files into container
COPY . .

# Install all binary executables into $GOBIN
RUN go install ./...
