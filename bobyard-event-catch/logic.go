package main

import (
	"encoding/json"
	"github.com/bobyard/indexer/db"
	"github.com/bobyard/indexer/models"
	"github.com/bobyard/indexer/pkg/logger"
	"log"
	"strconv"
	"strings"
	"time"
)

func CatchToDB(msg []byte) bool {
	data := string(msg)
	if strings.Contains(data, "ListEvent") {
		var listEvent db.ListEvent
		if err := json.Unmarshal(msg, &listEvent); err != nil {
			logger.Logger.Panic().Err(err)
		}

		list := new(models.Lists)
		list.ChainId = db.SUI
		list.TokenId = listEvent.MoveEvent.Fields.ListID
		list.SellerAddress = listEvent.MoveEvent.Fields.Owner
		s, err := strconv.Atoi(listEvent.MoveEvent.Fields.Ask)
		if err != nil {
			logger.Logger.Panic().Err(err)
		}

		list.SallerValue = int64(s)
		list.SellerCoinId = 1
		list.SellerEndTime = time.Now()

		_, err = db.Engine.Insert(list)
		if err != nil {
			logger.Logger.Panic().Err(err)
		}
		log.Printf("recver list event and sueccess inserted")

	} else if strings.Contains(data, "MarketCreateEvent") {
		var create db.MarketCreate
		if err := json.Unmarshal(msg, &create); err != nil {
			log.Panicf("%s", err)
		}
		log.Printf("recver Market Create Event")
	} else if strings.Contains(data, "BuyEvent") {
		var buy db.BuyEvent
		if err := json.Unmarshal(msg, &buy); err != nil {
			log.Panicf("%s", err)
		}

		list := new(models.Lists)
		_, err := db.Engine.Where("token_id = ?", buy.MoveEvent.Fields.ListID).Delete(list)
		if err != nil {
			log.Panicf("%s", err)
		}

		// add orders table
		order := new(models.Orders)
		order.TokenId = buy.MoveEvent.Fields.ListID
		order.SellerAddress = buy.MoveEvent.Fields.Owner
		order.BuyerAddress = buy.MoveEvent.Fields.Buyer
		order.Amount = buy.MoveEvent.Fields.Ask
		order.CoinId = db.SUI
		order.ChainId = 1
		order.Time = time.Now()
		_, err = db.Engine.Insert(list)
		if err != nil {
			log.Printf("%v", err)
		}
		log.Printf("recver Buy event and sueccess inserted")

	} else if strings.Contains(data, "OfferEvent") {
		var offer db.OfferToNftEvent
		if err := json.Unmarshal(msg, &offer); err != nil {
			log.Panicf("%s", err)
		}

		// add orders table
		offerDB := new(models.Offers)
		offerDB.TokenId = offer.MoveEvent.Fields.ListID
		offerDB.OfferId = offer.MoveEvent.Fields.OfferID
		offerDB.ChainId = db.SUI
		offerDB.BuyerAddress = offer.MoveEvent.Fields.Owner
		offerDB.Item = ""   //TODO
		offerDB.Amount = "" //TODO

		_, err := db.Engine.Insert(offerDB)

		if err != nil {
			log.Printf("%v", err)
		}

		log.Printf("recevr offer")

	} else if strings.Contains(data, "CancelOfferEvent") {
		var cancel db.CancelOfferEvent
		if err := json.Unmarshal(msg, &cancel); err != nil {
			log.Panicf("%s", err)
		}

		offer := new(models.Offers)
		_, err := db.Engine.Where("offer_id = ?", cancel.MoveEvent.Fields.OfferID).Delete(offer)
		if err != nil {
			log.Panicf("%s", err)
		}

		log.Printf("cancel offer")
	} else if strings.Contains(data, "AcceptOfferEvent") {
		var accpet db.AcceptOfferEvent
		if err := json.Unmarshal(msg, &accpet); err != nil {
			log.Panicf("%s", err)
		}

		list := new(models.Lists)
		_, err := db.Engine.Where("token_id = ?", accpet.MoveEvent.Fields.ListID).Delete(list)
		if err != nil {
			log.Panicf("%s", err)
		}
		// let all offer cancel or maybe owner take by self
		offer := new(models.Offers)
		_, err = db.Engine.Where("offer_id = ?", accpet.MoveEvent.Fields.OfferID).Delete(offer)
		if err != nil {
			log.Panicf("%s", err)
		}

		// add orders table
		order := new(models.Orders)
		order.TokenId = accpet.MoveEvent.Fields.ListID
		order.SellerAddress = accpet.MoveEvent.Fields.Owner
		order.BuyerAddress = accpet.MoveEvent.Fields.Buyer
		order.Amount = "1" //TODO make this real
		order.CoinId = db.SUI
		order.ChainId = 1
		order.Time = time.Now()
		_, err = db.Engine.Insert(list)
		if err != nil {
			log.Printf("%v", err)
		}

	} else {
		logger.Logger.Info().Msg("TODO -------------------")
	}

	return true
}
