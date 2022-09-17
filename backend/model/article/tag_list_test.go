package article_test

import (
	"reflect"
	"testing"

	"github.com/morning-night-guild/platform/model/article"
)

func TestNewTagList(t *testing.T) {
	t.Parallel()

	type args struct {
		value []article.Tag
	}

	tests := []struct {
		name    string
		args    args
		want    article.TagList
		wantErr bool
	}{
		{
			name: "5個のタグを含むタグリストが作成できる",
			args: args{
				value: []article.Tag{
					article.Tag("tag1"),
					article.Tag("tag2"),
					article.Tag("tag3"),
					article.Tag("tag4"),
					article.Tag("tag5"),
				},
			},
			want: article.TagList(
				[]article.Tag{
					article.Tag("tag1"),
					article.Tag("tag2"),
					article.Tag("tag3"),
					article.Tag("tag4"),
					article.Tag("tag5"),
				},
			),
			wantErr: false,
		},
		{
			name: "重複したタグは排除してタグリストが作成できる",
			args: args{
				value: []article.Tag{
					article.Tag("tag"),
					article.Tag("tag"),
					article.Tag("tag"),
					article.Tag("tag"),
					article.Tag("tag"),
				},
			},
			want: article.TagList(
				[]article.Tag{
					article.Tag("tag"),
				},
			),
			wantErr: false,
		},
		{
			name: "空タグリストが作成できる",
			args: args{
				value: []article.Tag{},
			},
			want: article.TagList(
				[]article.Tag{},
			),
			wantErr: false,
		},
		{
			name: "6個のタグを含むタグリストは作成に失敗する",
			args: args{
				value: []article.Tag{
					article.Tag("tag1"),
					article.Tag("tag2"),
					article.Tag("tag3"),
					article.Tag("tag4"),
					article.Tag("tag5"),
					article.Tag("tag6"),
				},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := article.NewTagList(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTagList() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTagList() = %v, want %v", got, tt.want)
			}
		})
	}
}
