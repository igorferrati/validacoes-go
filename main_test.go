package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestBuscarAlunoID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetUpRotasTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	path := "/alunos/" + strconv.Itoa(ID)
	requisicao, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock) //armazena o corpo da requisião no struct alunoMock

	assert.Equal(t, "Aluno Teste", alunoMock.Nome, "Nomes devem ser iguais")
	assert.Equal(t, "12345678910", alunoMock.CPF, "CPF devem corresponder ao aluno")
	assert.Equal(t, "123456789", alunoMock.RG, "RG devem corresponder ao aluno")
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestDeleteAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetUpRotasTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	path := "/alunos/" + strconv.Itoa(ID)
	requisicao, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestEditaAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetUpRotasTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "Aluno Teste", CPF: "47123456789", RG: "123456700"}
	valorJson, _ := json.Marshal(aluno)

	path := "/alunos/" + strconv.Itoa(ID)
	requisicao, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(valorJson)) //passando no body o conteúdo json para editar
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	var alunoAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoAtualizado) //preenchendo a struct com a respota para verificação

	assert.Equal(t, "47123456789", alunoAtualizado.CPF)
	assert.Equal(t, "123456700", alunoAtualizado.RG)
	assert.Equal(t, "Aluno Teste", alunoAtualizado.Nome)

}
