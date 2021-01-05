package view

import (
	"github.com/caos/zitadel/internal/eventstore/models"
	"github.com/caos/zitadel/internal/org/model"
	org_view "github.com/caos/zitadel/internal/org/repository/view"
	org_model "github.com/caos/zitadel/internal/org/repository/view/model"
	"github.com/caos/zitadel/internal/view/repository"
)

const (
	orgTable = "auth.orgs"
)

func (v *View) OrgByID(orgID string) (*org_model.OrgView, error) {
	return org_view.OrgByID(v.Db, orgTable, orgID)
}

func (v *View) OrgByPrimaryDomain(primaryDomain string) (*org_model.OrgView, error) {
	return org_view.OrgByPrimaryDomain(v.Db, orgTable, primaryDomain)
}

func (v *View) SearchOrgs(req *model.OrgSearchRequest) ([]*org_model.OrgView, uint64, error) {
	return org_view.SearchOrgs(v.Db, orgTable, req)
}

func (v *View) PutOrg(org *org_model.OrgView, event *models.Event) error {
	err := org_view.PutOrg(v.Db, orgTable, org)
	if err != nil {
		return err
	}
	return v.ProcessedOrgSequence(event)
}

func (v *View) GetLatestOrgFailedEvent(sequence uint64) (*repository.FailedEvent, error) {
	return v.latestFailedEvent(orgTable, sequence)
}

func (v *View) ProcessedOrgFailedEvent(failedEvent *repository.FailedEvent) error {
	return v.saveFailedEvent(failedEvent)
}

func (v *View) UpdateOrgSpoolerRunTimestamp() error {
	return v.updateSpoolerRunSequence(orgTable)
}

func (v *View) GetLatestOrgSequence(aggregateType string) (*repository.CurrentSequence, error) {
	return v.latestSequence(orgTable, aggregateType)
}

func (v *View) ProcessedOrgSequence(event *models.Event) error {
	return v.saveCurrentSequence(orgTable, event)
}
