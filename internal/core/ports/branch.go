package ports

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type Branch interface {
	CreateBranch(domain.Branch) (*domain.Branch, error)
}
