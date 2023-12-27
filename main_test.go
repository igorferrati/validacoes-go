package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/igorferrati/api-go-gin/controllers"
	"github.com/igorferrati/api-go-gin/database"
	"github.com/igorferrati/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetUpRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()

	return rotas
}

func CriaAlunoMock() {
	alunoTeste := models.Aluno{Nome: "Aluno Teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&alunoTeste)
	ID = int(alunoTeste.ID)
}

func DeletaAlunoMock() {
	var alunoTeste models.Aluno
	database.DB.Delete(&alunoTeste, ID)
}

func TestSaudacao(t *testing.T) {
	r := SetUpRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	requisicao, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	assert.Equal(t, http.StatusOK, resposta.Code, "MENSAGEM DE ERRO")

	mockResposta := `{"API diz:":"E ai gui, tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)

	assert.Equal(t, mockResposta, string(respostaBody))
}

func TestListaTodosAlunos(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetUpRotasTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	requisicao, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscarPorCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetUpRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	requisicao, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	assert.Equal(t, http.StatusOK, resposta.Code)
}
