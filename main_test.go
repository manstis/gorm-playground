package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	obj := Person{
		Address: Address{
			AddressType: "home",
			AddressLines: []AddressLine{{
				Text: "Line1",
			},
				{
					Text: "Line2",
				}},
		},
	}

	DB.Debug().Create(&obj)

	var result People
	dbConn := DB.Table("people")
	dbConn = dbConn.Select(
		"people.*",
		"a.*",
		"al.*",
	)
	dbConn = dbConn.Joins("LEFT JOIN addresses a ON a.id = people.address_id")
	dbConn = dbConn.Joins("LEFT JOIN address_lines al ON al.address_id = a.id")
	dbConn.Preload("Address")
	dbConn.Preload("Address.AddressLines")

	if err := dbConn.Debug().Find(&result, obj.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Println("Resultset size: ", len(result))

	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}
