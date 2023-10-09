package api

import (
	"encoding/json"
	"fmt"
	"github.com/ariandi/kilat-be-go1/cache"
	"github.com/ariandi/kilat-be-go1/services"
	util "github.com/ariandi/kilat-be-go1/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var merchantService services.MerchantInterface
var inqService services.InqInterface
var client = resty.New()

type Server struct {
	Router    *gin.Engine
	config    util.Config
	dbConnect *mongo.Database
}

func NewServer(config util.Config, redis cache.RedisInterface, dbConnect *mongo.Database) (*Server, error) {
	server := &Server{
		config:    config,
		dbConnect: dbConnect,
	}

	setupRouter(server)
	SetCache(redis)
	merchantService = services.GetMerchantService(config, dbConnect, client)
	inqService = services.GetInqService(config, dbConnect, client, redis)
	util.InitLogger()
	logrus.Println("================================================")
	logrus.Printf("Server running at port %s", config.ServerAddress)
	logrus.Println("================================================")
	return server, nil
}

func setupRouter(server *Server) {
	router := gin.Default()
	router.Use(CORSMiddleware())

	authRoutes := router.Group("/api/v2").Use(AuthMiddleware(server.config, []string{"roleName"}))
	authRoutes.POST("/inq", Inq)

	server.Router = router
}

func (server Server) Start(address string) error {
	err := server.Router.Run(address)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func SetCache(redis cache.RedisInterface) {
	_ = redis.TestConnection()
	prodList := []string{"PDAMTBK", "PDAMTKR", "PDAMTKR2"}
	bankList := []string{"PDAMTKR001", "PDAMTKR002", "PDAMTB001"}
	catList := "PDAM001"

	prod, _ := json.Marshal(prodList)
	bank, _ := json.Marshal(bankList)

	key := "PDAM_PROD_LIST"
	keyBank := "PDAM_BANK_LIST"
	keyCat := "CATEGORY"
	ttl := time.Duration(0) * time.Hour
	logrus.Println("set ttl what second : ", int(ttl.Seconds()))

	res, _ := redis.GetRedisVal(keyCat)
	logrus.Println("redis keyCat : ", res)

	_ = redis.SetRedisVal(key, string(prod), int(ttl.Seconds()))
	_ = redis.SetRedisVal(keyBank, string(bank), int(ttl.Seconds()))
	_ = redis.SetRedisVal(keyCat, catList, int(ttl.Seconds()))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
