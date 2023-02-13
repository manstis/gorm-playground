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

	DB.Create(&obj)

	var result People
	if err := DB.Preload("Address.AddressLines").Find(&result).Joins("Address").Joins("AddressLine").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Println("=STORE====================")
	fmt.Println("Resultset size: ", len(result))
	b, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(b))

	if err := DB.Delete(&result[0].Address).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	fmt.Println("=UNSCOPED LOAD====================")
	var result2 People
	if err := DB.Unscoped().Preload("Address.AddressLines").Find(&result2).Joins("Address").Joins("AddressLine").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	b2, _ := json.MarshalIndent(result2, "", "  ")
	fmt.Println(string(b2))

	fmt.Println("=SCOPED LOAD====================")
	var result3 People
	if err := DB.Preload("Address.AddressLines").Find(&result3).Joins("Address").Joins("AddressLine").Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
	b3, _ := json.MarshalIndent(result3, "", "  ")
	fmt.Println(string(b3))

	fmt.Println("=ADDRESSES====================")
	address := &Address{}
	DB.Model(&Address{}).Preload("AddressLines").Unscoped().Find(address)
	b4, _ := json.MarshalIndent(address, "", "  ")
	fmt.Println(string(b4))

}
