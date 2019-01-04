package db

import (
	"testing"
)
func init()  {
	InitDB()
}
/*func TestInsertCategory(t *testing.T) {
	category := &model.Category{}
	category.CategoryName = "python"
	category.CategoryNo = 5
	t.Logf("%#v",category)
	categoryId,err := InsertCategory(category)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(categoryId)

}*/

/*func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64
	categoryIds = append(categoryIds,1,2,3)
	categoryList,err := GetCategoryList(categoryIds)
	if err != nil{
		t.Errorf("错误信息：%d",err)
	}

	if len(categoryList) != len(categoryIds){
		t.Errorf("错误信息：%d,%d",len(categoryList),len(categoryIds))
	}

	for _,v :=  range categoryList{
		t.Logf("%d,%#v",v.Id,v)
	}
}*/

/*func TestGetAllCategoryList(t *testing.T) {

	categoryList,err := GetAllCategoryList()
	if err != nil{
		t.Errorf("错误信息：%d",err)
	}


	for _,v :=  range categoryList{
		t.Logf("%d,%#v",v.Id,v)
	}
}*/

/*func TestGetCategoryByid(t *testing.T) {
	category,err := GetCategoryByid(2)
	if err != nil{
		t.Errorf("错误信息：%d",err)
	}
	t.Logf("%v",category)

}
*/
/*func TestGetRelevantArticleByid(t *testing.T) {
	RelevantArticle ,err := GetRelevantArticleByid(1)
	if err != nil {
		fmt.Println(err)
	}

	for _,v := range RelevantArticle{
		t.Logf("%d,%s",v.Id,v.Title)
	}

}*/

/*func TestGetCategoryArticle(t *testing.T) {
	articleRecord, err := GetCategoryArticle(1,0,15)
	if err != nil {
		fmt.Println(err)
	}

	for _,v := range articleRecord{
		t.Logf("%d,%s",v.Id,v.Title)
	}
}
*/
func TestGetCategoryByid(t *testing.T) {
	categoryName,err := GetCategoryByid(1)
	if err != nil {
		return
	}
	t.Logf("%v",categoryName.CategoryName)

}