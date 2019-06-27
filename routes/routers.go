package routes

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/proximax-storage/xpx-catapult-faucet"
	"github.com/proximax-storage/xpx-catapult-faucet/services/blockchain"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
	"path"
	"path/filepath"
	"strings"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

var defaultdist *string

func NewRouter(dist *string) *gin.Engine {

	defaultdist = dist
	router := gin.New()

	router.Use(
		gin.Recovery(),
		Faucet.SetCORS(Faucet.Config.Server.AllowCrossDomain),
		gzip.Gzip(gzip.BestSpeed))

	utils.Logger(1, "registering routes")

	router.NoRoute(web)

	authorized := router.Group("/")

	for _, route := range routes {
		var handler gin.HandlerFunc
		handler = route.HandlerFunc
		handler = Faucet.ParseLogger(handler, route.Name)
		switch route.Method {
		case "GET":
			authorized.GET(route.Pattern, handler)
		}
	}
	return router
}

var routes = Routes{

	Route{
		"GetXpx",
		strings.ToUpper("Get"),
		"/api/faucet/GetXpx/:address",
		GetXpx,
	},
}

func GetXpx(ctx *gin.Context) {
	id, err := getAddressParam(ctx)
	if err != nil {
		utils.Logger(2, "%v", "GetXpx fail!")
		respError(ctx, err)
		return
	}

	err = blockchain.TransferXpx(*id, ctx.ClientIP())
	if err != nil {
		respError(ctx, err)
		utils.Logger(2, "%v", "GetXpx fail!")
		return
	} else {
		utils.Logger(0, "%v", "GetXpx complete!")
		respOk(ctx, "XPX sent!")
	}
}

func web(c *gin.Context) {
	dir, file := path.Split(c.Request.RequestURI)
	file = strings.Split(file, "?")[0]
	ext := filepath.Ext(file)
	if file == "" || ext == "" {
		c.File(*defaultdist + "index.html")
	} else {
		c.File(*defaultdist + path.Join(dir, file))
	}
}
