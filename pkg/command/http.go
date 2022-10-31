package command

import (
	"fmt"
	"github.com/nurana88/online-shopping/domain"
	"github.com/nurana88/online-shopping/pkg/monitoring/metrics"
	"github.com/nurana88/online-shopping/pkg/storage"
	"github.com/nurana88/online-shopping/presentation/http/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nurana88/online-shopping/pkg/homehandlers"
	"github.com/urfave/cli"
	"log"
	"net/http"
)

// https://github.com/fr05t1k/http-debug-server/blob/master/http-debug-server.go

type HTTPCommand struct {
	BaseCommand
}

func NewHTTPCommand(basecommand BaseCommand) *HTTPCommand {
	return &HTTPCommand{basecommand}
}

func (h *HTTPCommand) Run(cliCtx *cli.Context) error {
	//ctx := context.Background()

	fmt.Println("In The RUN func")
	//defer cancelCtx()

	prometheus.MustRegister(metrics.RegisterRequestCounter)
	db, err := h.newDBConnection()
	if err != nil {
		log.Fatal("Can't connect to DB", err)
	}

	fmt.Println("db:", db)
	//newRegisterUser := storage.NewUserDBInserter(db)

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {

		r.Use(middleware.Logger)
		fmt.Println("SHould wait")
		userDBActions := storage.NewUserDBActions(db)
		userCreateAction := domain.NewUserCreate(userDBActions)
		r.Method(http.MethodPost, "/registerauth", handlers.NewRegisterUserHandler(userCreateAction))
		r.Get("/welcome", homehandlers.Welcome)

		r.Get("/", homehandlers.Home)
		r.Get("/register", homehandlers.Register)
		r.Get("/login", homehandlers.Login)
		userGetAction := domain.NewGetUser(userDBActions)
		r.Method(http.MethodGet, "/loginauth", handlers.NewGetUserHandler(userGetAction))

		// http.Handle("/assets", http.StripPrefix("/assets", http.FileServer(http.Dir("assets/"))))
		fmt.Println("End of the router")
	})

	router.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting session on :8010...")
	log.Fatal(http.ListenAndServe(":8010", router))

	// log.Fatal(http.ListenAndServe(":9000", http.FileServer(http.Dir("../templates"))))
	return nil
}
