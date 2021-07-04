package administration

import (
	"github.com/Nistagram-Organization/nistagram-administration/src/dtos/inappropriate_post_report_decision"
	"github.com/Nistagram-Organization/nistagram-administration/src/services/administration"
	"github.com/Nistagram-Organization/nistagram-shared/src/utils/rest_error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdministrationController interface {
	DecideOnPost(ctx *gin.Context)
}

type administrationController struct {
	administration.AdministrationService
}

func NewAdministrationController(administrationService administration.AdministrationService) AdministrationController {
	return &administrationController{
		administrationService,
	}
}

func (c *administrationController) DecideOnPost(ctx *gin.Context) {
	var reportDecision inappropriate_post_report_decision.InappropriatePostReportDecision
	if err := ctx.ShouldBindJSON(&reportDecision); err != nil {
		restErr := rest_error.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status(), restErr)
		return
	}

	id, err := c.AdministrationService.DecideOnPost(reportDecision)
	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, id)
}
