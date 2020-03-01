package models

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

// 查看标签是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	// 查询标签
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

// 添加标签
func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:name,
		State:state,
		CreatedBy:createBy
	})

	return true
}
