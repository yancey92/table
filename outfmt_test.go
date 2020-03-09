package lksctl

import (
	"fmt"
	"testing"
)

func TestFmtGet(t *testing.T) {

	type Student struct {
		Name   string   `head:"NAME"`
		Age    int      `head:"AGE"`
		Gender string   `head:"GENDER"`
		Hobby  []string `head:"HOBBY"`
	}

	var stuSlice = []Student{
		{
			Name:   "Zhang san",
			Age:    23,
			Gender: "Man",
			Hobby:  []string{"Basketball", "Ping pang ball"},
		},
		{
			Name:   "Li si",
			Age:    22,
			Gender: "WoMan",
			Hobby:  []string{"Swimming"},
		},
		{
			Name:   "Wang wu",
			Age:    21,
			Gender: "Man",
		},
	}

	table, err := CreateTableCell(stuSlice)
	if err != nil {
		PrintTableBox(table)
	} else {
		fmt.Println(err)
	}

}
