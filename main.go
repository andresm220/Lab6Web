package main

//Andrés Mazariegos, 21749 
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//Definimos la estructura de un partido 

type Match struct {
	ID        int    `json:"id"`
	HomeTeam  string `json:"homeTeam"`
	AwayTeam  string `json:"awayTeam"`
	MatchDate string `json:"matchDate"`
}

//Slicy del API, como es simple no necesitamos base de datos 
var matches = []Match{}

var nextID = 1

func main() {
	r := gin.Default()

	// Configurar CORS 
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Grupo de rutas /api
	api := r.Group("/api")
	{
		api.POST("/matches", createMatch)
		api.GET("/matches", getAllMatches)
		api.GET("/matches/:id", getMatchByID)
		api.PUT("/matches/:id", updateMatch)
		api.DELETE("/matches/:id", deleteMatch)
		
		// Endpoints PATCH para operaciones especiales
		api.PATCH("/matches/:id/goals", registerGoal)
		api.PATCH("/matches/:id/yellowcards", registerYellowCard)
		api.PATCH("/matches/:id/redcards", registerRedCard)
		api.PATCH("/matches/:id/extratime", setExtraTime)
	}

	fmt.Println("Servidor corriendo en http://localhost:8080")
	r.Run(":8080")
}

// Handlers para los endpoints CRUD
func createMatch(c *gin.Context) {
	var newMatch Match
	if err := c.ShouldBindJSON(&newMatch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	newMatch.ID = nextID
	nextID++
	matches = append(matches, newMatch)
	c.JSON(http.StatusCreated, newMatch)
}

//Función para hacer get a todos los partidos de la api 
func getAllMatches(c *gin.Context) {
	c.JSON(http.StatusOK, matches)
}

//funcion para hacer get por ID 
func getMatchByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for _, match := range matches {
		if match.ID == id {
			c.JSON(http.StatusOK, match)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
}
//función para actualizar datos de un partido 
func updateMatch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var updatedMatch Match
	if err := c.ShouldBindJSON(&updatedMatch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	for i, match := range matches {
		if match.ID == id {
			matches[i] = Match{
				ID:        id,
				HomeTeam:  updatedMatch.HomeTeam,
				AwayTeam:  updatedMatch.AwayTeam,
				MatchDate: updatedMatch.MatchDate,
			}
			c.JSON(http.StatusOK, matches[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
}

//función para eliminar un partido 
func deleteMatch(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	for i, match := range matches {
		if match.ID == id {
			matches = append(matches[:i], matches[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Partido eliminado"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Partido no encontrado"})
}

// Handlers para operaciones especiales
func registerGoal(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	// Lógica para registrar gol 
	c.JSON(http.StatusOK, gin.H{"message": "Gol registrado", "matchId": id})
}

func registerYellowCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	// Lógica para tarjeta amarilla
	c.JSON(http.StatusOK, gin.H{"message": "Tarjeta amarilla registrada", "matchId": id})
}

func registerRedCard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	// Lógica para tarjeta roja
	c.JSON(http.StatusOK, gin.H{"message": "Tarjeta roja registrada", "matchId": id})
}

func setExtraTime(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	
	// Lógica para tiempo extra
	c.JSON(http.StatusOK, gin.H{"message": "Tiempo extra establecido", "matchId": id})
}
