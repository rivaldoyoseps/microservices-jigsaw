package config

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HttpServerConfig struct {
	DB 		*gorm.DB
	APP 	*fiber.App
}

func StartHttpServer(config *HttpServerConfig) {
	// repository := repository.NewScheduleRepository(config.DB)
	// usecase := usecase.NewScheduleUseCase(repository)
	// api := config.APP.Group("/api/v1")
	// http.NewScheduleHTTPHandler(api, usecase)


	// // Tambahkan ini supaya server jalan
	// if err := config.APP.Listen(":8080"); err != nil {
	// 	panic(err)
	// }


}