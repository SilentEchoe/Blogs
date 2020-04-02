package models

type ModelName struct {
	Model

	ModelName  string `json:"model_name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetModelNames(pageNum int, pageSize int, maps interface{}) (tags []ModelName) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetModelNameTotal(maps interface{}) (count int) {
	db.Model(&ModelName{}).Where(maps).Count(&count)

	return
}
func ExistModelNameByName(name string) bool {
	var modelName ModelName
	db.Select("id").Where("name = ?", name).First(&modelName)
	if modelName.ID > 0 {
		return true
	}

	return false
}

func AddModelName(name string, state int, createdBy string) bool {
	db.Create(&ModelName{
		ModelName: name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func ExistModelNameByID(id int) bool {
	var modelName ModelName
	db.Select("id").Where("id = ?", id).First(&modelName)
	if modelName.ID > 0 {
		return true
	}

	return false
}

func DeleteModelName(id int) bool {
	db.Where("id = ?", id).Delete(&ModelName{})

	return true
}

func EditModelName(id int, data interface{}) bool {
	db.Model(&ModelName{}).Where("id = ?", id).Updates(data)

	return true
}
