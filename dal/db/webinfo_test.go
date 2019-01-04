package db

import (
	"blogger/model"
	"testing"
)

/*func TestGetWebInfo(t *testing.T) {
	webInfo, err := GetWebInfo()
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("%s",webInfo.Keywords)

}
*/

func TestUpdateWebInfo(t *testing.T) {
	webInfo := &model.WebInfo{}
	webInfo.WebName = "1"
	webInfo.Keywords ="2"
	webInfo.Description = "3"
	webInfo.Url = "http://www.abc.com"

	effect_num, err := UpdateWebInfo(webInfo)

	if err != nil {
		t.Logf("err%v",err)
	}
	t.Logf("%v",effect_num)
}