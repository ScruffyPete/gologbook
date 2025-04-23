package domain

import "context"

type UnitOfWork interface {
	WithTx(ctx context.Context, fn func(repos RepoBundle) error) error
}
