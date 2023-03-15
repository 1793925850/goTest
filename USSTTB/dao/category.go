package dao

import (
	"USSTTB/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {

	rows, err := DB.Query("select *from categoryinfo")
	if err != nil {
		log.Println("Getcategory err!", err)
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println("取值出错")
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, err

}

func GetCategoryNameById(Cid int) string {
	row := DB.QueryRow("select *from categoryinfo where cid=?", Cid)
	if row.Err() != nil {
		log.Println("GetCategoryNameById出错", row.Err())

	}
	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}
