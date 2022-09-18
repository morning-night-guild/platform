package article

import (
	"fmt"
	"net/url"
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
func NewURL(value string) (URL, error) {
	u := URL(value)

	if err := u.validate(); err != nil {
		return URL(""), err
	}

	return u, nil
}

// validate URLを検証するメソッド.
func (u URL) validate() error {
	if _, err := url.Parse(u.String()); err != nil {
		return err
	}

	if !strings.HasPrefix(u.String(), "https://") {
		msg := fmt.Sprintf("url must be start with `https://`. value is %s", u.String())

		return errors.NewValidationError(msg)
	}

	return nil
}
