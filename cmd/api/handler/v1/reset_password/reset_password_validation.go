package reset_password

import (
	"github.com/gin-gonic/gin"
	"library/service/reset_password"
	"net/http"
)

var _ Handler = (*ResetPass)(nil)

func New(r reset_password.ServiceResetPass) *ResetPass {
	return &ResetPass{
		res: r,
	}
}

type ResetPass struct {
	res reset_password.ServiceResetPass
}

// SendResetPasswordValidation godoc
// @Summary send reset password validation
// @Schemes
// @Description
// @Tags auth
// @Accept  json
// @Produce  json
// @Param email body SendResetPasswordReq  true "enter email"
// @Success 204
// @Router /auth/reset-password [post]
func (r *ResetPass) SendResetPasswordValidation(c *gin.Context) {
	var icBody SendResetPasswordReq

	if err := c.ShouldBindJSON(&icBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.res.SendResetPasswordEmail(c.Request.Context(), icBody.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found or wrong email"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Validate godoc
// @Summary validate password reset
// @Schemes
// @Description
// @Tags auth
// @Accept  json
// @Produce  json
// @Param req body ValidateReq  true "enter credentials"
// @Success 204
// @Router /auth/validate-reset-password [post]
func (r *ResetPass) Validate(c *gin.Context) {
	var icBody ValidateReq

	if err := c.ShouldBindJSON(&icBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.res.Validate(c.Request.Context(), icBody.Email, icBody.Code, icBody.NewPassword, icBody.ConfirmNewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
