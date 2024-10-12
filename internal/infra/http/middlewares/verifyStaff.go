package middlewares

import "github.com/gin-gonic/gin"

func VerifyStaff(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.JSON(401, gin.H{"success": false, "message": "Missing user information"})
		c.Abort()
		return
	}

	userData, err := user.(User)
	if !err {
		c.JSON(401, gin.H{"success": false, "message": "User information is invalid"})
		c.Abort()
		return
	}

	if userData.RoleID != "1" {
		c.JSON(403, gin.H{"success": false, "message": "User is not authorized to perform this action"})
		c.Abort()
		return
	}

}
