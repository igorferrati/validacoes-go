package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/stretchr/testify/assert"
)

func SetUpRotasTeste() *gin.Engine {
	rotas := gin.Default()

	return rotas
}

func TestVerificaStatusCode(t *testing.T) {
	r := SetUpRotasTeste()
	r.GET("/:nome", controllers.Saudacao)
	requisicao, _ := http.NewRequest("GET", "/gui", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, requisicao)

	assert.Equal(t, http.StatusOK, resposta.Code, "MENSAGEM DE ERRO")

	mockResposta := `{"API diz":"E ai gui, Tudo beleza?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)

	assert.Equal(t, mockResposta, string(respostaBody))
}
