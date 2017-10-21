package controllers

import (
  "github.com/astaxie/beego";
  "github.com/jenazads/gojweto";
)

type GoJwetoController struct {
  beego.Controller
}

// Prepare, executes before Http Methods
func (o *GoJwetoController) Prepare() {
  tokenString := o.Ctx.Request.Header.Get(gojweto.GetHeaderKey())
  //tokenString := c.Ctx.Input.Query("tokenString")
  valido, _, _ := gojweto.ValidateHS256Token(tokenString)
  if !valido {
    o.Ctx.Output.SetStatus(401)
    o.Data["json"] = "Permission Deny"
    o.ServeJSON()
  }
  return
}
