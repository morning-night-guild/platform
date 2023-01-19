package controller

import (
	"encoding/base64"
	"strconv"

	"github.com/morning-night-guild/platform/internal/domain/repository"
)

// トークン.
type Token string

// NewToken トークンを生成する.
func NewToken(value string) Token {
	return Token(value)
}

// CreateTokenFromIndex インデックスからトークンを作成する.
func CreateTokenFromIndex(index repository.Index) Token {
	return Token(base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(index.Int()))))
}

// String トークンの文字列を提供する.
func (t Token) String() string {
	return string(t)
}

// ToIndex トークンからインデックスを作成する.
func (t Token) ToIndex() repository.Index {
	dec, err := base64.StdEncoding.DecodeString(t.String())
	if err != nil {
		return repository.Index(0)
	}

	index, err := strconv.Atoi(string(dec))
	if err != nil {
		return repository.Index(0)
	}

	return repository.Index(index)
}

// CreateNextToken ネクストトークンを作成する.
func (t Token) CreateNextToken(size repository.Size) Token {
	return Token(base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(t.ToIndex().Int() + size.Int()))))
}
