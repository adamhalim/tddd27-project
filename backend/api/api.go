package api

	"github.com/gin-contrib/cors"

func Start() {
	handleRequests()
}

func handleRequests() {
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"PUT", "GET", "POST"},
		AllowHeaders: []string{"Content-Type", "Origin"},
	}))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
