package handler

import (
	"auth/api/token"
	pb "auth/generated/users"
	"auth/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Foydalanuvchilarni olish
// @Description  Ushbu endpoint mavjud foydalanuvchilar ro'yxatini olish uchun ishlatiladi. Ro'yxat sahifalash orqali cheklangan.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        page    query     int    false  "Sahifa raqami, default: 1"
// @Param        limit   query     int    false  "Har bir sahifadagi foydalanuvchilar soni, default: 10"
// @Success      200  {object}  users.GetUserRes  "Foydalanuvchilar ro'yxati muvaffaqiyatli qaytarildi"
// @Failure      400  {object}  model.Error     "Xatolik yuz berdi"
// @Router       /user/getUsers [get]
func (h *handlerImpl) GetUsers(c *gin.Context) {
	var page, limit int
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	resp, err := h.UserService.GetUsers(c, &pb.GetUserReq{
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetUsers request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "GetUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchini o'chirish
// @Description  Ushbu endpoint ma'lum bir foydalanuvchini ID orqali o'chirish uchun ishlatiladi.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        id   path      string  true  "Foydalanuvchi IDsi"
// @Success      200  {object}  users.Massage  "Foydalanuvchi muvaffaqiyatli o'chirildi"
// @Failure      400  {object}  model.Error        "Xatolik yuz berdi"
// @Router       /user/deleteUser{id} [delete]
func (h *handlerImpl) DeleteUser(c *gin.Context) {
	resp, err := h.UserService.DeleteUsers(c, &pb.Id{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteUsers request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "DeleteUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchini ID orqali olish
// @Description  Ushbu endpoint ma'lum bir foydalanuvchini ID orqali olish uchun ishlatiladi.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        id   path      string  true  "Foydalanuvchi IDsi"
// @Success      200  {object}  users.User  "Foydalanuvchi ma'lumotlari muvaffaqiyatli olindi"
// @Failure      400  {object}  model.Error         "Xatolik yuz berdi"
// @Router       /user/getUser{id} [get]
func (h *handlerImpl) GetByIdUsers(c *gin.Context) {
	resp, err := h.UserService.GetByIdUsers(c, &pb.Id{Id: c.Param("id")})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetByIdUsers request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "GetByIdUsers request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchini kuzatish
// @Description  Ushbu endpoint orqali foydalanuvchi boshqa foydalanuvchini kuzatishi mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        id            body      string  true  "Kuzatilayotgan foydalanuvchi IDsi"
// @Success      200           {object}  users.Massage  "Kuzatish muvaffaqiyatli amalga oshirildi"
// @Failure      400           {object}  model.Error       "Xatolik yuz berdi"
// @Router       /user/subscribe [post]
func (h *handlerImpl) Subscribe(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var id string
	err = c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return
	}

	resp, err := h.UserService.Subscribe(c, &pb.FollowingReq{
		UserId:      claim.Id,
		FollowingId: id,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Subscribe request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Subscribe request error",
		})
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchining kuzatuvchilarini olish
// @Description  Ushbu endpoint orqali foydalanuvchining kuzatuvchilarini ro'yxatini olish mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        page          query     int     false "Qaysi sahifani ko'rmoqchisiz? Default: 1"
// @Param        limit         query     int     false "Bir sahifada nechta element bo'lishi kerak? Default: 10"
// @Success      200           {object}  users.GetaFollowersRes  "Foydalanuvchining kuzatuvchilar ro'yxati"
// @Failure      400           {object}  model.Error          "Xatolik yuz berdi"
// @Router       /user/followers [get]
func (h *handlerImpl) GetFollowers(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var page, limit int
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	resp, err := h.UserService.GetFollowers(c, &pb.GetFollowersReq{
		Id:    claim.Id,
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetFollowers request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "GetFollowers request error",
		})
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchidan kuzatuvchini o'chirish
// @Description  Ushbu endpoint orqali foydalanuvchidan kuzatuvchini o'chirishingiz mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        id            body      string  true  "O'chirilishi kerak bo'lgan kuzatuvchi IDsi"
// @Success      200           {object}  users.Massage  "Kuzatuvchini muvaffaqiyatli o'chirdi"
// @Failure      400           {object}  model.Error           "Xatolik yuz berdi"
// @Router       /user/deleteFollower/{id} [delete]
func (h *handlerImpl) DeleteFollower(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var id string
	err = c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
	}

	resp, err := h.UserService.DeleteFollower(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: id,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteFollower request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "DeleteFollower request error",
		})
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Kuzatuvchi ma'lumotlarini olish
// @Description  Ushbu endpoint orqali ma'lum bir kuzatuvchining ma'lumotlarini olish mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        id            body      string  true  "Kuzatuvchi IDsi"
// @Success      200           {object}  users.Follow  "Kuzatuvchi ma'lumotlari"
// @Failure      400           {object}  model.Error  "Xatolik yuz berdi"
// @Router       /user/getFollower/{id} [get]
func (h *handlerImpl) GetByIdFollower(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var id string
	err = c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return
	}

	resp, err := h.UserService.GetByIdFollower(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: id,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetByIdFollower request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "GetByIdFollower request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi kimlarni kuzatayotganini olish
// @Description  Ushbu endpoint orqali ma'lum bir foydalanuvchining kimlarni kuzatayotganini olish mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        page          query     int     false "Sahifa raqami"
// @Param        limit         query     int     false "Sahifadagi elementlar soni"
// @Success      200           {object}  users.GetaFollowingRes  "Kuzatadigan foydalanuvchilar ro'yxati"
// @Failure      400           {object}  model.Error  "Xatolik yuz berdi"
// @Router       /user/following [get]
func (h *handlerImpl) GetFollowing(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var page, limit int
	page, err = strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}
	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	resp, err := h.UserService.GetFollowing(c, &pb.GetFollowingReq{
		Id:    claim.Id,
		Limit: int32(limit),
		Page:  int32(page),
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetFollowing Request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "GetFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi kuzatayotgan foydalanuvchini o'chirish
// @Description  Ushbu endpoint orqali ma'lum bir foydalanuvchining kuzatayotgan boshqa foydalanuvchisini o'chirish mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        id            body      string  true  "Kuzatishni to'xtatmoqchi bo'lgan foydalanuvchi ID'si"
// @Success      200           {object}  users.Massage  "O'chirish muvaffaqiyatli amalga oshirildi"
// @Failure      400           {object}  model.Error  "Xatolik yuz berdi"
// @Router       /users/deleteFollowing/{id} [delete]
func (h *handlerImpl) DeleteFollowing(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var id string
	err = c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return
	}

	resp, err := h.UserService.DeleteFollowing(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: id,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("DeleteFollowing Request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "DeleteFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// @Summary      Foydalanuvchi kimni kuzatayotganini ID orqali olish
// @Description  Ushbu endpoint orqali ma'lum bir foydalanuvchining kimni kuzatayotganini ID orqali olish mumkin.
// @Tags         Users
// @Accept       json
// @Produce      json
// @Security 	 ApiKeyAuth
// @Param        Athorization  header    string  true  "Bearer token bilan birga keladi"
// @Param        id            body      string  true  "Kuzatmoqchi bo'lgan foydalanuvchi ID'si"
// @Success      200           {object}  users.Follow  "Kuzatmoqchi bo'lgan foydalanuvchining ma'lumotlari"
// @Failure      400           {object}  model.Error  "Xatolik yuz berdi"
// @Router       /users/getFollowing/{id} [get]
func (h *handlerImpl) GetByIdFollowing(c *gin.Context) {
	access := c.GetHeader("Athorization")
	claim, err := token.ExtractClaimToken(access)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Claimni olishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Claimdan ma'lumotlarni olishda xatolik",
		})
		return
	}

	var id string
	err = c.ShouldBindJSON(&id)
	if err != nil {
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return
	}

	resp, err := h.UserService.GetByIdFollowing(c, &pb.DeleteFollowerReq{
		UserId:     claim.Id,
		FollowerId: id,
	})
	if err != nil {
		h.Logger.Error(fmt.Sprintf("GetByIdFollowing Request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error:   err.Error(),
			Message: "GetByIdFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
