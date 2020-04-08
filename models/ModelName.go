package models

type Madalena struct {
	Model

	ModuleName string `json:"module_name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
	ParentId   int    `json:"parent_id"`
}

func GetModelNames(pageNum int, pageSize int, maps interface{}) (modelNames []Madalena) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&modelNames)

	return
}

func GetModelNameTotal(maps interface{}) (count int) {
	db.Model(&Madalena{}).Where(maps).Count(&count)

	return
}

func ExistModelNameByName(name string) bool {
	var modelName Madalena
	db.Select("id").Where("module_name = ?", name).First(&modelName)
	if modelName.ID > 0 {
		return true
	}

	return false
}

func AddModelName(name string, state int, createdBy string, parentId int) bool {
	db.Create(&Madalena{
		ModuleName: name,
		State:      state,
		CreatedBy:  createdBy,
		ParentId:   parentId,
	})

	return true
}

func ExistModelNameByID(id int) bool {
	var modelName Madalena
	db.Select("id").Where("id = ?", id).First(&modelName)
	if modelName.ID > 0 {
		return true
	}

	return false
}

func DeleteModelName(id int) bool {
	db.Where("id = ?", id).Delete(&Madalena{})

	return true
}

func EditModelName(id int, data interface{}) bool {
	db.Model(&Madalena{}).Where("id = ?", id).Updates(data)

	return true
}
