package service

import (
	"USSTTB/config"
	"USSTTB/dao"
	"USSTTB/models"
	"html/template"
)

func GetAllQiuZhuInfo(slug string, page, pagesize int) (*models.HomeData, error) {

	categorys, err := dao.GetAllCategory()

	var posts []models.Post
	var total int
	if slug == "" {
		posts, err = dao.GetPostPageBycategoryId(2, page, pagesize)
		total = dao.CountGetAllPostBycategoryId(2)
	} else {
		posts, err = dao.GetPostPageBySlug(slug, page, pagesize)
		total = dao.CountGetAllPostByslug(slug)

	}

	if err != nil {
		return nil, err
	}
	var postMores []models.PostMore
	for _, post := range posts {
		categoryname := dao.GetCategoryNameById(post.CategoryId)
		username := dao.GetuserNameById(post.UserId)

		postmore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(post.Content),
			post.CategoryId,
			categoryname,
			post.UserId,
			username,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postmore)
	}
	//查询总页数

	pagescount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagescount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeData{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page == pagescount,
	}
	return hr, nil
}
