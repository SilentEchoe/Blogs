package models

type MadalenaAttrValue struct {
	Model

	MadalenaTypeId int    `json:"madalena_type_id"`
	AttrKey        string `json:"attr_key"`
	AttrValue      string `json:"attr_value"`
	Bins           int    `json:"bin_template"`
	Version        string `json:"version"`
	Sn             string `json:"sn"`
}

type Bins struct {
	Model

	MadalenaAttrValueId string `json:"madalena_attr_value_id"`
	FilePath            string `json:"file_location"`
	FileName            string `json:"file_name"`
	IsDelete            int    `json:"is_delete"`
	BinFiles            string `json:"base_64"`
}

// 查找bin文件

func GetBin(madalenaTypeId int, attrKey string, attrValue string, version string) (madalenaAttrValue []MadalenaAttrValue) {
	db.Where(&MadalenaAttrValue{MadalenaTypeId: madalenaTypeId, AttrKey: attrKey, AttrValue: attrValue, Version: version}).First(&madalenaAttrValue)

	return
}
