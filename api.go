package main

import (
	"github.com/gin-gonic/gin"
	//"io"
	"net/http"
	//"os"
	//"html/template"
	"fmt"
	//"path/filepath"
	"strconv"
	//"time"
	"encoding/json"
)



func (r *gin.Engine) UseApi() error {


	
    eventsV1 := r.Group("/api/v1/events")
    {		
		eventsV1.GET("/:eventId/:orderId/confirm", confirmOrder)
		eventsV1.GET("/:eventId/:orderId", getOrder)
		eventsV1.GET("/:eventId/:orderId/state", getState)
		eventsV1.POST("/:eventId/:orderId/billinginfo", updateBillingInfo)
        eventsV1.POST("/:eventId", createOrder)
    }


	return nil
}

func getOrder (c *gin.Context)  {

}

func getState (c *gin.Context)  {

}


func createOrder (c *gin.Context)  {

}

func updateBillingInfo (c *gin.Context)  {

}


func confirmOrder (c *gin.Context)  {

}