package route

import (
	"fooddel-app/database"
	"fooddel-app/model"

	"github.com/gofiber/fiber/v2"
)
//images
func GetImages(c *fiber.Ctx) error {
	je := []model.Images{}
	database.DB.Find(&je)
	return c.JSON(je)
}

func GetImage(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	je := model.Images{}
	database.DB.Find(&je, "image_id", id)
	return c.JSON(je)

}
