# Inventory App using layered architecture  

Endpoint List Items:
curl http://localhost:8080 | json_pp
Endpoint Add Item:
curl --header "Content-Type: application/json" --request POST --data "@add_item.json" http://localhost:8080/add
Endpoint Change Open Status for Item:
curl -sS 'http://localhost:8080/open?id=5&open=true'
Endpoint Delete Item:
curl http://localhost:8080/del?id=2
