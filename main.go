package main

import (
	"log"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit-layout/startup"
	"github.com/FarmerChillax/tkit/app"
	"github.com/FarmerChillax/tkit/config"
)

func main() {
	builder, err := app.NewBuilder(&tkit.Application{
		Name: "login-demo",
		Config: &config.Config{
			Database: &config.DatabseConfig{
				Driver: "sqlite3",
				Dsn:    ":memory:",
			},
		},
		RegisterCallback: startup.RegisterCallback(),
	})
	if err != nil {
		log.Fatalln("app.New err: ", err)
	}

	err = builder.ListenGinServer(&tkit.GinApplication{
		RegisterHttpRoute: startup.Register,
	})
	if err != nil {
		log.Fatalln("app.New err: ", err)
	}

}
