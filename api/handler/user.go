package handler

import (
	"auth/api/token"
	"auth/model"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *handlerImpl) Register(c *gin.Context) {
	req := model.RegisterReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Ma'lumotlarni olishda xatolik",
		})
		return
	}
	mail := isValidEmail(req.Email)
	if !mail {
		h.Logger.Error(fmt.Sprintf("Invalid email: %v", "Kiritilgan email satri, email emas"))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   "Invalid email",
			Message: "Kiritilgan email satri email emas",
		})
	}

	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Passwordni hashlashda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Passwordni hashlashda xatolik",
		})
		return
	}
	req.Password = string(hashpassword)

	resp, err := h.Service.RegisterUser(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Register request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Register request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handlerImpl) Login(c *gin.Context) {
	req := model.LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Malumotlarni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Malumotlarni olishda xatolik",
		})
		return
	}

	mail := isValidEmail(req.Email)
	if !mail {
		h.Logger.Error(fmt.Sprintf("Invalid email: %v", "Kiritilgan email satri, email emas"))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   "Invalid email",
			Message: "Kiritilgan email satri email emas",
		})
	}

	user, err := h.Service.GetUserByEmail(req.Email)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Bunday user mavjud emas, invalid email: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Bunday user mavjud emas, invalid email",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Password xato: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Password xato",
		})
	}

	access, err := token.GenerateAccessToke(user)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Access token generatsiya qilishda xatolik: %v", err))
		c.JSON(http.StatusOK, model.Error{
			Error: err.Error(),
			Message: "Access token generatsiya qilinmadi",
		})
		return 
	}
	refresh, err := token.GenerateRefreshToken(user)
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Rfresh token generatsiya qilishda xatolik: %v", err))
		c.JSON(http.StatusOK, model.Error{
			Error: err.Error(),
			Message: "Rfresh token generatsiya qilinmadi",
		})
		return 
	}

	err = h.Service.SaveRefreshToken(&model.SaveToken{
		UserId: user.Id,
		RefreshToken: refresh,
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Refresh token databazaga saqlanmadi: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Refresh token databazaga saqlanmadi",
		})
	}
	c.JSON(http.StatusOK, model.LogenResp{
		AccessToken: access,
		RefreshToken: refresh,
	})
}

func isValidEmail(email string) bool {
	// Email formatini tekshirish uchun regex andozasi
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
