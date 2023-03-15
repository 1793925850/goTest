package dao

import (
	"USSTTB/models"
	"fmt"
	"log"
)

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from postinfo")
	_ = rows.Scan(&count)
	return
}

func CountGetAllPostByslug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from postinfo where slug=?", slug)
	_ = rows.Scan(&count)
	return
}

func CountGetAllPostBycategoryId(cid int) (count int) {
	rows := DB.QueryRow("select count(1) from postinfo where category_id=?", cid)
	_ = rows.Scan(&count)
	return
}

func GetPostPage(page, pagesize int) ([]models.Post, error) {
	page = (page - 1) * pagesize
	rows, err := DB.Query("select * from postinfo limit ?,?", page, pagesize)
	if err != nil {
		fmt.Println("没有载入post")
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content,
			&post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		post.Content = post.Content[0:30]
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select * from postinfo ")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content,
			&post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBycategoryId(cid, page, pagesize int) ([]models.Post, error) {
	page = (page - 1) * pagesize
	rows, err := DB.Query("select * from postinfo where category_id = ? limit ?,?", cid, page, pagesize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content,
			&post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageBySlug(slug string, page, pagesize int) ([]models.Post, error) {
	page = (page - 1) * pagesize
	rows, err := DB.Query("select * from postinfo where slug = ? limit ?,?", slug, page, pagesize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content,
			&post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostById(pid int) (models.Post, error) {
	rows := DB.QueryRow("select * from postinfo where pid=?", pid)
	var post models.Post
	if rows.Err() != nil {
		return post, rows.Err()
	}
	err := rows.Scan(&post.Pid, &post.Title, &post.Content,
		&post.Markdown, &post.CategoryId, &post.UserId,
		&post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
	if err != nil {
		return post, err
	}
	return post, nil
}

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into postinfo (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at )"+
		"values(?,?,?,?,?,?,?,?,?,?)", post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update postinfo set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid =?",
		post.Title, post.Content, post.Markdown, post.CategoryId, post.Type, post.Slug, post.UpdateAt,
		post.Pid)
	if err != nil {
		log.Println(err)
	}

}

func GetPostSearch(conditon string) ([]models.Post, error) {
	rows, err := DB.Query("select * from postinfo where title like ?", "%"+conditon+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content,
			&post.Markdown, &post.CategoryId, &post.UserId,
			&post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
