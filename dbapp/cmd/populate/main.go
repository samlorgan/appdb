package main

import (
	"dbapp/db"
	"dbapp/handlers"
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Running cmd/populate main")

	app := handlers.App{Database: db.Init()}
	fmt.Println("****POPULATING DATABASE****")

	CreateADGroups(app)
	CreateApplicationCategories(app)
	CreateCommunityCategories(app)
	CreatePartnerPageLinkFragments(app)
	CreateCommunities(app)
	CreateApplications(app)
	CreatePartners(app)
	fmt.Println("****READING DATABASE****")
	adGroups := app.Database.GetAllADGroups()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(adGroups), adGroups, spew.Sdump(adGroups))
	appCats := app.Database.GetAllApplicationCategories()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(appCats), appCats, spew.Sdump(appCats))
	apps := app.Database.GetAllApplications()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(apps), apps, spew.Sdump(apps))
	comms := app.Database.GetAllCommunities()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(comms), comms, spew.Sdump(comms))
	commCats := app.Database.GetAllCommunityCategories()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(commCats), commCats, spew.Sdump(commCats))
	pageFrags := app.Database.GetAllPartnerPageLinkFragments()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(pageFrags), pageFrags, spew.Sdump(pageFrags))
	partners := app.Database.GetAllPartners()
	fmt.Printf("Read %d %T \n Breakdown %v\n", len(partners), partners, spew.Sdump(partners))
	fmt.Println("****DONE****")

}

func CreateADGroups(app handlers.App) {
	app.Database.CreateADGroup("ADGroup1")
	app.Database.CreateADGroup("ADGroup2")
	app.Database.CreateADGroup("ADGroup3")
	app.Database.CreateADGroup("ADGroup4")
	app.Database.CreateADGroup("ADGroup5")
}

func CreateApplicationCategories(app handlers.App) {
	app.Database.CreateApplicationCategory("ApplicationCategory1", 1)
	app.Database.CreateApplicationCategory("ApplicationCategory2", 2)
	app.Database.CreateApplicationCategory("ApplicationCategory3", 3)
	app.Database.CreateApplicationCategory("ApplicationCategory4", 4)
	app.Database.CreateApplicationCategory("ApplicationCategory5", 5)

}

func CreateCommunityCategories(app handlers.App) {
	app.Database.CreateCommunityCategory("CommunityCategory1", 1)
	app.Database.CreateCommunityCategory("CommunityCategory2", 2)
	app.Database.CreateCommunityCategory("CommunityCategory3", 3)
	app.Database.CreateCommunityCategory("CommunityCategory4", 4)
	app.Database.CreateCommunityCategory("CommunityCategory5", 5)

}

func CreatePartnerPageLinkFragments(app handlers.App) {
	app.Database.CreatePartnerPageLinkFragment("LinkText1", "URLSuffix1", 1)
	app.Database.CreatePartnerPageLinkFragment("LinkText2", "URLSuffix2", 2)
	app.Database.CreatePartnerPageLinkFragment("LinkText3", "URLSuffix3", 3)
	app.Database.CreatePartnerPageLinkFragment("LinkText4", "URLSuffix4", 4)
	app.Database.CreatePartnerPageLinkFragment("LinkText5", "URLSuffix5", 5)

}

func CreateApplications(app handlers.App) {
	appCats := app.Database.GetAllApplicationCategories()
	ac := appCats[0]
	app.Database.CreateApplication("AppName1", "Description1", "AltText1", "URI1", "IconURI1", true, time.Now(), time.Now(), ac)
	app.Database.CreateApplication("AppName2", "Description2", "AltText2", "URI2", "IconURI2", true, time.Now(), time.Now(), ac)
	app.Database.CreateApplication("AppName3", "Description3", "AltText3", "URI3", "IconURI3", true, time.Now(), time.Now(), ac)
	app.Database.CreateApplication("AppName4", "Description4", "AltText4", "URI4", "IconURI4", true, time.Now(), time.Now(), ac)
	app.Database.CreateApplication("AppName5", "Description5", "AltText5", "URI5", "IconURI5", true, time.Now(), time.Now(), ac)
}

func CreatePartners(app handlers.App) {
	app.Database.CreatePartner(1, "Title1", "Description1", "KeycloakOrganisation1", "PartnerURL1", time.Now(), time.Now(), nil)
	app.Database.CreatePartner(2, "Title2", "Description2", "KeycloakOrganisation2", "PartnerURL2", time.Now(), time.Now(), nil)
	app.Database.CreatePartner(3, "Title3", "Description3", "KeycloakOrganisation3", "PartnerURL3", time.Now(), time.Now(), nil)
	app.Database.CreatePartner(4, "Title4", "Description4", "KeycloakOrganisation4", "PartnerURL4", time.Now(), time.Now(), nil)
	app.Database.CreatePartner(5, "Title5", "Description5", "KeycloakOrganisation5", "PartnerURL5", time.Now(), time.Now(), nil)

}
func CreateCommunities(app handlers.App) {
	app.Database.CreateCommunity(1, "Title1", "Description1", "CommunityURL1", time.Now(), time.Now(), time.Now(), time.Now(), nil, nil)
	app.Database.CreateCommunity(2, "Title2", "Description2", "CommunityURL2", time.Now(), time.Now(), time.Now(), time.Now(), nil, nil)
	app.Database.CreateCommunity(3, "Title3", "Description3", "CommunityURL3", time.Now(), time.Now(), time.Now(), time.Now(), nil, nil)
	app.Database.CreateCommunity(4, "Title4", "Description4", "CommunityURL4", time.Now(), time.Now(), time.Now(), time.Now(), nil, nil)
	app.Database.CreateCommunity(5, "Title5", "Description5", "CommunityURL5", time.Now(), time.Now(), time.Now(), time.Now(), nil, nil)
}

// type IDatabase interface {
// 	CreateApplicationCategory(name string, displayOrder int) (*ent.ApplicationCategory, error)
// 	CreateApplication(name string, description string, altText string, uri string, iconURI string, isFavourite bool, validFrom time.Time, validTo time.Time, applicationCategory *ent.ApplicationCategory) (*ent.Application, error)
// 	CreateADGroup(name string) (*ent.ADGroup, error)
// 	CreateCommunity(whamSiteID int, whamTitle string, whanDescription string, whamCommunityURL string, whamCreated time.Time, whamUpdated time.Time, featuredFrom time.Time, featuredTo time.Time, adGroups []*ent.ADGroup, communityCategories []*ent.CommunityCategory) (*ent.Community, error)
// 	CreateCommunityCategory(name string, displayOrder int) (*ent.CommunityCategory, error)
// 	CreatePartner(whamSiteID int, whamTitle string, whamDescription string, KeycloakOrganisation string, whamPartnerURL string, whamCreated time.Time, whamUpdated time.Time, applications []*ent.Application) (*ent.Partner, error)
// 	CreatePartnerPageLinkFragment(linkText string, whamPartnerURLSuffix string, displayOrder int) (*ent.PartnerPageLinkFragment, error)
// }
