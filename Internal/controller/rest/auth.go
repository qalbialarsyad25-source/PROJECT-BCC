package rest

import (
	"bcc-geazy/pkg/oauth"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (p *V1) LoginGoogle(c *gin.Context) {
	state := oauth.GenerateRandomState()

	c.SetCookie(
		"google_state",
		state,
		3600,
		"/",
		"",
		os.Getenv("APP_ENV") != "development",
		true,
	)

	url := p.usecase.AuthUsecase.GoogleLogin(state)
	c.Redirect(http.StatusFound, url)
}

func (p *V1) CallbackGoogle(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	oauth2state, err := c.Cookie("google_state")
	if err != nil || state != oauth2state {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid oauth state"})
		return
	}

	ctx := c.Request.Context()
	token, err := p.usecase.AuthUsecase.GoogleCallback(ctx, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
