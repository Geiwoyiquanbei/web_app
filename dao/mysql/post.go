package mysql

import (
	"web_app/models"
)

func CreatePost(p *models.Post) {
	sqlStr := `insert into post (post_id,title ,content,author_id,community_id) values (?,?,?,?,?)`
	db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
}
func GetPostDetail(id int64) (data *models.Post, err error) {
	sqlStr := `select post_id ,title ,content ,author_id ,community_id,create_time from post where post_id= ? `
	data = new(models.Post)
	err = db.Get(data, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func GetPostList(limit, offset int64) (data []*models.Post, err error) {
	//limit 条数  offset 页数
	sqlStr := `select post_id ,title ,content ,author_id ,community_id,create_time from post limit ?,?`
	data = make([]*models.Post, 0, 2)
	db.Select(&data, sqlStr, (offset-1)*limit, limit)
	return data, nil
}
