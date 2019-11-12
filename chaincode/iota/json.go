package iota

// https://stackoverflow.com/a/49006864

import (
  "encoding/json"
  "fmt"
)

type Animal struct {
	Name  string
	Order string
}

var st = `[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`


func ConvertAnimal() {
	var kpi interface{} = st
	var a []Animal
	err := json.Unmarshal([]byte(kpi.(string)), &a)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(a)
}

func ConvertUser() {
    // ********************* Marshal *********************
    u := map[string]interface{}{}
    u["name"] = "kish"
    u["age"] = 28
    u["work"] = "engine"
    u["hobbies"] = []string{"art", "football"}
    // u["hobbies"] = "art"

    b, err := json.Marshal(u)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(b))

    // ********************* Unmarshal *********************
    var a interface{}
    err = json.Unmarshal(b, &a)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(a)
}
