package main

import (
	"github.com/bobyard/indexer/db"
	"github.com/bobyard/indexer/models"
	"github.com/bobyard/indexer/suimodels"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bobyard_db := os.Getenv("BOBYARD")
	sui_db := os.Getenv("SUIDB")
	db.Connect(bobyard_db)
	db.ConnectSui(sui_db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("collections", func(c *gin.Context) {
		collections := make([]*models.Collections, 0)
		err := db.Engine.Where("chain_id = ?", db.SUI).Find(&collections)
		if err != nil {
			log.Printf("%v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		c.JSON(http.StatusOK, collections)
	})

	r.GET("collections/:id", func(c *gin.Context) {
		id := c.Param("id")

		collection := new(models.Collections)
		_, err := db.Engine.ID(id).Get(collection)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		objects := make([]*suimodels.Objects, 0)
		err = db.Sui.Where("object_type = ? and object_status != 'DELETED' ", collection.CollectionId).Find(&objects)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		list := make([]*models.Lists, 0)
		err = db.Engine.Where("chain_id = ? and object_type", db.SUI, collection.CollectionId).Find(list)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"lists":   list,
			"objects": objects,
		})
	})

	r.GET("profile/:wallet", func(c *gin.Context) {
		wallet := c.Param("wallet")

		//ALL objects
		objects := make([]*suimodels.Objects, 0)
		err := db.Sui.Where("owner_address = ? and object_status != 'DELETED' ", wallet).Find(&objects)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		// query owner list
		list := make([]*models.Lists, 0)
		err = db.Engine.Where("chain_id = ? and seller_address", db.SUI, wallet).Find(list)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		// query owner offer
		offer := make([]*models.Offers, 0)
		err = db.Engine.Where("chain_id = ? and buyer_address =? ", db.SUI, wallet).Find(offer)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"lists":   list,
			"offers":  offer,
			"objects": objects,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
