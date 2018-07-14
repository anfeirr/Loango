package admin

import (
	"github.com/gin-gonic/gin"
	. "Loango/models"
	"net/http"
	"fmt"
	"strconv"
)

func BannerManage(c *gin.Context)  {
	banners,_ := ReadBanners()
	c.HTML(200,"adm_bannerManage.html",gin.H{"banners":banners})

}

func BannerAdd(c *gin.Context){
	c.HTML(200,"adm_bannerAdd.html",gin.H{})
}

func BannerAddPost(c *gin.Context){
	bannerImg,err := c.FormFile("banner")
	if err != nil{
		c.String(http.StatusBadRequest,"a bad requtest")
		return
	}
	if err := c.SaveUploadedFile(bannerImg,"statics/img/banners/"+bannerImg.Filename);err != nil{
		c.String(http.StatusBadRequest,fmt.Sprintf("upload banner err:%s"),err.Error())
		return
	}
	url := c.PostForm("url")
	weightString := c.PostForm("weight")
	weight,_ := strconv.ParseInt(weightString,64,10)

	banner := Banner{ImgName:bannerImg.Filename,Url:url,Weight:weight}

	banner.Create()
	c.Redirect(301,"/admin/bannerManage")
}


func BannerDelete(c *gin.Context) {
	id := c.Param("id")
	err := DeleteBannerById(id)
	c.JSON(200,gin.H{
		"id :":id,
		"err:":err,
	})
	//fmt.Println(id)
	//fmt.Println(err)
	//c.Redirect(301, "/admin/bannerManage")

}

