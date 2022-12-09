package main

import (
    "Login-Backend/src/cache"
    "Login-Backend/src/config"
    "Login-Backend/src/controller"
    "Login-Backend/src/repository"
    // "Login-Backend/src/model"
    // "Login-Backend/src/repository"
)

func main() {
    if err := config.Initialize(); err != nil {
        panic("init config error")
    }
    if err := cache.Initialize(); err != nil {
        panic("init cache error")
    }
    if err := repository.Initialize(); err != nil {
        panic("init repository error")
    }
    if err := controller.Initialize(); err != nil {
        panic("init controller error")
    }
    controller.Run()
}
