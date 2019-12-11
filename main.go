package main

import (
	"fmt"
	// "github.com/sirupsen/logrus"
	"encoding/json"
	ameritrade "gitlab.com/austin-millan/tdameritrade_openapi/openapi"
)

func main() {
	msg := ameritrade.Order{
		Session: "AM",
		Duration: "GOOD_TILL_CANCEL",
	}
	res, _ := json.Marshal(msg)
	fmt.Println(string(res))
}
