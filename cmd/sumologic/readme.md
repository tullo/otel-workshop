# Sumo Logic

## Start opentelemetry-collector

```sh
# Launch and access multipass VM instance.
multipass launch -n sumologic && multipass shell sumologic

# Install docker.
sudo apt install docker.io

# Get config template.
wget -O config.yaml https://raw.githubusercontent.com/SumoLogic/opentelemetry-collector-contrib/main/examples/non-kubernetes/gateway-configuration-template.yaml

# Edit config.yaml - replace traces_endpoint with generated collector endpoint (sumo web ui).

# Start collector container.
sudo docker run --rm -p 4317:4317 -p 55680:55680 -p 55681:55681 \
    -p 9411:9411   -p 6831:6831/udp -p 14250:14250 -p 14268:14268 \
    -v "${PWD}/config.yaml":/conf/config.yaml \
    public.ecr.aws/sumologic/opentelemetry-collector \
    --config /conf/config.yaml
```

## Start Sample App:

```sh
export SUMO_LOGIC_IP=$(multipass info sumologic | grep IPv4 | awk '{print $2}')
export OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=http://${SUMO_LOGIC_IP}:55681
export OTEL_RESOURCE_ATTRIBUTES=service.name=fib,application=workshop
cd cmd/sumologic
go build
./sumologic
Your server is live!
Try to navigate to: http://127.0.0.1:3000/fib?n=6
```
