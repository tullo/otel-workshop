go build
env $(cat .env | grep -v '^#') ./hotrod all
