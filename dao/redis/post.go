package redis

import (
	"web_app/models"
)

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	//2.确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//3. zrevrange  按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(key, start, end).Result()
}
