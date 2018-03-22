package gojweto

func Example_BeeGo()(){
  type GoJwetoController struct {
    beego.Controller
  }

  // Prepare, executes before Http Methods
  func (o *GoJwetoController) Prepare() {
    tokenString := o.Ctx.Request.Header.Get(models.GojwtObject.GetHeaderKey())
    //tokenString := c.Ctx.Input.Query("tokenString")
    valido, _, _ := GojwtObject.ValidateToken(tokenString)
    if !valido {
      //o.Ctx.Output.SetStatus(401)
      //o.Ctx.ResponseWriter.WriteHeader(401)
      o.Abort("401")
      //o.Data["json"] = "Permission Deny"
      //o.ServeJSON()
    }
    return
  }
}
