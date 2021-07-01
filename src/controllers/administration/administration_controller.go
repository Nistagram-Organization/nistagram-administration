package administration

import (
	"github.com/Nistagram-Organization/nistagram-administration/src/services/administration"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdministrationController interface {
	TerminateProfile(ctx *gin.Context)
}

type administrationController struct {
	administration.AdministrationService
}

func NewAdministrationController(administrationService administration.AdministrationService) AdministrationController {
	return &administrationController{
		administrationService,
	}
}

func (c *administrationController) TerminateProfile(ctx *gin.Context) {
	email := ctx.Query("email")

	if err := c.AdministrationService.TerminateProfile(email); err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, email)
}
