go build
echo "http://localhost:8090/books/"
env $(cat .env | grep -v '^#') ./bookstore
