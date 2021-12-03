go build -tags appsec
env --debug $(cat .env | grep -v '^#') ./datadog
