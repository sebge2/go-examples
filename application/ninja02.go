package application

import (
	"encoding/json"
	"fmt"
)

func Ninja02() {
	type person struct {
		First   string
		Last    string
		Age     int
		Sayings []string
	}

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	var persons []person

	err := json.Unmarshal([]byte(s), &persons)
	if err != nil {
		fmt.Printf("Error appened %s\n", err)
	} else {
		fmt.Println(persons)
	}
}
