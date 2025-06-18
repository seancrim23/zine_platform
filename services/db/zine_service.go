package services

import (
	"context"
	"zine_platform/models"
)

// This interface will encompass all actions of loading Zine data into a database
type ZinePlatformZineService interface {
	GetZine(context.Context, string) (*models.Zine, error)
	CreateZine(context.Context, models.Zine) (*models.Zine, error)
	UpdateZine(context.Context, string, models.Zine) (*models.Zine, error)
	DeleteZine(context.Context, string) error
}
