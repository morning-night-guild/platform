package article

import (
	"fmt"
	"strings"

	"github.com/morning-night-guild/platform/model/errors"
)

// URL 記事のURL.
type URL string

// String URLを文字列として提供するメソッド.
func (u URL) String() string {
	return string(u)
}

// NewURL URLを新規作成する関数.
func NewURL(u string) (URL, error) {
	url := URL(u)

	if err := url.validate(); err != nil {
		return URL(""), err
	}

	return url, nil
}

// validate URLを検証するメソッド.
func (u URL) validate() error {
	if !strings.HasPrefix(u.String(), "https://") {
		msg := fmt.Sprintf("url must be start with `https://`. value is %s", u.String())

		return errors.NewValidationError(msg)
	}

	return nil
}
