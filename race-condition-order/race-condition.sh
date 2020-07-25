curl --request POST \
  --url http://localhost:9500/api/order/save \
  --header 'content-type: application/json' \
  --cookie refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyZWZyZXNoIjp0cnVlLCJpYXQiOjE1OTQ0NzI1ODcsImV4cCI6MTU5NDU1ODk4N30.zMcNmsk3JoOpFIZzToy3bqKpoJLR8QU2XL0PUfhYeZI \
  --data '{
	"phoneNumber": "08112232320",
	"created": "nabil",
	"total": 300000,
	"items": [
		{
			"sku": "product-1",
			"quantity": 2
		}
	]
}' & curl --request POST \
  --url http://localhost:9500/api/order/save \
  --header 'content-type: application/json' \
  --cookie refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyZWZyZXNoIjp0cnVlLCJpYXQiOjE1OTQ0NzI1ODcsImV4cCI6MTU5NDU1ODk4N30.zMcNmsk3JoOpFIZzToy3bqKpoJLR8QU2XL0PUfhYeZI \
  --data '{
	"phoneNumber": "08112232320",
	"created": "nabil",
	"total": 300000,
	"items": [
		{
			"sku": "product-1",
			"quantity": 2
		}
	]
}'