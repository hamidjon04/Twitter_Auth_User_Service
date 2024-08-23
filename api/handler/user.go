package handler

import (
	"auth/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func(h *handlerImpl) Register(c *gin.Context){
	req := model.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return 
	}

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Passwordni hashlashda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Passwordni hashlashda xatolik",
		})
		return 
	}
	req.Password = string(hashpassword)

	resp, err := h.Service.RegisterUser(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Register request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Register request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) Login(c *gin.Context){
	req := model.LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Malumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Malumotlarni olishda xatolik",
		})
		return 
	}

	user, err := h.Service.GetUserByEmail(req.Email)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Bunday user mavjud emas, invalid email: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Bunday user mavjud emas, invalid email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Password xato: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Password xato",
		})
	}

	
}