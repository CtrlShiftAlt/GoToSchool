package iniconf

// IniRedis Redis类
type IniRedis struct {
	Host    string `tag:"host"`
	Port    string `tag:"port"`
	Auth    string `tag:"auth"`
	Select  string `tag:"select"`
	Expire  string `tag:"expire"`
	Prefix  string `tag:"prefix"`
	TimeOut string `tag:"timeout"`
}
