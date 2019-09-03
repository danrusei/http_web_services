# Inventory App using layered architecture

Inspired by Matt Ryer article:
https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831
and Florin Patan demo:
https://www.youtube.com/watch?v=wxkEQxvxs3w

Endpoint List Items: 
curl http://localhost:5000 | json_pp
Endpoint Add Item:
curl --header "Content-Type: application/json" --request POST --data "@add_item.json" http://localhost:5000/add
Endpoint Change Open Status for Item:
curl -sS 'http://localhost:5000/open?id=5&open=true'
Endpoint Delete Item:
curl http://localhost:5000/del?id=2
