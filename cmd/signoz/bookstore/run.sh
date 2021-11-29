go build
env $(cat .env | grep -v '^#') ./gin-bookstore
