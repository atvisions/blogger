package db

import (
	"blogger/model"
	"fmt"
)

//获取网站信息
func GetWebInfo()(webInfo *model.WebInfo, err error){
	webInfo = &model.WebInfo{}
	sqlStr := "select web_name,keywords,description,url from web_info"
	err = DB.Get(webInfo, sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	return
}

//更新网站信息
func UpdateWebInfo(webInfo *model.WebInfo)(effect_num int64, err error){
	sqlStr := `update web_info SET web_name=?,keywords=?,description=?,url=? `
	result, err := DB.Exec(sqlStr, webInfo.WebName,webInfo.Keywords, webInfo.Description, webInfo.Url)
	if err != nil {
		fmt.Println(err)
	}
	effect_num, err = result.RowsAffected()
	return effect_num, err
}


