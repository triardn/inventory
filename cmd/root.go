package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/triardn/inventory/common"
	"github.com/triardn/inventory/config"
	"github.com/triardn/inventory/driver"
	"github.com/triardn/inventory/handler"
	"github.com/triardn/inventory/middleware"
	"github.com/triardn/inventory/repository"
	"github.com/triardn/inventory/service"
	"github.com/urfave/negroni"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "inventory",
	Short: "A simple API for manage inventory",
	Long:  `Lorem ipsum dolor sit amet`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		InitApp()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.inventory.toml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".inventory" (without extension).
		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("./params")
		viper.AddConfigPath("/opt/inventory/bin")
		viper.AddConfigPath("/opt/inventory/bin/params")
		viper.AddConfigPath("/etc/inventory")
		viper.AddConfigPath("/usr/local/etc/inventory")
		viper.AddConfigPath("/etc/inventory")
		viper.SetConfigType("toml")
		viper.SetConfigName(".inventory")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// Initialize application
func InitApp() {
	config, err := config.NewAppConfig()
	if err != nil {
		log.Fatalf("Config error : %s", err)
	}

	// cache, err := driver.NewCache(config.GetCacheOption())
	// if err != nil {
	// 	log.Fatalf("%s : %v", "Cache error", err)
	// }

	db, err := driver.NewSqliteDatabase(config.GetPathToDBSqlite())
	if err != nil {
		log.Fatalf("%s : %v", "DB error", err)
		panic(err)
	}
	defer db.Close()

	stdOutLogger := log.New()
	stdOutLogger.SetOutput(os.Stdout)
	stdOutLogger.SetLevel(log.DebugLevel)

	stdErrLogger := log.New()
	stdErrLogger.SetOutput(os.Stderr)
	stdOutLogger.SetReportCaller(true)
	stdErrLogger.SetLevel(log.DebugLevel)

	logger := common.NewAPILogger(stdOutLogger, stdErrLogger)

	repo := WiringUpRepository(db, logger, config)

	service, err := WiringUpService(repo, logger, config)
	if err != nil {
		panic(err)
	}

	urlHandler := handler.NewHandler(service)

	r := mux.NewRouter()

	r.Handle("/products/export", handler.HttpHandler{logger, urlHandler.ExportProduct}).Methods(http.MethodGet)
	r.Handle("/orders/export", handler.HttpHandler{logger, urlHandler.ExportOrder}).Methods(http.MethodGet)

	r.Use(middleware.CommonHeaderMiddleware)

	// product
	r.Handle("/products", handler.HttpHandler{logger, urlHandler.GetAllProduct}).Methods(http.MethodGet)
	r.Handle("/products", handler.HttpHandler{logger, urlHandler.CreateProduct}).Methods(http.MethodPost)
	r.Handle("/products/{id}", handler.HttpHandler{logger, urlHandler.GetProductDetail}).Methods(http.MethodGet)
	r.Handle("/products/{id}", handler.HttpHandler{logger, urlHandler.UpdateProduct}).Methods(http.MethodPut)

	// order
	r.Handle("/orders", handler.HttpHandler{logger, urlHandler.GetAllOrder}).Methods(http.MethodGet)
	r.Handle("/orders", handler.HttpHandler{logger, urlHandler.CreateOrder}).Methods(http.MethodPost)
	r.Handle("/orders/{id}", handler.HttpHandler{logger, urlHandler.GetOrder}).Methods(http.MethodGet)

	// order detail
	r.Handle("/order-details", handler.HttpHandler{logger, urlHandler.GetAllOrderDetail}).Methods(http.MethodGet)

	// restock
	r.Handle("/restocks", handler.HttpHandler{logger, urlHandler.GetAllRestock}).Methods(http.MethodGet)
	r.Handle("/restocks", handler.HttpHandler{logger, urlHandler.CreateRestockData}).Methods(http.MethodPost)

	// product sold
	r.Handle("/solds", handler.HttpHandler{logger, urlHandler.GetAllSoldProduct}).Methods(http.MethodGet)
	r.Handle("/solds", handler.HttpHandler{logger, urlHandler.CreateSoldProductData}).Methods(http.MethodPost)

	r.HandleFunc("/health_check", urlHandler.HealthCheck).Methods(http.MethodGet)

	n := negroni.Classic()
	n.UseHandler(r)

	var srv http.Server
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		stdOutLogger.Printf("Server shutdown.. \n")
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			stdErrLogger.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	srv.Addr = fmt.Sprintf("%s:%d", config.App.Host, config.App.Port)
	fmt.Printf("Server address is at: %s\n", srv.Addr)
	srv.Handler = n
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		stdErrLogger.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
	stdOutLogger.Printf("Bye. \n")
}

func WiringUpService(repository *repository.Repository, logger *common.APILogger, appConfig *config.Config) (*service.Service, error) {
	svc := service.NewService()

	personService := service.NewPersonService(repository.Person, logger)
	svc.SetPersonService(personService)

	productService := service.NewProductService(repository.Product, repository.Restock, logger)
	svc.SetProductService(productService)

	orderService := service.NewOrderService(repository.Order, repository.OrderDetail, repository.Restock, logger)
	svc.SetOrderService(orderService)

	orderDetailService := service.NewOrderDetailService(repository.OrderDetail, logger)
	svc.SetOrderDetailService(orderDetailService)

	restockService := service.NewRestockService(repository.Restock, repository.Product, logger)
	svc.SetRestockService(restockService)

	soldService := service.NewSoldService(repository.Sold, repository.Product, logger)
	svc.SetSoldService(soldService)

	return svc, nil
}

func WiringUpRepository(db *gorm.DB, logger *common.APILogger, appConfig *config.Config) *repository.Repository {
	repo := repository.NewRepository()

	option := repository.RepositoryOption{
		DB:        db,
		Logger:    logger,
		AppConfig: appConfig,
	}

	productRepo := repository.NewProductRepository(option)
	repo.SetProductRepository(productRepo)

	orderRepo := repository.NewOrderRepository(option)
	repo.SetOrderRepository(orderRepo)

	orderDetailRepo := repository.NewOrderDetailRepository(option)
	repo.SetOrderDetailRepository(orderDetailRepo)

	restockRepository := repository.NewRestockRepository(option)
	repo.SetRestockRepository(restockRepository)

	soldRepository := repository.NewSoldRepository(option)
	repo.SetSoldRepository(soldRepository)

	return repo
}
