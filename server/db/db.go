package db

const MongoDbNameEnvName = "MONGO_DB_NAME"
const MongoTestDbNameEnvName = "MONGO_TEST_DB_NAME"

type Pagination struct {
	Limit int64
	Page  int64
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
