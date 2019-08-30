package main

import (
	"encoding/json"
	"fmt"
)

// PopulateItems insert the items in empty database
func (a *api) PopulateItems() {
	defaultItems := []byte(`[{
			"id":1,
			"name":"Milk",
			"manufactured":"23-07-2019",
			"expdate":"23-08-2019",
			"expopen":10,
			"isopen":false,
			"opened":"30-07-2019"},{
			"id":2,
			"name":"Butter",
			"manufactured":"15-07-2019",
			"expdate":"23-10-2019",
			"expopen":20,
			"isopen":false,
			"opened":"20-08-2019"}]`)

	if err := json.Unmarshal(defaultItems, &a.db.Items); err != nil {
		fmt.Println("Could not unmarshal data:", err)
		return
	}

}
