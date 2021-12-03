# Datadog

## Start Datadog Agent

```sh
# Launch and access multipass VM instance.
multipass launch -n datadog && multipass shell datadog

# Install datadog-agent
DD_AGENT_MAJOR_VERSION=7 DD_API_KEY=... DD_SITE="datadoghq.eu" \
    bash -c "$(curl -L https://s3.amazonaws.com/dd-agent/scripts/install_script.sh)"

# Bind to 0.0.0.0
sed -i 's/^bind_host: localhost$/bind_host: 0.0.0.0/g' /etc/datadog-agent/datadog.yaml

# Restart datadog-agent
sudo systemctl restart datadog-agent

sudo lsof -i -P -n | grep LISTEN
...
trace-age 43229        dd-agent    8u  IPv6  44257      0t0  TCP *:8126 (LISTEN)
```

## Start Sample App:

```sh
export DD_AGENT_IP=$(multipass info datadog | grep IPv4 | awk '{print $2}')
export DD_AGENT_ADDRESS=${DD_AGENT_IP}:8126
export SERVICE_NAME=fib
export DD_APPSEC_ENABLED=true
export DD_DEPLOYMENT=v0.3.3
export DD_ENV=dev

cd cmd/datadog
go build -tags appsec
./datadog
Your server is live!
Try to navigate to: http://127.0.0.1:3000/fib?n=6
```
