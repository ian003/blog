package controllers

import "github.com/astaxie/beego"
import "blog/models"

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["isLogin"] = checkAccount(this.Ctx)
	this.Data["isTopic"] = true
	this.TplNames = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["topics"] = topics
	}
}

func (this *TopicController) Post() {

	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	// 解析表单
	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")

	var err error

	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)

}

func (this *TopicController) Add() {
	this.TplNames = "topic_add.html"
}

func (this *TopicController) View() {
	this.TplNames = "topic_view.html"
	// '/view/1/3' 中this.Ctx.Input.Params["0"]的值为1,this.Ctx.Input.Params["1"]的值为3
	topic, err := models.GetTopic(this.Ctx.Input.Params["0"]) //自动路由，获取参数
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Params["0"]
}

func (this *TopicController) Modify() {
	this.TplNames = "topic_modify.html"

	tid := this.Input().Get("tid")

	topic, err := models.GetTopic(tid)

	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["topic"] = topic
	this.Data["tid"] = tid
}

func (this *TopicController) Delete() {

	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteTopic(this.Ctx.Input.Params["0"])

	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}
