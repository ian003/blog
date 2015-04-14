package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["isHome"] = true
	c.TplNames = "home.html"
	c.Data["isLogin"] = checkAccount(c.Ctx)

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["topics"] = topics
	}
	/*c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"

	c.Data["TrueCond"] = true
	c.Data["FalseCond"] = false

	type user struct {
		Name string
		Age  int
		Sex  string
	}

	u := &user{
		Name: "ian",
		Age:  20,
		Sex:  "Male",
	}

	c.Data["User"] = u

	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Data["Nums"] = nums

	c.Data["Var"] = "hello world"

	c.Data["Html"] = "<div>hello beego</div>"

	c.Data["Pipe"] = "<div>hello beego</div>"*/
}
