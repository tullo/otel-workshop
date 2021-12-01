# SigNoz

Single pane for complete metrics and traces, no need to shift to different systems

- [SigNoz vs Jaeger](../../cmd/signoz/readme.md)
- [Bookstore REST API using Gin and Gorm](../../cmd/signoz/bookstore/README.md)
- [Hot R.O.D. - Rides on Demand](../../cmd/signoz/hotrod/README.md)
- [Getting Started](https://signoz.io/docs/)
- [Technical Architecture](https://signoz.io/docs/architecture/)
- [An alternative to DataDog and New Relic](https://signoz.io/blog/datadog-vs-newrelic/#an-alternative-to-datadog-and-new-relic---signoz)
- [Self-hosting software & why it may be worth considering again now](https://signoz.io/blog/self-hosting-software-observability/)

## SigNoz APM and hotrod

### Start SigNoz:

```sh
multipass launch -n signoz
multipass shell signoz
git clone https://github.com/SigNoz/signoz.git
cd signoz/deploy
./install.sh
# Terminate services and remove named volumes.
sudo docker-compose -f docker/clickhouse-setup/docker-compose.yaml down -v
# Uncomment 'hotrod' & 'load-hotrod' services in docker-compose.yaml.
sudo docker-compose -f docker/clickhouse-setup/docker-compose.yaml up
```

### Start Hotrod:

```sh
export SIGNOZ_IP=$(multipass info signoz | grep IPv4 | awk '{print $2}')
export JAEGER_ENDPOINT=http://SIGNOZ_IP:14268/api/traces
cd /cmd/signoz/hotrod
./run.sh
```

### Start requesting "Rides On Demand":

- http://[::1]:8080/

### SigNoz UI

- http://SIGNOZ_IP:3000/application
