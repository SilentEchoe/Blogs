package models

import (
	_ "time"

	_ "github.com/jinzhu/gorm"
)

type ModuleName struct {
	Model
	Id         string `json:"id"`
	ModuleName string `json:"module_name"`
	ParentId   string `json:"parent_id"`
	State      int    `json:"is_delete"`
}

func GetModuleName(pageNum int, pageSize int, maps interface{}) (moduleNames []ModuleName) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&moduleNames)

	return
}
