package article

import (
	"fmt"

	"github.com/morning-night-guild/platform/model/errors"
)

// Tag タグリスト.
type TagList []Tag

// maxTagLength タグリストに含まれるタグの個数の最大値.
const maxTagLength = 5

// NewTagList タグリストを新規作成する関数.
func NewTagList(value []Tag) (TagList, error) {
	tags := TagList(value)

	if err := tags.validate(); err != nil {
		return nil, err
	}

	return tags, nil
}

// validate タグリストを検証するメソッド.
func (t TagList) validate() error {
	if t.Len() > maxTagLength {
		msg := fmt.Sprintf("must be less than or equal to %d. length is %d", maxTagLength, t.Len())

		return errors.NewValidationError(msg)
	}

	return nil
}

// Len タグリストに含まれるのタグの個数を提供するメソッド.
func (t TagList) Len() int {
	return len(t)
}
