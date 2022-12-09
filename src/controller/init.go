package controller

import (
    "github.com/gin-gonic/gin"
)

var engine *gin.Engine

func Initialize() error {
    gin.SetMode(gin.DebugMode)
    engine = gin.New()
    SetupRouting(engine)
    return nil
}

func Run() error {
    err := engine.Run(":8080")
    if err != nil {
        return err
    }
    return nil
}
