package model

type SettingsQuery interface {
	SettingByKey(key string) *Setting
}

type Setting struct {
	ID    string
	Key   string
	Value string
}

func (Setting) TableName() string {
	return "settings"
}

func (q *DBDataQuery) SettingByKey(key string) *Setting {
	var setting Setting
	if q.DB.Where("key = ?", key).First(&setting).RecordNotFound() {
		return nil
	}
	return &setting
}
