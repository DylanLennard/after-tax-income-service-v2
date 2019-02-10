# Go here for example of why we should build this way: 
# https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .
zip main.zip main