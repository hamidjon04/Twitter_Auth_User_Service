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

func(h *handlerImpl) DeleteFollower(c *gin.Context){
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
	if err != nil{
		h.Logger.Error("Ma'lumotlarni o'qishda xatolik: %v", err)
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
	}

	resp, err := h.UserService.DeleteFollower(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: id,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("DeleteFollower request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "DeleteFollower request error",
		})
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetByIdFollower(c *gin.Context){
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
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return 
	}

	resp, err := h.UserService.GetByIdFollower(c, &pb.FollowerReq{
		UserId: claim.Id,
		FollowerId: id,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetByIdFollower request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "GetByIdFollower request error",
		})
		return 
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetFollowing(c *gin.Context){
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
		Id: claim.Id,
		Limit: int32(limit),
		Page: int32(page),
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetFollowing Request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "GetFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) DeleteFollowing(c *gin.Context){
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
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return 
	}

	resp, err := h.UserService.DeleteFollowing(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: id,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("DeleteFollowing Request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "DeleteFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func(h *handlerImpl) GetByIdFollowing(c *gin.Context){
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
	if err != nil{
		h.Logger.Error(fmt.Sprintf("Ma'lumotlarni o'qishda xatolik: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "Ma'lumotlarni o'qishda xatolik",
		})
		return 
	}

	resp, err := h.UserService.GetByIdFollowing(c, &pb.DeleteFollowerReq{
		UserId: claim.Id,
		FollowerId: id,
	})
	if err != nil{
		h.Logger.Error(fmt.Sprintf("GetByIdFollowing Request error: %v", err))
		c.JSON(http.StatusBadRequest, model.Error{
			Error: err.Error(),
			Message: "GetByIdFollowing request error",
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
