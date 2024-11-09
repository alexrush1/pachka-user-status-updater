package restapi

import "github.com/gin-gonic/gin"

func SaveUser(c *gin.Context) {
	var newUser User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

}
