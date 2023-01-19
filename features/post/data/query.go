package data

import (
	"errors"
	"log"
	"simple-social-media-API/features/comment"
	"simple-social-media-API/features/post"

	"gorm.io/gorm"
)

type postData struct {
	db *gorm.DB
}

func Isolation(db *gorm.DB) post.PostData {
	return &postData{
		db: db,
	}
}

func (pd *postData) Add(userID uint, newPost post.Core) (post.Core, error) {
	convert := CoreToData(newPost)
	convert.UserID = userID

	err := pd.db.Create(&convert).Error
	if err != nil {
		log.Println("add book query error", err.Error())
		return post.Core{}, err
	}

	newPost.ID = convert.ID
	newPost.UserID = convert.UserID

	return newPost, nil
}

func (pd *postData) Update(postID uint, userID uint, updatedPost post.Core) (post.Core, error) {
	convert := CoreToData(updatedPost)
	qry := pd.db.Where("id = ? AND user_id = ?", postID, userID).Updates(&convert)
	if qry.RowsAffected <= 0 {
		log.Println("update post query error : data not found")
		return post.Core{}, errors.New("not found")
	}

	if err := qry.Error; err != nil {
		log.Println("update post query error :", err.Error())
		return post.Core{}, errors.New("not found")
	}
	log.Println("update di query", convert.ID)
	return DataToCore(convert), nil
}

func (pd *postData) Delete(postID uint, userID uint) error {
	qry := pd.db.Where("user_id = ?", userID).Delete(&Post{}, postID)

	affrows := qry.RowsAffected

	if affrows == 0 {
		log.Println("no rows affected")
		return errors.New("data not found")
	}

	err := qry.Error
	if err != nil {
		log.Println("delete post query error")
		return errors.New("data not found")
	}

	return nil
}

func (pd *postData) MyPosts(userID uint) ([]post.MyPostsResp, error) {
	myPosts := []post.Core{}
	cnvMyPosts := []post.MyPostsResp{}
	comments := []comment.Core{}

	err := pd.db.Raw("SELECT posts.id, posts.content, posts.photo, posts.user_id, name Writer, posts.created_at  FROM posts JOIN users u ON user_id = u.id  WHERE posts.deleted_at is NULL AND user_id = ?", userID).Scan(&myPosts).Error
	if err != nil {
		log.Println("my book query error")
		return []post.MyPostsResp{}, errors.New("data not found")
	}

	for i, _ := range myPosts {
		tmp := post.ToMyPostResp(myPosts[i])
		cnvMyPosts = append(cnvMyPosts, tmp)
	}

	for i, val := range cnvMyPosts {
		pd.db.Raw("SELECT c.id, user_id, name UserName, post_id, c.created_at, content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.deleted_at IS NULL AND post_id = ?", val.ID).Scan(&comments)
		cnvMyPosts[i].Comments = append(val.Comments, comments...)
	}

	return cnvMyPosts, nil
}

func (pd *postData) AllPosts() ([]post.MyPostsResp, error) {
	allPosts := []post.Core{}
	cnvAllPosts := []post.MyPostsResp{}
	comments := []comment.Core{}

	err := pd.db.Raw("SELECT posts.id, posts.content, posts.photo, posts.user_id, name Writer, posts.created_at FROM posts JOIN users u ON user_id = u.id  WHERE posts.deleted_at is NULL").Scan(&allPosts).Error
	if err != nil {
		log.Println("all book query error")
		return []post.MyPostsResp{}, err
	}

	for i, _ := range allPosts {
		tmp := post.ToMyPostResp(allPosts[i])
		cnvAllPosts = append(cnvAllPosts, tmp)
	}

	for i, val := range cnvAllPosts {
		pd.db.Raw("SELECT c.id, user_id, name UserName, post_id, c.created_at, content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.deleted_at IS NULL AND post_id = ?", val.ID).Scan(&comments)
		cnvAllPosts[i].Comments = append(val.Comments, comments...)
	}

	return cnvAllPosts, nil
}

func (pd *postData) GetPostById(postID uint, userID uint) (post.MyPostsResp, error) {
	res := post.Core{}
	comments := []comment.Core{}

	err := pd.db.Raw("SELECT posts.id, posts.content, posts.photo, posts.user_id, name Writer, posts.created_at FROM posts JOIN users u ON user_id = u.id  WHERE posts.deleted_at is NULL AND posts.id = ?", postID).Scan(&res).Error
	if err != nil {
		log.Println("GetPostById query error")
		return post.MyPostsResp{}, err
	}

	cnvRes := post.ToMyPostResp(res)
	pd.db.Raw("SELECT c.id, user_id, name UserName, post_id, c.created_at, content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.deleted_at IS NULL AND post_id = ?", postID).Scan(&comments)
	cnvRes.Comments = append(cnvRes.Comments, comments...)

	return cnvRes, nil
}
