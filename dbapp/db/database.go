package db

import (
	"context"
	"dbapp/ent"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type IDatabase interface {
	CreateApplicationCategory(name string, displayOrder int) (*ent.ApplicationCategory, error)
	GetAllApplicationCategories() []*ent.ApplicationCategory
	UpdateApplicationCategory(applicationCategory *ent.ApplicationCategory) (*ent.ApplicationCategory, error)
	DeleteApplicationCategory(applicationCategory *ent.ApplicationCategory) error
	CreateApplication(name string, description string, altText string, uri string, iconURI string, isFavourite bool, validFrom time.Time, validTo time.Time, applicationCategory *ent.ApplicationCategory) (*ent.Application, error)
	GetAllApplications() []*ent.Application
	UpdateApplication(application *ent.Application) (*ent.Application, error)
	DeleteApplication(application *ent.Application) error
	CreateADGroup(name string) (*ent.ADGroup, error)
	GetAllADGroups() []*ent.ADGroup
	UpdateADGroup(adGroup *ent.ADGroup) (*ent.ADGroup, error)
	DeleteADGroup(adGroup *ent.ADGroup) error
	CreateCommunity(whamSiteID int, whamTitle string, whanDescription string, whamCommunityURL string, whamCreated time.Time, whamUpdated time.Time, featuredFrom time.Time, featuredTo time.Time, adGroups []*ent.ADGroup, communityCategories []*ent.CommunityCategory) (*ent.Community, error)
	GetAllCommunities() []*ent.Community
	UpdateCommunity(community *ent.Community) (*ent.Community, error)
	DeleteCommunity(community *ent.Community) error
	CreateCommunityCategory(name string, displayOrder int) (*ent.CommunityCategory, error)
	GetAllCommunityCategories() []*ent.CommunityCategory
	UpdateCommunityCategory(communityCategory *ent.CommunityCategory) (*ent.CommunityCategory, error)
	DeleteCommunityCategory(communityCategory *ent.CommunityCategory) error
	CreatePartner(whamSiteID int, whamTitle string, whamDescription string, KeycloakOrganisation string, whamPartnerURL string, whamCreated time.Time, whamUpdated time.Time, applications []*ent.Application) (*ent.Partner, error)
	GetAllPartners() []*ent.Partner
	UpdatePartner(partner *ent.Partner) (*ent.Partner, error)
	DeletePartner(partner *ent.Partner) error
	CreatePartnerPageLinkFragment(linkText string, whamPartnerURLSuffix string, displayOrder int) (*ent.PartnerPageLinkFragment, error)
	GetAllPartnerPageLinkFragments() []*ent.PartnerPageLinkFragment
}

type Database struct {
	client *ent.Client
}

func Init() *Database {
	dbuser := "appuser"
	dbpass := "pass"
	dbhost := "localhost"
	dbport := "3306"
	dbname := "appdb"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbuser, dbpass, dbhost, dbport, dbname)
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	return &Database{client: client}

}
func (d *Database) CreateApplicationCategory(name string, displayOrder int) (*ent.ApplicationCategory, error) {
	createdApplicationCategory, err := d.client.ApplicationCategory.Create().
		SetName(name).SetDisplayOrder(displayOrder).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating application category: %v\n", err)
		return nil, fmt.Errorf("failed creating application category: %v", err)
	}
	return createdApplicationCategory, nil

}
func (d *Database) GetAllApplicationCategories() []*ent.ApplicationCategory {
	applicationCategories, err := d.client.ApplicationCategory.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all application categories: %v\n", err)
	}
	// fmt.Printf("Application Categories: %+v\n", applicationCategories)
	return applicationCategories
}
func (d *Database) CreateApplication(name string, description string, altText string, uri string, iconURI string,
	isFavourite bool, validFrom time.Time, validTo time.Time, applicationCategory *ent.ApplicationCategory) (*ent.Application, error) {
	application, err := d.client.Application.Create().
		SetName(name).
		SetDescription(description).
		SetAltText(altText).
		SetURI(uri).
		SetIconURI(iconURI).
		SetIsFavourite(isFavourite).
		SetValidFrom(validFrom).
		SetValidTo(validTo).
		SetApplicationCategory(applicationCategory).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating application: %v\n", err)
		return nil, fmt.Errorf("failed creating application category: %v", err)
	}
	// fmt.Printf("Application: %+v\n", application)
	return application, nil
}
func (d *Database) GetAllApplications() []*ent.Application {
	applications, err := d.client.Application.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all applications: %v\n", err)
	}
	// fmt.Printf("Applications: %+v\n", applications)
	return applications
}
func (d *Database) CreateADGroup(name string) (*ent.ADGroup, error) {
	adGroup, err := d.client.ADGroup.Create().
		SetName(name).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating ad group: %v\n", err)
		return nil, fmt.Errorf("failed creating application category: %v", err)
	}
	// fmt.Printf("ADGroup: %+v\n", adGroup)
	return adGroup, nil
}
func (d *Database) GetAllADGroups() []*ent.ADGroup {
	adGroups, err := d.client.ADGroup.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all ad groups: %v\n", err)
	}
	// fmt.Printf("ADGroups: %+v\n", adGroups)
	return adGroups
}
func (d *Database) CreateCommunity(whamSiteID int, whamTitle string, whamDescription string,
	whamCommunityURL string, whamCreated time.Time, whamUpdated time.Time, featuredFrom time.Time,
	featuredTo time.Time, adGroups []*ent.ADGroup, communityCategories []*ent.CommunityCategory) (*ent.Community, error) {
	community, err := d.client.Community.Create().
		SetWhamSiteID(whamSiteID).
		SetWhamTitle(whamTitle).
		SetWhanDescription(whamDescription).
		SetWhamCommunityURL("Test Community").
		SetWhamCreated(time.Now()).
		SetWhamUpdated(time.Now()).
		SetFeaturedFrom(time.Now()).
		SetFeaturedTo(time.Now()).
		AddAdgroup(adGroups...).
		AddCommunityCategory(communityCategories...).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating community: %v\n", err)
		return nil, fmt.Errorf("failed creating community: %v", err)
	}
	// fmt.Printf("Community: %+v\n", community)
	return community, nil
}
func (d *Database) GetAllCommunities() []*ent.Community {
	communities, err := d.client.Community.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all communities: %v\n", err)
	}
	// fmt.Printf("Communities: %+v\n", communities)
	return communities
}
func (d *Database) CreatePartner(WhamSiteID int, WhamTitle string, WhamDescription string, KeycloakOrganisation string,
	WhamPartnerURL string, WhamCreated time.Time, WhamUpdated time.Time, applications []*ent.Application) (*ent.Partner, error) {
	partner, err := d.client.Partner.Create().
		SetWhamSiteID(WhamSiteID).
		SetWhamTitle(WhamTitle).
		SetWhamDescription(WhamDescription).
		SetKeycloakOrganisation(KeycloakOrganisation).
		SetWhamPartnerURL(WhamPartnerURL).
		SetWhamCreated(WhamCreated).
		SetWhamUpdated(WhamUpdated).
		AddApplication(applications...).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating partner: %v\n", err)
		return nil, fmt.Errorf("failed creating partner: %v", err)
	}
	// fmt.Printf("Partner: %+v\n", partner)
	return partner, nil
}
func (d *Database) GetAllPartners() []*ent.Partner {
	partners, err := d.client.Partner.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all partners: %v\n", err)
	}
	// fmt.Printf("Partners: %+v\n", partners)
	return partners
}

