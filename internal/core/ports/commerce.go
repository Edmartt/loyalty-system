package ports

import "github.com/Edmartt/loyalty-system/internal/core/domain"

type Commerce interface {
	CreateCommerce(domain.Commerce) (*domain.Commerce, error)
}
