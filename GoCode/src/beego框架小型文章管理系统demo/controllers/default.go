package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"shanghaiyiqi/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["data"] = "china"
	c.TplName = "test.html"
}

func(c*MainController)Post(){
	c.Data["data"] = "上海一期最棒"
	c.TplName = "test.html"
}
func(c*MainController)ShowGet(){
	//获取ORM对象
	o := orm.NewOrm()
	//执行某个操作函数  增删改查
	//插入操作
	/*
	//插入对象
	var user models.User
	//给插入对象赋值
	user.Name = "heima"
	user.PassWord = "chuanzhi"

	//插入操作
	id,err := o.Insert(&user)
	if err != nil{
		beego.Error("插入失败")
	}
	beego.Info(id)
*/

	//查询操作
	/*
	var user models.User
	user.Id = 1

	err := o.Read(&user)
	if err != nil{
		beego.Error("查询失败")
	}
	//返回结果
	beego.Info(user)
	*/
	//更新操作
	/*
	var user models.User
	user.Id = 1
	err := o.Read(&user)
	if err != nil{
		beego.Error("要更新的数据不存在")
	}
	user.Name = "shanghaiyiqi"
	count,err := o.Update(&user)
	if err != nil{
		beego.Error("更新失败")
	}
	beego.Info(count)
	*/

	//删除操作
	var user models.User
	user.Id = 1

	//如果不查询，直接删除，删除对象的主键要有值

	count,err:=o.Delete(&user)
	if err != nil{
		beego.Info("删除失败")
	}
	beego.Info(count)

	c.Data["data"] = "上海"
	c.TplName = "test.html"
}