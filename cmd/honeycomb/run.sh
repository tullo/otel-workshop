go build
env --debug $(cat .env | grep -v '^#') ./honeycomb