func (d *Database) CreatePartnerPageLinkFragment(linkText string, whamPartnerURLSuffix string, displayOrder int) (*ent.PartnerPageLinkFragment, error) {
	partnerPageLinkFragment, err := d.client.PartnerPageLinkFragment.Create().
		SetLinkText(linkText).
		SetWhamPartnerURLSuffix(whamPartnerURLSuffix).
		SetDisplayOrder(displayOrder).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating partner page link fragment: %v\n", err)
		return nil, fmt.Errorf("failed creating partner page link fragment: %v", err)
	}
	// fmt.Printf("Partner Page Link Fragment: %+v\n", partnerPageLinkFragment)
	return partnerPageLinkFragment, nil
}

func (d *Database) GetAllPartnerPageLinkFragments() []*ent.PartnerPageLinkFragment {
	partnerPageLinkFragments, err := d.client.PartnerPageLinkFragment.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all partner page link fragments: %v\n", err)
	}
	// fmt.Printf("Partner Page Link Fragments: %+v\n", partnerPageLinkFragments)
	return partnerPageLinkFragments
}
func (d *Database) CreateCommunityCategory(name string, displayOrder int) (*ent.CommunityCategory, error) {
	communityCategory, err := d.client.CommunityCategory.Create().
		SetName(name).
		SetDisplayOrder(displayOrder).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed creating community category: %v\n", err)
		return nil, fmt.Errorf("failed creating community category: %v", err)
	}
	// fmt.Printf("Community Category: %+v\n", communityCategory)
	return communityCategory, nil
}

