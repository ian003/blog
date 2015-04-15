package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	this.Data["isLogin"] = checkAccount(this.Ctx)
	op := this.Input().Get("op")

	switch op {
	case "add":
		name := this.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		this.Redirect("/category", 301)
		return
	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := models.DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)

	}

	var err error
	this.Data["categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.Data["isCategory"] = true
	this.TplNames = "category.html"
}
