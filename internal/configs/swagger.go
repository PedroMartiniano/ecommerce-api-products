package configs

import "github.com/swaggo/swag"

func SwaggerConfigure(infos *swag.Spec) {
	infos.Title = "E-commerce API Products"
	infos.Description = "API Rest for product and stock management"
	infos.Host = "localhost:8091"
	infos.BasePath = "/"
	infos.Version = "1.0"
	infos.Schemes = []string{"http", "https"}
}
