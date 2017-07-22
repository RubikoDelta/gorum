package controllers

import (
	"fmt"
	"github.com/opennut/gorum/app/models"
	"github.com/revel/revel"
	"html/template"
	"modules/gorm/app"

	"github.com/russross/blackfriday"

	"github.com/xeonx/timeago"
	"time"
)

func InitializeDB() {
	gorm.InitDB()
	gorm.DB.AutoMigrate(&models.User{})
	gorm.DB.AutoMigrate(&models.Tag{})
	gorm.DB.AutoMigrate(&models.Discussion{})
	gorm.DB.AutoMigrate(&models.DiscussionComment{})
	// var firstUser = models.User{Name: "Demo", Email: "demo@demo.com"}
	// firstUser.SetNewPassword("demo")
	// firstUser.Active = true
	// gorm.DB.Create(&firstUser)
}

func init() {
	revel.OnAppStart(InitializeDB)
	revel.InterceptMethod((*gorm.GormController).Begin, revel.BEFORE)
	revel.InterceptMethod(PublicApp.AddUser, revel.BEFORE)
	revel.InterceptMethod(App.checkUser, revel.BEFORE)
	revel.InterceptMethod(AdminApp.checkUser, revel.BEFORE)
	revel.InterceptMethod((*gorm.GormController).Commit, revel.AFTER)
	revel.InterceptMethod((*gorm.GormController).Rollback, revel.FINALLY)

	revel.TemplateFuncs["markdown"] = func(str interface{}) template.HTML {
		s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", str)))
		return template.HTML(s)
	}

	revel.TemplateFuncs["timeagos"] = func(time time.Time) string {
		s := timeago.English.Format(time)
		return s
	}
}