func (d *Database) GetAllCommunityCategories() []*ent.CommunityCategory {
	communityCategories, err := d.client.CommunityCategory.Query().All(context.Background())
	if err != nil {
		fmt.Printf("failed getting all community categories: %v\n", err)
	}
	// fmt.Printf("Community Categories: %+v\n", communityCategories)
	return communityCategories
}

func (d *Database) UpdateApplicationCategory(applicationCategory *ent.ApplicationCategory) (*ent.ApplicationCategory, error) {
	foundApplicationCategory, err := d.client.ApplicationCategory.Get(context.Background(), applicationCategory.ID)
	if err != nil {
		fmt.Printf("failed updating application category: %v\n", err)
		return nil, fmt.Errorf("failed updating application category: %v", err)
	}
	updatedApplicationCategory, err := foundApplicationCategory.Update().
		SetName(applicationCategory.Name).
		SetDisplayOrder(applicationCategory.DisplayOrder).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed updating application category: %v\n", err)
		return nil, fmt.Errorf("failed updating application category: %v", err)
	}
	// fmt.Printf("Application Category was updated from: %+v\n to: %+v\n", spew.Sdump(*foundApplicationCategory), spew.Sdump(*updatedApplicationCategory))
	return updatedApplicationCategory, nil
}
func (d *Database) DeleteApplicationCategory(applicationCategory *ent.ApplicationCategory) error {
	foundApplicationCategory, err := d.client.ApplicationCategory.Get(context.Background(), applicationCategory.ID)
	if err != nil {
		fmt.Printf("failed deleting application category: %v\n", err)
		return fmt.Errorf("failed deleting application category: %v", err)
	}
	err = d.client.ApplicationCategory.DeleteOne(foundApplicationCategory).Exec(context.Background())
	if err != nil {
		fmt.Printf("failed deleting application category: %v\n", err)
		return fmt.Errorf("failed deleting application category: %v", err)
	}
	// fmt.Printf("Application Category was deleted: %+v\n", spew.Sdump(*foundApplicationCategory))
	return nil
}
func (d *Database) UpdateApplication(application *ent.Application) (*ent.Application, error) {
	foundApplication, err := d.client.Application.Get(context.Background(), application.ID)
	if err != nil {
		fmt.Printf("failed updating application: %v\n", err)
		return nil, fmt.Errorf("failed updating application: %v", err)
	}
	updatedApplication, err := foundApplication.Update().
		SetName(application.Name).
		SetDescription(application.Description).
		SetAltText(application.AltText).
		SetURI(application.URI).
		SetIconURI(application.IconURI).
		SetIsFavourite(application.IsFavourite).
		SetValidFrom(application.ValidFrom).
		SetValidTo(application.ValidTo).
		SetApplicationCategory(application.Edges.ApplicationCategory).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed updating application: %v\n", err)
		return nil, fmt.Errorf("failed updating application: %v", err)
	}
	// fmt.Printf("Application was updated from: %+v\n to: %+v\n", spew.Sdump(*foundApplication), spew.Sdump(*updatedApplication))
	return updatedApplication, nil
}

