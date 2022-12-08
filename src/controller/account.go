package controller

import (
	"Login-Backend/src/model"
	"Login-Backend/src/repository"

	"github.com/gin-gonic/gin"
)

func init() {
	RegisterApiRoute(func(router *gin.RouterGroup) {
		AccountRouter := router.Group("/account", DontLoginRequired())
		{
			// AccountRouter.POST("/join",AccountJoinHandler);
			AccountRouter.POST("/login", AccountLoginHandler)
			AccountRouter.GET("/haha", AccounthahaHandler)
		}
		AuthorizedAccountRouter := router.Group("/account", LoginRequired())
		{
			AuthorizedAccountRouter.DELETE("/logout", AccountLogoutHandler)
			// AuthorizedAccountRouter.GET("/profile", AccountGetSelfProfileHandler)
			// AuthorizedAccountRouter.PATCH("/profile", AccountChangeProfileHandler)
			// AuthorizedAccountRouter.PATCH("/change-password", AccountChangePasswordHandler)
		}
	})
}

type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func AccounthahaHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "haha",
	})
}

func AccountLoginHandler(ctx *gin.Context) {
	var req LoginRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	// fmt.Println("haha: "+req.Account)
	user, err := repository.ValidateUserPassword(req.Account, req.Password)
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
		return
	}
	loginUser(ctx, user)
	Ok(ctx)
}

func AccountLogoutHandler(ctx *gin.Context) {
	_ = RemoveToken(ctx)
	Ok(ctx)
}

func loginUser(ctx *gin.Context, user model.User) {
	err := DistributeToken(ctx, TokenClaims{
		UserID:   user.ID,
		NickName: user.NickName,
		Admin:    user.Admin,
	})
	if err != nil {
		InternalFailedWithMessage(ctx, err.Error())
		ctx.Abort()
	}
}
