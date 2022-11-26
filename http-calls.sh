curl --location --request POST 'http://localhost:8080/api/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account": "third",
    "operation": "add",
    "amount": 140.7
}'

curl --location --request POST 'http://localhost:8080/api/balance' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account": "third",
    "operation": "withdraw",
    "amount": 37.44
}'