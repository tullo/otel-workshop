# Elastic

## Start Elastic APM

```sh
# Launch and access multipass VM instance.
multipass launch -c3 -d 10G -m 4G -n elastic \
    && multipass shell elastic

# Install docker.
sudo apt install docker.io python3-pip

# Install docker-compose
sudo pip3 install docker-compose

# Install go
wget https://go.dev/dl/go1.17.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
echo 'PATH=$PATH:/usr/local/go/bin' >> .bashrc

# Clone git repo
git clone https://github.com/riferrei/otel-with-golang.git \
    && cd otel-with-golang

# Build go binary
go mod download && go mod tidy && go build .

# Launch elastic apm service containers
docker-compose -f run-without-collector.yaml up -d

# Check service status
docker-compose -f run-without-collector.yaml ps
```

## Start Sample App:

```sh
export ELASTIC_IP=$(multipass info elastic | grep IPv4 | awk '{print $2}')
export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=http://${ELASTIC_IP}:8200
export OTEL_RESOURCE_ATTRIBUTES=service.name=fib,application=workshop
export SERVICE_NAME=fib

cd cmd/elastic
go build
./elastic
Your server is live!
Try to navigate to: http://127.0.0.1:3000/fib?n=6
```
