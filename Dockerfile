# Start from golang base image
FROM golang:1.19.5-alpine3.17

# Set maintainer (IT'S ME!!)
LABEL maintainer = "dozheiny"

RUN mkdir /app
WORKDIR /app

COPY ./ ./

# Start builder.
RUN go build -o . ./...

EXPOSE 80
CMD ["./http"]