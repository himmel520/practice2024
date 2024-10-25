package usecase

type (
	Usecase struct {
		Keyword KeywordUc
	}

	KeywordUc interface {
		GetMapping() map[string][]string
	}
)

func New() *Usecase {
	return &Usecase{
		Keyword: NewKeywordUsecase(),
	}
}
