# Start from the latest golang base image
FROM golang:1.14
ARG APP=main

# Setting ENv
ENV APP ${APP}
ENV HOST "db"

# Add Maintainer Info
LABEL maintainer="Punit Bamaniya <punitkumarbamaniya@gmail.com>"
 
# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Build the Go app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags static_all -a -installsuffix cgo -ldflags '-extldflags "-static"' -o './bin/'${APP} cmd/${APP}/${APP}.go
 
# Expose port 8080 to the outside world
EXPOSE 9090

ADD assets/asset.sql /docker-entrypoint-initdb.d/ 

# Command to run the executable 
CMD ["bash", "-c", "./bin/${APP}"]
