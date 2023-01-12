package controller

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/morning-night-guild/platform/internal/model/article"
	"github.com/morning-night-guild/platform/internal/usecase/port"
	"github.com/morning-night-guild/platform/internal/usecase/repository"
	articlev1 "github.com/morning-night-guild/platform/pkg/connect/proto/article/v1"
)

// Article.
type Article struct {
	key   string
	share port.ShareArticle
	list  port.ListArticle
}

// NewArticle 記事のコントローラを新規作成する関数.
func NewArticle(
	key string,
	share port.ShareArticle,
	list port.ListArticle,
) *Article {
	return &Article{
		key:   key,
		share: share,
		list:  list,
	}
}

// Share 記事を共有するコントローラメソッド.
func (a *Article) Share(
	ctx context.Context,
	req *connect.Request[articlev1.ShareRequest],
) (*connect.Response[articlev1.ShareResponse], error) {
	if req.Header().Get("X-Api-Key") != a.key {
		return nil, ErrUnauthorized
	}

	url, err := article.NewURL(req.Msg.Url)
	if err != nil {
		return nil, handleError(ctx, err)
	}

	title, err := article.NewTitle(req.Msg.Title)
	if err != nil {
		return nil, handleError(ctx, err)
	}

	description, err := article.NewDescription(req.Msg.Description)
	if err != nil {
		return nil, handleError(ctx, err)
	}

	thumbnail, err := article.NewThumbnail(req.Msg.Thumbnail)
	if err != nil {
		return nil, handleError(ctx, err)
	}

	input := port.ShareArticleInput{
		URL:         url,
		Title:       title,
		Description: description,
		Thumbnail:   thumbnail,
	}

	output, err := a.share.Execute(ctx, input)
	if err != nil {
		return nil, handleError(ctx, err)
	}

	return connect.NewResponse(&articlev1.ShareResponse{
		Article: &articlev1.Article{
			Id:          output.Article.ID.String(),
			Title:       output.Article.Title.String(),
			Url:         output.Article.URL.String(),
			Description: output.Article.Description.String(),
			Thumbnail:   output.Article.Thumbnail.String(),
			Tags:        output.Article.TagList.StringSlice(),
		},
	}), nil
}

// List 記事を取得するコントローラメソッド.
func (a *Article) List(
	ctx context.Context,
	req *connect.Request[articlev1.ListRequest],
) (*connect.Response[articlev1.ListResponse], error) {
	token := NewToken(req.Msg.PageToken)

	index := token.ToIndex()

	size, err := repository.NewSize(int(req.Msg.MaxPageSize))
	if err != nil {
		return nil, handleError(ctx, err)
	}

	input := port.ListArticleInput{
		Index: index,
		Size:  size,
	}

	output, err := a.list.Execute(ctx, input)
	if err != nil {
		return nil, handleError(ctx, err)
	}

	result := make([]*articlev1.Article, len(output.Articles))

	for i, article := range output.Articles {
		result[i] = &articlev1.Article{
			Id:          article.ID.String(),
			Title:       article.Title.String(),
			Url:         article.URL.String(),
			Description: article.Description.String(),
			Thumbnail:   article.Thumbnail.String(),
			Tags:        article.TagList.StringSlice(),
		}
	}

	next := token.CreateNextToken(size).String()
	if len(output.Articles) < size.Int() {
		next = ""
	}

	return connect.NewResponse(&articlev1.ListResponse{
		Articles:      result,
		NextPageToken: next,
	}), nil
}
