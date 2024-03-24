package route

import (
	"gorm_demo/handler"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	server := echo.New()

	//User group
	userGroup := server.Group("api/v1/user")
	userGroup.GET("/:id", handler.GetUser)
	userGroup.GET("", handler.GetUsers)
	userGroup.GET("/where", handler.GetUsersWhereConditions)
	userGroup.GET("/struct", handler.GetUsersStructConditions)
	userGroup.GET("/map", handler.GetUsersMapConditions)
	userGroup.GET("/slice", handler.GetUsersSliceConditions)
	userGroup.GET("/or", handler.GetUsersOrConditions)
	userGroup.GET("/select", handler.GetUsersSelectFields)
	userGroup.GET("/dto", handler.GetUsersDTO)
	userGroup.POST("", handler.CreateUser)
	userGroup.PUT("/save/:id", handler.UpdateUserBySave)
	userGroup.PUT("/:id", handler.UpdateUser)
	userGroup.DELETE("", handler.DeleteUsers)

	// //Note group
	// noteGroup := server.Group("api/v1/note")
	// noteGroup.GET("/get-one", handler.GetNote)
	// noteGroup.GET("", handler.GetNotes)
	// noteGroup.POST("", handler.CreateNote)
	// noteGroup.PUT("", handler.UpdateNote)
	// noteGroup.DELETE("", handler.DeleteNote)

	// //Credit card group
	// creditCardGroup := server.Group("api/v1/credit_card")
	// creditCardGroup.GET("/get-one", handler.GetCreditCard)
	// creditCardGroup.GET("", handler.GetCreditCards)
	// creditCardGroup.POST("", handler.CreateCreditCards)
	// creditCardGroup.PUT("", handler.UpdateCreditCard)
	// creditCardGroup.DELETE("", handler.DeleteCreditCard)

	return server
}