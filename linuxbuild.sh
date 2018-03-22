rm ec2-metadata-proxy
GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s'
