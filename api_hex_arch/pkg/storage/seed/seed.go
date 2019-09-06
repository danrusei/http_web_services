package seed

import (
	"encoding/json"
	"fmt"

	"github.com/Danr17/http_web_services/api_hex_arch/pkg/adding"
)

// PopulateItems insert the items in empty database
func PopulateItems() []adding.Item {
	defaultItems := []byte(`[{
			"id":1,
			"name":"Milk-(False)",
			"manufactured":"2019-07-23T00:00:00Z",
			"expdate":"2019-08-23T00:00:00Z",
			"expopen":10,
			"isopen":true,
			"opened":"2019-07-30T00:00:00Z"},{
			"id":2,
			"name":"Milk2-(False)",
			"manufactured":"2019-07-23T00:00:00Z",
			"expdate":"2019-12-23T00:00:00Z",
			"expopen":10,
			"isopen":true,
			"opened":"2019-07-30T00:00:00Z"},{
			"id":3,
			"name":"CannedFish-(True)",
			"manufactured":"2019-11-15T00:00:00Z",
			"expdate":"2020-10-10T00:00:00Z",
			"expopen":30,
			"isopen":true,
			"opened":"2019-08-20T00:00:00Z"},{
			"id":4,
			"name":"Butter-(False)",
			"manufactured":"2019-07-15T00:00:00Z",
			"expdate":"2019-08-23T00:00:00Z",
			"expopen":20,
			"isopen":false},{
			"id":5,
			"name":"CannedBeans-(True)",
			"manufactured":"2019-02-24T00:00:00Z",
			"expdate":"2020-08-10T00:00:00Z",
			"expopen":5,
			"isopen":false}]`)

	data := make([]adding.Item, 4)
	if err := json.Unmarshal(defaultItems, &data); err != nil {
		fmt.Println("Could not unmarshal data:", err)
	}
	return data
}
