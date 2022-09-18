package gateway

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/morning-night-guild/platform/model"
	"github.com/morning-night-guild/platform/model/article"
	"github.com/morning-night-guild/platform/usecase/repository"
)

var _ repository.OGP = (*OGP)(nil)

// OGP.
type OGP struct {
	client *http.Client
}

// NewOGP OGPゲートウェイを新規作成する関数.
func NewOGP() *OGP {
	return &OGP{
		client: http.DefaultClient,
	}
}

// Create urlからarticleモデルを作成するメソッド.
func (o *OGP) Create(ctx context.Context, url article.URL) (model.Article, error) {
	gr, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		//
		return model.Article{}, err
	}

	res, err := o.client.Do(gr.WithContext(ctx))
	if err != nil {
		//
		return model.Article{}, err
	}

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	og := opengraph.NewOpenGraph()

	err = og.ProcessHTML(strings.NewReader(string(body)))
	if err != nil {
		//
		return model.Article{}, err
	}

	title, err := article.NewTitle(og.Title)
	if err != nil {
		return model.Article{}, err
	}

	imageURL := ""
	if len(og.Images) > 0 {
		imageURL = og.Images[0].URL
	}

	thumbnail, err := article.NewThumbnail(imageURL)
	if err != nil {
		return model.Article{}, err
	}

	return model.CreateArticle(
		title,
		url,
		thumbnail,
		article.TagList{},
	), nil
}
