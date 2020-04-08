package models

type MadalenaType struct {
	Model

	CompatibleType   string `json:"compatible_type"`
	Type             string `json:"type"`
	ProcessingMethod string `json:"processing_method"`
	ThresholdValue   int    `json:"threshold_value"`
	State            int    `json:"state"`
}

func GetModelTypes(pageNum int, pageSize int, maps interface{}) (madalenaType []MadalenaType) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&madalenaType)

	return
}
