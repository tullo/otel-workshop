go build -o country-info
env $(cat .env) ./country-info
