package models

type ConfigFiles struct {
	Model

	ID             int    `json:"id"`
	ConfigFileName string `json:"config_file_name"`
	ConfigType     string `json:"config_type"`
	PowerDelay     string `json:"power_delay"`
	PageWriteDelay string `json:"page_write_delay"`
	ConfigPassword string `json:"config_password"`
	PageWriteByte  string `json:"page_write_byte"`
	PreOperation   string `json:"pre_operation"`
	IsDelete       int    `json:"isdelete"`
}

type ConfigFileManger struct {
	Model

	ID             int    `json:"id"`
	ModalId        int    `json:"modal_id"`
	CompatibleType string `json:"compatible_type"`
	ConfigFiles    string `json:"config_files"`
	enable         string `json:"enable"`
}

// 查询对应的配置文件
func GetConfigFileById(id []int) (confides []ConfigFiles) {
	db.Where("id IN (?)", id).Find(&confides)
	return
}

// 根据ID查到配置文件
func GetConfigsById(id int) string {
	var confideManger ConfigFileManger
	db.Where("modal_id = ?", id).First(&confideManger)
	return confideManger.ConfigFiles
}
