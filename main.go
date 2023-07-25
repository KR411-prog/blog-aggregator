package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // underscore means we are not calling it directly
	"blogaggregator/internal/database" //blogaggregator is coming from go.mod
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not in env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("url not in env")
	}

	conn,err := sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal("unable to connect to DB")
	}


	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()
	// cors options is used by server to tell the browser that it should allow http,post,put methods etc. Because the client is connecting to server via a browser.
	router.Use(cors.Handler((cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})))

	//common standard in http routing. Using subrouter to route to specific version and then to paths.
	//router.HandleFunc("/healthz", handlerReadiness) //to scope only for Get requests, we use router.Get
	router.Get("/healthz", handlerReadiness)
	router.Get("/err", handlerErr)
	router.Post("/users", apiCfg.handlerCreateUser)
	router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))
	router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	router.Mount("/v1", router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	log.Printf("Server starting on port %v", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("welcome"))
	// })

	// http.ListenAndServe(":3000", router)

}
