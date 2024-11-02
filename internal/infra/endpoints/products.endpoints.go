package endpoints

import (
	"github.com/gofiber/fiber/v3"
	createproductservice "github.com/henrique998/go-ecommerce/internal/app/services/products-services/create-product-service"
	createproductcontroller "github.com/henrique998/go-ecommerce/internal/infra/controllers/products-controllers/create-product-controller"
	uploadimagecontroller "github.com/henrique998/go-ecommerce/internal/infra/controllers/products-controllers/upload-image-controller"
	"github.com/henrique998/go-ecommerce/internal/infra/database/repositories"
	"github.com/henrique998/go-ecommerce/internal/infra/endpoints/middlewares"
	"github.com/henrique998/go-ecommerce/internal/infra/providers"
	"github.com/jackc/pgx/v5"
)

func productsEndpoints(app *fiber.App, conn *pgx.Conn) {
	accountsRepository := repositories.NewPGAccountsRepository(conn)
	accountRolesRepo := repositories.NewPGAccountRolesRepository(conn)
	productsRepo := repositories.NewPGProductsRepo(conn)
	commerceProvider := providers.NewStripeCustommerProvider()

	createProductService := createproductservice.New(productsRepo, commerceProvider)

	createProductController := createproductcontroller.NewCreateProductController(createProductService)

	app.Post("/products", middlewares.AuthMiddleware(accountsRepository), middlewares.OnlyAdminsMiddleware(accountRolesRepo), createProductController.Handle)
	app.Post("/products/image/upload", uploadimagecontroller.UploadImageController)
}
