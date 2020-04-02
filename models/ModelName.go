package models

type madame struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetModelNames(pageNum int, pageSize int, maps interface{}) (modelNames []madame) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&modelNames)

	return
}

func GetModelNameTotal(maps interface{}) (count int) {
	db.Model(&madame{}).Where(maps).Count(&count)

	return
}
func ExistModelNameByName(name string) bool {
	var modelName madame
	db.Select("id").Where("name = ?", name).First(&modelName)
	if modelName.ID > 0 {
		return true
	}

	return false
}

func AddModelName(name string, state int, createdBy string) bool {
	db.Create(&madame{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})

	return true
}

func ExistModelNameByID(id int) bool {
	var modelName madame
	db.Select("id").Where("id = ?", id).First(&modelName)
	if modelName.ID > 0 {
		return true
	}

	return false
}

func DeleteModelName(id int) bool {
	db.Where("id = ?", id).Delete(&madame{})

	return true
}

func EditModelName(id int, data interface{}) bool {
	db.Model(&madame{}).Where("id = ?", id).Updates(data)

	return true
}
