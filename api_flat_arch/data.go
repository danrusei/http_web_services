package main

import (
	"encoding/json"
	"fmt"
)

// PopulateItems insert the items in empty database
func (a *api) PopulateItems() {
	defaultItems := []byte(`[{
			"id":1,
			"name":"Milk-(False)",
			"manufactured":"23-07-2019",
			"expdate":"23-08-2019",
			"expopen":10,
			"isopen":true,
			"opened":"30-07-2019"},{
			"id":2,
			"name":"Milk2-(False)",
			"manufactured":"23-07-2019",
			"expdate":"23-12-2019",
			"expopen":10,
			"isopen":true,
			"opened":"30-07-2019"},{
			"id":3,
			"name":"CannedFish-(True)",
			"manufactured":"15-11-2018",
			"expdate":"10-10-2020",
			"expopen":30,
			"isopen":true,
			"opened":"20-08-2019"},{
			"id":4,
			"name":"Butter-(False)",
			"manufactured":"15-07-2019",
			"expdate":"23-08-2019",
			"expopen":20,
			"isopen":false},{
			"id":5,
			"name":"CannedBeans-(True)",
			"manufactured":"24-02-2019",
			"expdate":"10-08-2020",
			"expopen":5,
			"isopen":false}]`)

	data := make([]Item, 4)
	if err := json.Unmarshal(defaultItems, &data); err != nil {
		fmt.Println("Could not unmarshal data:", err)
		return
	}
	a.addGood(data...)
}
