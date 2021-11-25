package model

import "goPhotoV2/utils"

// Info 结构体对应mysql中的字段
type Info struct {
	ID   int64
	Name string
	Path string
	Note string
	Unix int64
}

// AddInfo 上传图片信息
func AddInfo(mod *Info) error {
	str := "insert into info (name,unix,path,note) values(?,?,?,?)"
	_, err := utils.Db.Exec(str, mod.Name, mod.Unix, mod.Path, mod.Note)
	if err != nil {
		return err
	}
	return nil
}

// GetInfoByID 获取图片详细信息
func GetInfoByID(id int64) (mod *Info, err error) {
	str := "select * from info where id = ?"
	row := utils.Db.QueryRow(str, id)
	mod = &Info{}
	err = row.Scan(&mod.ID, &mod.Name, &mod.Unix, &mod.Path, &mod.Note)
	if err != nil {
		return mod, err
	}
	return mod, nil
}

// GetAllInfo 获取所有信息
func GetAllInfo() {
}
