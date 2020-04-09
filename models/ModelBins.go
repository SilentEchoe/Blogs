package models

type MadalenaAttrValue struct {
	Model

	MadalenaTypeId string `json:"madalena_type_id"`
	attrKey        string `json:"attr_key"`
	attrValue      string `json:"attr_value"`
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
func GetBin(pageNum int, pageSize int, maps interface{}) (bins []Bins) {

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&bins)
	return
}
