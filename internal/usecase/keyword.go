package usecase

import "github.com/himmel520/practice2024/internal/infrastucture/storage"

type KeywordUsecase struct {
	keywordsMapping map[string][]string
}

func NewKeywordUsecase() *KeywordUsecase{
	return &KeywordUsecase{
		keywordsMapping: storage.KeywordsMapping,
	}
}

func(u *KeywordUsecase) GetMapping() map[string][]string {
	return u.keywordsMapping
}
