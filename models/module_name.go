package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type ModuleName struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetModuleName(pageNum int, pageSize int, maps interface{}) (moduleNames []ModuleName) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&moduleNames)

	return
}

func GetModuleNameTotal(maps interface{}) (count int) {
	db.Model(&ModuleName{}).Where(maps).Count(&count)

	return
}

func ExistModuleNameByName(name string) bool {
	var madalena ModuleName
	db.Select("id").Where("module_name = ?", name).First(&madalena)
	if madalena.ID > 0 {
		return true
	}

	return false
}

func AddModuleName(name string, state int, createdBy string) bool {
	db.Create(&ModuleName{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func EditModuleName(id int, data interface{}) bool {
	db.Model(&ModuleName{}).Where("id = ?", id).Updates(data)

	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
