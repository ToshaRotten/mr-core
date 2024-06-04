package group

import (
	"encoding/json"
	"main/models/database"
	"main/models/response"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type createGroupRequest struct {
	Group database.Group `json:"group"`
}

type createGroupResponse struct {
	response.Response
}

func CreateGroup(c fiber.Ctx, db *gorm.DB) error {
	var request createGroupRequest
	if err := json.Unmarshal(c.Body(), &request); err != nil {
		return c.JSON(createGroupResponse{response.Error()})
	}
	db.Create(&request.Group)

	return c.JSON(createGroupResponse{response.OK()})
}