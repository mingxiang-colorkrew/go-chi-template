package tenantfactory

import (
	"go_chi_template/config"
	"go_chi_template/db/go_chi_template/public/model"
	"go_chi_template/src/mutation"
)

func SeedTenant(app *config.App, opts ...TenantOption) (*model.Tenant, error) {
	tenant := model.Tenant{
		Name:      "Default Tenant Name",
		ShortCode: "aaay1",
	}

	for _, opt := range opts {
		opt(&tenant)
	}

	inserted, err := mutation.InsertTenant(app, &tenant)

	return inserted, err
}

type TenantOption func(*model.Tenant)

func WithName(name string) TenantOption {
	return func(t *model.Tenant) {
		t.Name = name
	}
}

func WithShortCode(shortCode string) TenantOption {
	return func(t *model.Tenant) {
		t.ShortCode = shortCode
	}
}
