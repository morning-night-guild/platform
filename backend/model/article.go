package model

import "github.com/morning-night-guild/platform/model/article"

// Article 記事モデル.
type Article struct {
	ID        article.ID        // ID
	Title     article.Title     // タイトル
	URL       article.URL       // 記事のURL
	Thumbnail article.Thumbnail // サムネイル
	TagList   article.TagList   // タグリスト
}

// NewArticle 記事モデルのファクトリー関数.
func NewArticle(
	id article.ID,
	title article.Title,
	url article.URL,
	thumbnail article.Thumbnail,
	tags article.TagList,
) (Article, error) {
	article := Article{
		ID:        id,
		Title:     title,
		URL:       url,
		Thumbnail: thumbnail,
		TagList:   tags,
	}

	if err := article.validate(); err != nil {
		return Article{}, err
	}

	return article, nil
}

// validate 記事を検証するメソッド.
func (a Article) validate() error {
	return nil
}

// CreateArticle 記事モデルを新規作成する関数.
func CreateArticle(
	title article.Title,
	url article.URL,
	thumbnail article.Thumbnail,
	tags article.TagList,
) Article {
	id := article.GenerateID()

	return Article{
		ID:        id,
		Title:     title,
		URL:       url,
		Thumbnail: thumbnail,
		TagList:   tags,
	}
}
