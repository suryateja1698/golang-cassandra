package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suryateja1698/golang-cassandra/db"
	"github.com/suryateja1698/golang-cassandra/models"
)

type PlayerDetailsRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	Country     string `json:"country"`
	Position    string `json:"position"`
	JoinedYear  string `json:"joined_year"`
	JoinedMonth string `json:"joined_month"`
}

type PlayerDetailsResponse struct {
	Status int `json:"status"`
	ID     int `json:"id"`
}

func AddPlayerDetails(c *gin.Context) {
	var req PlayerDetailsRequest
	var res PlayerDetailsResponse

	if err := c.ShouldBindJSON(&req); err != nil {
		res.Status = http.StatusBadRequest
		c.JSON(res.Status, res)
		return
	}

	playerDetails := &models.PlayerDetails{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Age:         req.Age,
		Country:     req.Country,
		Position:    req.Position,
		JoinedYear:  req.JoinedYear,
		JoinedMonth: req.JoinedMonth,
	}

	id, err := db.AddPlayerDetails(c, playerDetails)
	if err != nil {
		res.Status = http.StatusInternalServerError
		c.JSON(res.Status, err.Error())
		return
	}

	if len(id) == 0 {
		res.Status = http.StatusNotFound
		c.JSON(res.Status, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"ID": id,
	})
}
