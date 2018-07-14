package admin

import (
	. "Loango/models"
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"
	"strconv"
)

//管理页面GET
func LoanManage(c *gin.Context) {
	loans, _ := GetAllProduct()
	print(loans)
	c.HTML(200, "adm_loanManage.html", gin.H{"loans": loans})
}

//删除
func LoanDelete(c *gin.Context) {
	id := c.Param("id")
	DeleteProductByID(id)
	c.Redirect(301, "/admin/loanManage")

}

//新增页面GET
func LoanAdd(c *gin.Context) {
	c.HTML(200, "adm_loanAdd.html", gin.H{})

}

//新增提交POST
func LoanAddPost(c *gin.Context) {

	logoFile, err := c.FormFile("logo")
	if err != nil {
		c.String(http.StatusBadRequest, "a bad requtest")
		return
	}
	if err := c.SaveUploadedFile(logoFile, "statics/img/loan/"+logoFile.Filename); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload logo err:%s"), err.Error())
		return
	}

	name := c.DefaultPostForm("name", "")
	url := c.DefaultPostForm("url", "")
	logo := logoFile.Filename
	amountMax,_ := strconv.ParseInt(c.DefaultPostForm("amount_max","0"), 10, 64)
	period := c.PostForm("period")
	advantage := c.PostForm("advantage")
	applicants,_ := strconv.ParseInt(c.DefaultPostForm("applicants","0") , 10, 64)
	weight,_ := strconv.ParseInt(c.DefaultPostForm("weight","0") , 10, 64)
	company := c.PostForm("company")
	cooperateMethod := c.PostForm("cooperate")

	product := Product{Name:name,Url:url,Logo:logo,AmountMax:amountMax,Period:period,Advantage:advantage,Applicants:applicants,Weight:weight,Company:company,CooperateMethod:cooperateMethod}

	err = product.Create()

	if err == nil {
		c.Redirect(301,"/admin/loanManage")

	} else {
		c.JSON(200, gin.H{
			"error:": product,
		})
	}

}

//编辑页面GET
func LoanEdit(c *gin.Context){
	id := c.Param("id")
	loan,_ := GetProductByID(id)
	c.HTML(200,"adm_loanEdit.html",gin.H{"product":loan})
}


//编辑提交POST
