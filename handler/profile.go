package handler

import (
	"net/http"
	"strconv"
	"waysbeans/dto"
	"waysbeans/dto/result"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

type ProfileHandler struct {
	ProfileRepository repositories.ProfileRepository
}

func HandlerProfile(ProfileRepository repositories.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{ProfileRepository}
}

func (h *ProfileHandler) FindProfile(c echo.Context) error {
	profile, err := h.ProfileRepository.FindProfile()
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: profile})
}

func (h *ProfileHandler) GetProfile(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var profile models.Profile
	profile, err := h.ProfileRepository.GetProfile(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, result.ErrorResult{Status: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, result.SuccessResult{Status: http.StatusOK, Data: convProfile(profile)})
}

func convProfile(u models.Profile) dto.ProfileResponse {
	return dto.ProfileResponse{
		ID:      u.ID,
		Phone:   u.Phone,
		Gender:  u.Gender,
		Address: u.Address,
		UserID:  u.UserID,
		User:    u.User,
	}
}
