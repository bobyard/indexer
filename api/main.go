package main

import (
	"github.com/bobyard/indexer/models"
	"github.com/bobyard/indexer/suimodels"
	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	APTOS = 0
	SUI   = 1
)

var Engine *xorm.Engine
var Sui *xorm.Engine

func ConnectSui() {
	connStr := "user=sui_indexer password=!Woaini521 dbname=indexer host=127.0.0.1 port=5432 sslmode=disable"
	var err error
	Sui, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Panicf("%v", err)
	}
}

func Connect() {

	connStr := "user=obj password=!Woaini521 dbname=objdb host=127.0.0.1 port=5432 sslmode=disable"
	var err error
	Engine, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		log.Panicf("%v", err)
	}
	//Engine.SetLogger(de)
}

func main() {
	Connect()
	ConnectSui()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("collections", func(c *gin.Context) {
		collections := make([]*models.Collections, 0)
		err := Engine.Where("chain_id = ?", SUI).Find(&collections)
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
		_, err := Engine.ID(id).Get(collection)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		objects := make([]*suimodels.Objects, 0)
		err = Sui.Where("object_type = ? and object_status != 'DELETED' ", collection.CollectionId).Find(&objects)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		list := make([]*models.Lists, 0)
		err = Engine.Where("chain_id = ? and object_type", SUI, collection.CollectionId).Find(list)
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
		err := Sui.Where("owner_address = ? and object_status != 'DELETED' ", wallet).Find(&objects)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		// query owner list
		list := make([]*models.Lists, 0)
		err = Engine.Where("chain_id = ? and seller_address", SUI, wallet).Find(list)
		if err != nil {
			log.Printf("faild %v", err)
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}

		// query owner offer
		offer := make([]*models.Offers, 0)
		err = Engine.Where("chain_id = ? and buyer_address =? ", SUI, wallet).Find(offer)
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
