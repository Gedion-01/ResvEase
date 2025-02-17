# ResvEase

## Find Your Perfect Stay

Discover luxury accommodations worldwide for unforgettable experiences.

## Overview

ResvEase is a modern booking platform that allows users to search, filter, and reserve accommodations seamlessly. Built with Vue.js, Go, Fiber, and MongoDB, it provides a fast and intuitive experience with secure payment processing via Stripe.

## Features

- **Search & Filter**: Find hotels and properties based on location, date range, and number of guests.
- **Advanced Filtering**: Refine search results by price range, amenities, and rating.
- **Room Availability**: View available rooms in a selected property along with their availability count.
- **Secure Booking**: Reserve a room by filling out a form and completing payment via Stripe.
- **Booking Management**: Check your booking history and details on the "My Bookings" page.
- **Date-Based Availability**: Book rooms based on date range, ensuring availability if not already booked by another user.

## Tech Stack

- **Frontend**: Vue.js
- **Backend**: Go, Fiber
- **Database**: MongoDB
- **Payment Processing**: Stripe

## Installation

### Prerequisites

- Node.js
- Go
- MongoDB

### Backend Setup

1. Clone the repository:

   ```sh
   git clone https://github.com/Gedion-01/ResvEase
   cd ResvEase/server
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables in `.env`:
   ```sh
   MONGO_URI=your_mongo_connection_string
   STRIPE_SECRET_KEY=your_stripe_secret_key
   HTTP_LISTEN_ADDRESS=your_http_listen_address e.g :5000
   JWT_SECRET=your_jwt_secret
   MONGO_DB_NAME=your_db_Name
   MONGO_TEST_DB_NAME=test_db_name
   MONGO_DB_URL=your_mongo_db_url
   STRIPE_SECRET_KEY=your_stripe_secret_key
   DOMAIN=your_domain e.g http://localhost:5173
   STRIPE_WEBHOOK_SECRET=your_stripe_webhook_secret
   ```
4. Run the server:
   ```sh
   make run
   ```
5. For testing, add sample data by running:
   ```sh
   make seed
   ```

### Frontend Setup

1. Navigate to the client directory:
   ```sh
   cd ../client
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Set up environment variables in `.env`:
   ```sh
   VITE_API_ENDPOINT=your_back_end_api_endpoint e.g http://localhost:5000/api/v1
   ```
4. Start the development server:
   ```sh
   npm run dev
   ```

## Stripe Webhook Setup

For handling Stripe payments and ensuring proper event processing, you need to set up a webhook:

### Local Testing

1. Install the Stripe CLI.
2. Run the following command to forward events to your local server:
   ```sh
   stripe listen --forward-to localhost:5000/api/v1/stripe/webhook
   ```

### Production Setup

1. Set up your webhook endpoint in the Stripe dashboard to receive real-time event notifications.
2. Ensure your backend is configured to handle Stripe webhook events securely.

## Usage

1. Search for accommodations by entering a location, check-in/out dates, and number of guests.
2. Filter results based on price, amenities, and ratings.
3. Select a property to view available rooms and their availability.
4. Reserve a room by filling out the form and completing payment through Stripe.
5. Book rooms based on a date range if they are available and not already booked by another user.
6. View and manage your bookings under the "My Bookings" section.

## Contributing

Feel free to fork the repository and submit pull requests. Contributions are welcome!

## License

This project is licensed under the MIT License.

## Contact

For inquiries, please reach out to [Your Contact Info].
