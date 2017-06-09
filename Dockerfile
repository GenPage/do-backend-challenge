FROM golang:1.8
MAINTAINER Dylan Page

# Install system
RUN go get github.com/genpage/do-backend-challenge

# Expose the system on port 8080
EXPOSE 8080

# run the system
CMD ["do-backend-challenge", "run"]
