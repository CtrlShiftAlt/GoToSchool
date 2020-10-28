package iniconf

// IniMySQL MySQL类
type IniMySQL struct {
	UserName string `tag:"username"`
	PassWord string `tag:"password"`
	Host     string `tag:"host"`
	Port     string `tag:"port"`
}
