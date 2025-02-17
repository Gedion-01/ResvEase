package main

import (
	"context"
	"hotel-reservation/api"
	"hotel-reservation/db"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: api.ErrorHandler,
}

func main() {
	mongoEndPoint := os.Getenv("MONGO_DB_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(mongoEndPoint))
	if err != nil {
		log.Fatal(err)
	}
	var (
		hotelStore   = db.NewMongoHotelStore(client)
		bookingStore = db.NewMongoBookingStore(client)
		roomStore    = db.NewMongoRoomStore(client, hotelStore, bookingStore)
		userStore    = db.NewMongoUserStore(client)
		store        = &db.Store{
			Hotel:   hotelStore,
			Room:    roomStore,
			User:    userStore,
			Booking: bookingStore,
		}
		userHandler          = api.NewUserHandler(userStore)
		hotelHandler         = api.NewHotelHandler(store)
		authHandler          = api.NewAuthHandler(userStore)
		roomHandler          = api.NewRoomHandler(store)
		bookingHandler       = api.NewBookingHandler(store)
		stripeWebhookHandler = api.NewStripeWebhookHandler(store)
		app                  = fiber.New(config)
		auth                 = app.Group("/api/v1")
		apiv1                = app.Group("/api/v1")
		admin                = apiv1.Group("/admin")
	)

	// stripe webhook
	app.Post("/api/v1/stripe/webhook", stripeWebhookHandler.HandleStripeWebhook)
	app.Use(cors.New())
	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)
	auth.Post("/signup", userHandler.HandlePostUser)

	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	apiv1.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/api/v1/auth" || c.Path() == "/api/v1/signup" {
			return c.Next()
		}
		return api.JWTAuthentication(userStore)(c)
	})

	// user handlers
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Put("/user/:id", userHandler.HandlePutUser)

	// rooms handlers
	apiv1.Get("/room", roomHandler.HandleGetRooms)

	// bookings handlers
	apiv1.Get("/bookings/user/:id", bookingHandler.HandleGetUserBookings)
	apiv1.Get("/booking/:id", bookingHandler.HandleGetUserBooking)
	apiv1.Post("/room/book", roomHandler.HandleBookRoom)
	apiv1.Get("/booking/:id/cancel", bookingHandler.HandleCancelBooking)

	// admin handlers
	admin.Use(api.AdminAuth)
	// admin.Get("/booking", bookingHandler.HandleGetBookings)

	listenAddr := os.Getenv("HTTP_LISTEN_ADDRESS")
	app.Listen(listenAddr)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

}
