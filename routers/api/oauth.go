package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/sun-wenming/gin-auth/pkg/app"
	"github.com/sun-wenming/gin-auth/pkg/e"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/oauth"
	"github.com/sun-wenming/gin-auth/pkg/util"
	"github.com/sun-wenming/gin-auth/service/userser"
	"golang.org/x/oauth2"
	"net/http"
	"strconv"
)

var oauthStateString = "random-user"

// LoginGithub github登录/注册
func LoginGithub(c *gin.Context) {
	oauthStateString = util.GetRandomSalt()
	url := oauth.GithubOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	c.Redirect(http.StatusMovedPermanently, url)
}

// CallBackGithub 登录成功
func CallBackGithub(c *gin.Context) {
	state, _ := c.GetQuery("state")
	code, _ := c.GetQuery("code")
	appG := app.GetGin(c)

	// TODO 如果服务器重启了, oauthStateString就失效了
	if state != oauthStateString {
		err := fmt.Sprintf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		logging.GetLogger().Error(err)
		appG.ResponseFailError(util.ErrNewCode(e.ErrorOauthState))
		// c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	token, err := oauth.GithubOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		msg := fmt.Sprintf("oauthConf.Exchange() failed with '%s'\n", err)
		logging.GetLogger().Error(msg)
		appG.ResponseFailError(util.ErrNewCode(e.ErrorOauthCode))
		// c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	oauthClient := oauth.GithubOauthConfig.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		msg := fmt.Sprintf("client.Users.Get() faled with '%s'\n", err)
		logging.GetLogger().Error(msg)
		appG.ResponseFailError(util.ErrNewCode(e.ErrorOauthInfo))
		// c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 3: Github
	userID := strconv.FormatInt(*(user.ID), 10)

	userService := userser.UserOauth{OauthID: userID, OauthType: 3, OauthAccessToken: token.AccessToken, OauthExpires: "3600"}
	exist, merr := userService.ExistUserOauth()
	if merr != nil {
		appG.ResponseFailError(merr)
		return
	}

	if exist {
		goto Success
	}
	// 不存在创建一个
	if err := userService.LoginGithub(); err != nil {
		appG.ResponseFailError(err)
		return
	}
Success:
	appG.ResponseSuc(token.AccessToken)

}
