package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/santaniello/apis/configs"
	_ "github.com/santaniello/apis/docs"
	"github.com/santaniello/apis/internal/entity"
	"github.com/santaniello/apis/internal/infra/database"
	"github.com/santaniello/apis/internal/infra/webserver/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Wesley Willians
// @contact.url    http://www.fullcycle.com.br
// @contact.email  atendimento@fullcycle.com.br

// @license.name   Full Cycle License
// @license.url    http://www.fullcycle.com.br

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configsLoaded, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	/*
		Os middlewares abaixo que estão fora de uma rota em específico serão executados para todas as requests
	*/
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	/*
	 Registrando o midleware de log personalizado que criamos
	*/
	r.Use(CustomLoggerMiddleware)

	// Middleware que faz com que a nossa aplicão não pare de funcionar em caso de panic
	r.Use(middleware.Recoverer)

	// Abaixo usamos um middleware do chi que vai injetar os valores  jwt e JwtExperesIn no contexto da request para que possamos recuperar em qualquer parte da nossa app sem precisar ficar passando ...
	r.Use(middleware.WithValue("jwt", configsLoaded.TokenAuth))
	r.Use(middleware.WithValue("JwtExperesIn", configsLoaded.JwtExperesIn))

	r.Route("/products", func(r chi.Router) {
		/*
			Os middlewares abaixo irão funcionar somente para as rotas de products
		*/

		/* O middlerware abaixo vai procurar por um JWT token em uma http request, na ordem:
			   1. 'jwt' URI query parameter
			   2. 'Authorization: BEARER T' request header
			   3. Cookie 'jwt' value

		       Após encontrar o token, ele seta o *jwt.Token object no request context.
		*/
		r.Use(jwtauth.Verifier(configsLoaded.TokenAuth))
		/*
		   O middlerware abaixo vai validar o token encontrado pelo middleware verifier
		*/
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)
	})
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Post("/generate_token", userHandler.GetJWT)
	})

	// Add endpoint do swagger
	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))
	http.ListenAndServe(":8000", r)
}

/*
	     Um middleware em Go é basicamente uma função que recebe e retorna um http.Handler.
		 Abaixo, a função CustomLoggerMiddleware recebe um parâmetro de entrada que é um handler HTTP. Este é o próximo handler na cadeia de middleware ou no roteador que será chamado.
		 No retorno, retornamos um http.Handler que pode ser outro middleware ou o handler final.
*/
func CustomLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware personalizado, Requested URL:", r.URL.Path)
		/*
			A funço abaixo chama o próximo manipulador na cadeia. Isso é essencial em um middleware, pois permite que a solicitação prossiga para o próximo manipulador (que pode ser outro middleware ou o manipulador final).
		*/
		next.ServeHTTP(w, r)
	})
}
