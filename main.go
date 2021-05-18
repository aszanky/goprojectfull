package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aszanky/goprojectfull/internal/constants"
	"github.com/aszanky/goprojectfull/internal/handler"
	"github.com/aszanky/goprojectfull/internal/helper"
	"github.com/aszanky/goprojectfull/internal/infrastructure/elastic"
	"github.com/aszanky/goprojectfull/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	elasticHost = "http://localhost:9200"
	acqLeadsGen = "acquisition-leads-gen"
)

func main() {
	fmt.Println("GO Tutorial Service Initialization ...... ")
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/user", handler.User)

	//Init Elasticsearch
	elasticInfra, err := elastic.New(elasticHost)
	if err != nil {
		log.Println(err)
	}

	ctx := context.Background()
	// electronics := models.Store{
	// 	WPCode:   "WP9035908135",
	// 	Name:     "Toko Harian",
	// 	Type:     "Seafood",
	// 	Location: "Sabang",
	// }

	// doc := elastic.FlexibleDoc{
	// 	Index: "acquisiton-store",
	// 	Data:  electronics,
	// }

	//Insert Data
	// err = elasticInfra.Insert(ctx, doc)
	// if err != nil {
	// 	log.Println(err)
	// }

	// index := []string{"acquisiton-store"}

	// //Update Data by query
	// query := fmt.Sprintf(constants.UpdateLevelLeadsGen, "Toko Rombongan", "WP9035908135")
	// err = elasticInfra.UpdateByQuery(ctx, index, query)
	// if err != nil {
	// 	log.Println(err)
	// }

	leadsData := models.LeadsGen{
		LeadsID:    9,
		KodeWP:     "WP121313469",
		Level:      5,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Kota:       "Sapudi",
		NoHP:       "08785647469",
		NamaWarung: "Poday",
	}

	data, err := helper.StructJSON(leadsData)
	if err != nil {
		log.Println(err)
	}

	query := fmt.Sprintf(constants.UpsertLeadsGen, data)

	doc := elastic.FlexibleDoc{
		ID:    "290",
		Index: "acquisition-leads-gen",
	}

	//Upsert
	err = elasticInfra.Update(ctx, query, doc)
	if err != nil {
		log.Println(err)
	}

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
