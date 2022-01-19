FROM golang:1.17

# create a directory named /app
RUN mkdir /app

# make /app our working directory
WORKDIR /app

#copy all files from this directory to app
COPY ./ /app

#build/compile
RUN go build -o api

# the command that should be executed when
#the container is started
CMD ./api