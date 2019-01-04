package logic

import (
	"blogger/model"
	"encoding/json"
	"fmt"
	"os"
)

func UpWebInfo(web_name, keywords, description, url string)(webInfo *model.WebInfo, err error){
	//添加结构图数据
	//data,err := db.GetWebInfo()

	data := &model.WebInfo{}
	data.WebName = web_name
	data.Keywords = keywords
	data.Description = description
	data.Url = url

	//创建文件并打开
	file, err := os.Create("/config/config.json")

	//关闭文件
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		fmt.Printf("编码失败%v",err)
	} else {
		fmt.Printf("编码成功%v",err)
	}

	return
}
