package controllers

import "github.com/gin-gonic/gin"

func RespondWithHttpError(c *gin.Context, message string, status int, err error) {
	var errDescription string = ""

	if err != nil {
		c.Error(err)
		errDescription = err.Error()
	}

	er := HTTPError{
		Error:       message,
		Description: errDescription,
	}
	c.JSON(status, er)
}

type HTTPError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}
