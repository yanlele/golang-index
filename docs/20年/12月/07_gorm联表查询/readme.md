## gorm联表查询

```go
/*
获取单个文章

实体关联：https://gorm.io/zh_CN/docs/associations.html
*/
func GetArticle(id int) (*Article, error) {
	var article Article

	// 方式一 关联查询
	err := db.Where("id = ?", id).First(&article).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&article).Association("Tag").Find(&article.Tag)
	if err != nil {
		return nil, err
	}

	// 方式二 分别写两个sql 就完事儿
	//db.Where("id = ?", id).First(&article)
	//db.Where("id = ?", article.TagID).First(&article.Tag)

	// 方式三
	//err := db.Preload("Tag").Where("id = ?", id).Find(&article).Error
	return &article, err
}
```