func (d *Database) DeleteApplication(application *ent.Application) error {
	foundApplication, err := d.client.Application.Get(context.Background(), application.ID)
	if err != nil {
		fmt.Printf("failed deleting application: %v\n", err)
		return fmt.Errorf("failed deleting application: %v", err)
	}
	err = d.client.Application.DeleteOne(foundApplication).Exec(context.Background())
	if err != nil {
		fmt.Printf("failed deleting application: %v\n", err)
		return fmt.Errorf("failed deleting application: %v", err)
	}
	// fmt.Printf("Application was deleted: %+v\n", spew.Sdump(*foundApplication))
	return nil
}
func (d *Database) UpdateADGroup(adGroup *ent.ADGroup) (*ent.ADGroup, error) {
	foundADGroup, err := d.client.ADGroup.Get(context.Background(), adGroup.ID)
	if err != nil {
		fmt.Printf("failed updating ad group: %v\n", err)
		return nil, fmt.Errorf("failed updating ad group: %v", err)
	}
	updatedADGroup, err := foundADGroup.Update().
		SetName(adGroup.Name).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed updating ad group: %v\n", err)
		return nil, fmt.Errorf("failed updating ad group: %v", err)
	}
	// fmt.Printf("AD Group was updated from: %+v\n to: %+v\n", spew.Sdump(*foundADGroup), spew.Sdump(*updatedADGroup))
	return updatedADGroup, nil
}
func (d *Database) DeleteADGroup(adGroup *ent.ADGroup) error {
	foundADGroup, err := d.client.ADGroup.Get(context.Background(), adGroup.ID)
	if err != nil {
		fmt.Printf("failed deleting ad group: %v\n", err)
		return fmt.Errorf("failed deleting ad group: %v", err)
	}
	err = d.client.ADGroup.DeleteOne(foundADGroup).Exec(context.Background())
	if err != nil {
		fmt.Printf("failed deleting ad group: %v\n", err)
		return fmt.Errorf("failed deleting ad group: %v", err)
	}
	// fmt.Printf("AD Group was deleted: %+v\n", spew.Sdump(*foundADGroup))
	return nil
}
func (d *Database) UpdateCommunity(community *ent.Community) (*ent.Community, error) {
	foundCommunity, err := d.client.Community.Get(context.Background(), community.ID)
	if err != nil {
		fmt.Printf("failed updating community: %v\n", err)
		return nil, fmt.Errorf("failed updating community: %v", err)
	}
	updatedCommunity, err := foundCommunity.Update().
		SetWhamSiteID(community.WhamSiteID).
		SetWhamTitle(community.WhamTitle).
		SetWhanDescription(community.WhanDescription).
		SetWhamCommunityURL(community.WhamCommunityURL).
		SetWhamCreated(community.WhamCreated).
		SetWhamUpdated(community.WhamUpdated).
		SetFeaturedFrom(community.FeaturedFrom).
		SetFeaturedTo(community.FeaturedTo).
		ClearAdgroup().
		AddAdgroup(community.Edges.Adgroup...).
		ClearCommunityCategory().
		AddCommunityCategory(community.Edges.CommunityCategory...).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed updating community: %v\n", err)
		return nil, fmt.Errorf("failed updating community: %v", err)
	}
	// fmt.Printf("Community was updated from: %+v\n to: %+v\n", spew.Sdump(*foundCommunity), spew.Sdump(*updatedCommunity))
	return updatedCommunity, nil
}
func (d *Database) DeleteCommunity(community *ent.Community) error {
	foundCommunity, err := d.client.Community.Get(context.Background(), community.ID)
	if err != nil {
		fmt.Printf("failed deleting community: %v\n", err)
		return fmt.Errorf("failed deleting community: %v", err)
	}
	err = d.client.Community.DeleteOne(foundCommunity).Exec(context.Background())
	if err != nil {
		fmt.Printf("failed deleting community: %v\n", err)
		return fmt.Errorf("failed deleting community: %v", err)
	}
	// fmt.Printf("Community was deleted: %+v\n", spew.Sdump(*foundCommunity))
	return nil
}
func (d *Database) UpdateCommunityCategory(communityCategory *ent.CommunityCategory) (*ent.CommunityCategory, error) {
	foundCommunityCategory, err := d.client.CommunityCategory.Get(context.Background(), communityCategory.ID)
	if err != nil {
		fmt.Printf("failed updating community category: %v\n", err)
		return nil, fmt.Errorf("failed updating community category: %v", err)
	}
	updatedCommunityCategory, err := foundCommunityCategory.Update().
		SetName(communityCategory.Name).
		SetDisplayOrder(communityCategory.DisplayOrder).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed updating community category: %v\n", err)
		return nil, fmt.Errorf("failed updating community category: %v", err)
	}
	// fmt.Printf("Community Category was updated from: %+v\n to: %+v\n", spew.Sdump(*foundCommunityCategory), spew.Sdump(*updatedCommunityCategory))
	return updatedCommunityCategory, nil
}
func (d *Database) DeleteCommunityCategory(communityCategory *ent.CommunityCategory) error {
	foundCommunityCategory, err := d.client.CommunityCategory.Get(context.Background(), communityCategory.ID)
	if err != nil {
		fmt.Printf("failed deleting community category: %v\n", err)
		return fmt.Errorf("failed deleting community category: %v", err)
	}
	err = d.client.CommunityCategory.DeleteOne(foundCommunityCategory).Exec(context.Background())
	if err != nil {
		fmt.Printf("failed deleting community category: %v\n", err)
		return fmt.Errorf("failed deleting community category: %v", err)
	}
	// fmt.Printf("Community Category was deleted: %+v\n", spew.Sdump(*foundCommunityCategory))
	return nil
}
func (d *Database) UpdatePartner(partner *ent.Partner) (*ent.Partner, error) {
	foundPartner, err := d.client.Partner.Get(context.Background(), partner.ID)
	if err != nil {
		fmt.Printf("failed updating partner: %v\n", err)
		return nil, fmt.Errorf("failed updating partner: %v", err)
	}
	updatedPartner, err := foundPartner.Update().
		SetWhamSiteID(partner.WhamSiteID).
		SetWhamTitle(partner.WhamTitle).
		SetWhamDescription(partner.WhamDescription).
		SetKeycloakOrganisation(partner.KeycloakOrganisation).
		SetWhamPartnerURL(partner.WhamPartnerURL).
		SetWhamCreated(partner.WhamCreated).
		SetWhamUpdated(partner.WhamUpdated).
		ClearApplication().
		AddApplication(partner.Edges.Application...).
		Save(context.Background())
	if err != nil {
		fmt.Printf("failed updating partner: %v\n", err)
		return nil, fmt.Errorf("failed updating partner: %v", err)
	}
	// fmt.Printf("Partner was updated from: %+v\n to: %+v\n", spew.Sdump(*foundPartner), spew.Sdump(*updatedPartner))
	return updatedPartner, nil
}
func (d *Database) DeletePartner(partner *ent.Partner) error {
	foundPartner, err := d.client.Partner.Get(context.Background(), partner.ID)
	if err != nil {
		fmt.Printf("failed deleting partner: %v\n", err)
		return fmt.Errorf("failed deleting partner: %v", err)
	}
	err = d.client.Partner.DeleteOne(foundPartner).Exec(context.Background())
	if err != nil {
		fmt.Printf("failed deleting partner: %v\n", err)
		return fmt.Errorf("failed deleting partner: %v", err)
	}
	// fmt.Printf("Partner was deleted: %+v\n", spew.Sdump(*foundPartner))
	return nil
}
