package mock

import (
	"context"

	platform "github.com/influxdata/influxdb"
)

var _ platform.OrganizationService = &OrganizationService{}

// OrganizationService is a mock organization server.
type OrganizationService struct {
	FindOrganizationByIDF func(ctx context.Context, id platform.ID) (*platform.Organization, error)
	FindOrganizationF     func(ctx context.Context, filter platform.OrganizationFilter) (*platform.Organization, error)
	FindOrganizationsF    func(ctx context.Context, filter platform.OrganizationFilter, opt ...platform.FindOptions) ([]*platform.Organization, int, error)
	CreateOrganizationF   func(ctx context.Context, b *platform.Organization) error
	UpdateOrganizationF   func(ctx context.Context, id platform.ID, upd platform.OrganizationUpdate) (*platform.Organization, error)
	UpdateOrgLimitsF      func(ctx context.Context, orgID platform.ID, l platform.OrgLimits) (*platform.Organization, error)
	DeleteOrganizationF   func(ctx context.Context, id platform.ID) error
}

// NewOrganizationService returns a mock OrganizationService where its methods will return
// zero values.
func NewOrganizationService() *OrganizationService {
	return &OrganizationService{
		FindOrganizationByIDF: func(ctx context.Context, id platform.ID) (*platform.Organization, error) { return nil, nil },
		FindOrganizationF: func(ctx context.Context, filter platform.OrganizationFilter) (*platform.Organization, error) {
			return nil, nil
		},
		FindOrganizationsF: func(ctx context.Context, filter platform.OrganizationFilter, opt ...platform.FindOptions) ([]*platform.Organization, int, error) {
			return nil, 0, nil
		},
		CreateOrganizationF: func(ctx context.Context, b *platform.Organization) error { return nil },
		UpdateOrganizationF: func(ctx context.Context, id platform.ID, upd platform.OrganizationUpdate) (*platform.Organization, error) {
			return nil, nil
		},
		DeleteOrganizationF: func(ctx context.Context, id platform.ID) error { return nil },
	}
}

//FindOrganizationByID calls FindOrganizationByIDF.
func (s *OrganizationService) FindOrganizationByID(ctx context.Context, id platform.ID) (*platform.Organization, error) {
	return s.FindOrganizationByIDF(ctx, id)
}

//FindOrganization calls FindOrganizationF.
func (s *OrganizationService) FindOrganization(ctx context.Context, filter platform.OrganizationFilter) (*platform.Organization, error) {
	return s.FindOrganizationF(ctx, filter)
}

//FindOrganizations calls FindOrganizationsF.
func (s *OrganizationService) FindOrganizations(ctx context.Context, filter platform.OrganizationFilter, opt ...platform.FindOptions) ([]*platform.Organization, int, error) {
	return s.FindOrganizationsF(ctx, filter, opt...)
}

// CreateOrganization calls CreateOrganizationF.
func (s *OrganizationService) CreateOrganization(ctx context.Context, b *platform.Organization) error {
	return s.CreateOrganizationF(ctx, b)
}

// UpdateOrganization calls UpdateOrganizationF.
func (s *OrganizationService) UpdateOrganization(ctx context.Context, id platform.ID, upd platform.OrganizationUpdate) (*platform.Organization, error) {
	return s.UpdateOrganizationF(ctx, id, upd)
}

// DeleteOrganization calls DeleteOrganizationF.
func (s *OrganizationService) DeleteOrganization(ctx context.Context, id platform.ID) error {
	return s.DeleteOrganizationF(ctx, id)
}

// UpdateOrgLimits calls UpdateOrgLimitsF.
func (s *OrganizationService) UpdateOrgLimits(ctx context.Context, orgID platform.ID, l platform.OrgLimits) (*platform.Organization, error) {
	return s.UpdateOrgLimitsF(ctx, orgID, l)
}
