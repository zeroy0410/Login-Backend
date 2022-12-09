package controller

import (
    "Login-Backend/src/model"
    "Login-Backend/src/repository"
    "Login-Backend/src/utility"
    "fmt"

    // "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/copier"
)

func init() {
    RegisterApiRoute(func(router *gin.RouterGroup) {
        AccountRouter := router.Group("/account", DontLoginRequired())
        {
            AccountRouter.POST("/join", AccountJoinHandler)
            AccountRouter.POST("/login", AccountLoginHandler)
        }
        AuthorizedAccountRouter := router.Group("/account", LoginRequired())
        {
            AuthorizedAccountRouter.DELETE("/logout", AccountLogoutHandler)
            AuthorizedAccountRouter.GET("/profile", AccountGetSelfProfileHandler)
            // AuthorizedAccountRouter.PATCH("/profile", AccountChangeProfileHandler)
            // AuthorizedAccountRouter.PATCH("/change-password", AccountChangePasswordHandler)
        }
    })
}

type LoginRequest struct {
    Account  string `json:"account"`
    Password string `json:"password"`
}

type JoinRequest struct {
    Uid      string `json:"uid"`
    Password string `json:"password"`
    Email    string `json:"email"`
    NickName string `json:"nick-name"`
}

func AccountJoinHandler(ctx *gin.Context) {
    var req JoinRequest
    if err := ctx.ShouldBind(&req); err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    if !utility.VerifyEmailFormat(req.Email) ||
        !utility.VerifyPasswordFormat(req.Password) {
        InternalFailedWithMessage(ctx, "invalid format")
        ctx.Abort()
        return
    }
    var user model.User
    err := copier.Copy(&user, req)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user.Admin = false
    hashedPassword, err := utility.HashPassword(user.Password)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user.Password = hashedPassword
    fmt.Println("Before User is: ", user)
    err = repository.CreateUser(&user)
    fmt.Println("Now User is: ", user)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    loginUser(ctx, user)
    Ok(ctx)
}

func AccountLoginHandler(ctx *gin.Context) {
    var req LoginRequest
    err := ctx.ShouldBind(&req)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user, err := repository.ValidateUserPassword(req.Account, req.Password)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    loginUser(ctx, user)
    Ok(ctx)
}

func AccountGetSelfProfileHandler(ctx *gin.Context) {
    userID, err := utility.GetUintFromContext(ctx, "id")
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user, err := repository.GetUserByID(userID)
    if err != nil {
        InternalFailedWithMessage(ctx, err.Error())
        ctx.Abort()
        return
    }
    user.Password = "It's a secret"
    OkWithData(ctx, user)
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
