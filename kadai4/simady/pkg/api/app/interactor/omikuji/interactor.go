package interactor

import (
	"context"

	"omikuji-app/pkg/api/app/presenter"
	service "omikuji-app/pkg/api/domain/service/omikuji"
)

type OmikujiInteractor interface {
	Draw(ctx context.Context) (string, error)
}

type omikujiInteractor struct {
	presenter      presenter.Presenter
	omikujiService service.OmikujiService
}

func New(p presenter.Presenter, s service.OmikujiService) OmikujiInteractor {
	return &omikujiInteractor{presenter: p, omikujiService: s}
}

func (i *omikujiInteractor) Draw(ctx context.Context) (string, error) {
	rs := i.omikujiService.Draw(ctx)
	output, err := i.presenter.Output(rs)
	if err != nil {
		return "", err
	}
	return output, nil
}
