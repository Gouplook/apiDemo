/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  output
 * @Version: 1.0.0
 * @Date: 2021/1/10 20:51
 */
package functions

import "github.com/gin-gonic/gin"

type output struct {
	Ctx *gin.Context
	encry bool
	data map[string]interface{}
	controller *Controller  // functions.controller

	templates string
}
