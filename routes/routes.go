package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igorferrati/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")   //carrega html
	r.Static("/assets", "./assets") //carrega css
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.DELETE("/alunos/:id", controllers.DeletaAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNaoEncontrada)
	r.Run()
}
