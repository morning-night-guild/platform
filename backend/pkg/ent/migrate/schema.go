// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ArticlesColumns holds the columns for the "articles" table.
	ArticlesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "title", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "description", Type: field.TypeString},
		{Name: "thumbnail", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// ArticlesTable holds the schema information for the "articles" table.
	ArticlesTable = &schema.Table{
		Name:       "articles",
		Columns:    ArticlesColumns,
		PrimaryKey: []*schema.Column{ArticlesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "article_title",
				Unique:  false,
				Columns: []*schema.Column{ArticlesColumns[1]},
			},
		},
	}
	// ArticleTagsColumns holds the columns for the "article_tags" table.
	ArticleTagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "tag", Type: field.TypeString},
		{Name: "article_id", Type: field.TypeUUID},
	}
	// ArticleTagsTable holds the schema information for the "article_tags" table.
	ArticleTagsTable = &schema.Table{
		Name:       "article_tags",
		Columns:    ArticleTagsColumns,
		PrimaryKey: []*schema.Column{ArticleTagsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "article_tags_articles_tags",
				Columns:    []*schema.Column{ArticleTagsColumns[2]},
				RefColumns: []*schema.Column{ArticlesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "articletag_tag_article_id",
				Unique:  true,
				Columns: []*schema.Column{ArticleTagsColumns[1], ArticleTagsColumns[2]},
			},
			{
				Name:    "articletag_tag",
				Unique:  false,
				Columns: []*schema.Column{ArticleTagsColumns[1]},
			},
			{
				Name:    "articletag_article_id",
				Unique:  false,
				Columns: []*schema.Column{ArticleTagsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ArticlesTable,
		ArticleTagsTable,
	}
)

func init() {
	ArticleTagsTable.ForeignKeys[0].RefTable = ArticlesTable
}
