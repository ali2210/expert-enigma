
# create dockerfile 

# pull golang-apline3
FROM golang:1.17.8-alpine3.15

# enable c option (c and go are friends now)
ENV CGO_ENABLED = 0

# open port 
ENV PORT = 3000

# create app folder in the current working directory
RUN mkdir -p /app

# Create application
ADD . /app

# change your directory where app reside
WORKDIR /app

# copy go-modules from the application
COPY go.mod go.sum ./

# optimized modules 
RUN go mod tidy 

# download latest modules if required
RUN go mod download


# build application inside the container directory
RUN go build -o wisdomenigma

# open port for the application traffic
# EXPOSE 3000

# Some information about the application
LABEL Info = "Application started with in a minutes please wait..."

LABEL Listener = "http://localhost:3000/"

LABEL version = "0.0.0-alpha"

# initiate the application
CMD ["/app/wisdomenigma"]