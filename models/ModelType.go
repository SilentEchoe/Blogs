package models

type MadalenaType struct {
	Model

	MadalenaId int      `json:"madalena_id" gorm:"index"`
	Madalena   Madalena `json:"Madalena"`

	CompatibleType   string `json:"compatible_type"`
	Type             string `json:"type"`
	ProcessingMethod string `json:"processing_method"`
	ThresholdValue   int    `json:"threshold_value"`
	State            int    `json:"state"`
}

func GetModelTypes(pageNum int, pageSize int, maps interface{}) (madalenaType []MadalenaType) {
	db.Preload("Madalena").Where(maps).Offset(pageNum).Limit(pageSize).Find(&madalenaType)

	return
}

func GetModelTypeTotal(maps interface{}) (count int) {
	db.Model(&MadalenaType{}).Where(maps).Count(&count)

	return
}

func GetModelTypeId(modelId int, compatibleType string) int {
	var madalena MadalenaType
	db.Where(&MadalenaType{CompatibleType: compatibleType, MadalenaId: modelId}).First(&madalena)
	return madalena.ID
}

// 查询bin模板
func GetBinTemplate(maps interface{}) (madalenaType []MadalenaType) {
	db.Where(maps).First(&madalenaType)
	return
}
