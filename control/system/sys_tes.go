package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Tes struct {
	Name     string `json:name`
	Password string `json:password`
	Age      int    `json:age`
}

// @Summary 测试接口
// @Router /base/tes [get]
func (b *BaseApi) Tes(c *gin.Context) {
	var tes Tes
	c.ShouldBindJSON(&tes)
	fmt.Println("name:---", tes.Name)
	fmt.Println("password:---", tes.Password)
	fmt.Println("age:---", tes.Age)
}
