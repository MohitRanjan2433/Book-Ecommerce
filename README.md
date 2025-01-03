# BookEcom - Book Ecommerce Application

## Overview

**BookEcom** is a book ecommerce platform designed with a backend built using **Go Fiber** . The project supports user authentication, book management, cart, orders, reviews, OTP-based registration, and more.

---

## Features

- **User Authentication**: Supports login and signup with password hashing and JWT token-based authentication.
- **Book Management**: Create, read, update, and delete book details such as ISBN, title, author, and price.
- **Cart and Orders**: Manage users' shopping cart and process orders.
- **OTP Verification**: One-time password (OTP) generation and verification for user registration.
- **Reviews**: Users can review books by providing ratings and comments.

---

## Tech Stack

- **Backend**: Go with Fiber framework
- **Database**: PostgreSQL (via GORM)
- **Authentication**: JWT (JSON Web Tokens) and bcrypt for password hashing
- **Email Service**: SMTP for sending OTP emails
- **Middleware**: CORS and Logger

---

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/bookecom.git
    cd bookecom
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

3. Set up environment variables by creating a `.env` file:

    ```plaintext
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER_NAME=your_db_user
    DB_USER_PASSWORD=your_db_password
    DB_NAME=bookStore
    PORT=4000
    EMAIL=your_email
    EMAIL_PASSWORD=your_password
    PRODUCTION=true
    ACCESS_TOKEN_SECRET=books
    REFRESH_TOKEN_SECRET=newbooks
    ACCESS_TOKEN_EXPIRY=15m
    REFRESH_TOKEN_EXPIRY=168h
    TIMEZONE=Asia/Shanghai
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

---

## Routes

### Authentication Routes

- **POST** `/user/signup` - Register a new user
- **POST** `/user/login` - Login with email and password
- **POST** `/user/verify-otp/:userID` - Verify OTP during registration
- **GET** `/user/resend` - Resend OTP for verification

### Book Routes

- **POST** `/book` - Create a new book (Token validation required)
- **GET** `/book/allBooks` - Get all books (Token validation required)
- **GET** `/book/:bookId` - Get book details by ID
- **DELETE** `/book/:bookId` - Delete a book by ID (Token validation required)

### Cart Routes

- **POST** `/cart/items` - Add an item to the cart (Token validation required)
- **DELETE** `/cart/items` - Delete an item from the cart (Token validation required)
- **GET** `/cart` - Get the user's cart (Token validation required)
- **GET** `/cart/all` - Get all carts

### Order Routes

- **POST** `/order` - Create a new order (Token validation required)
- **GET** `/order` - Get all orders for a user (Token validation required)
- **GET** `/order/:orderId` - Get details of a specific order (Token validation required)

### Review Routes

- **POST** `/review` - Create a new review for a book (Token validation required)
- **GET** `/review/book/:bookId` - Get reviews for a specific book (Token validation required)
- **DELETE** `/review/:reviewId` - Delete a review (Token validation required)

### User Routes

- **GET** `/user/me` - Get user details (Token validation required)
- **POST** `/user/me/activate` - Activate the user
- **DELETE** `/user/me/delete` - Delete the user (Token validation required)

---

## Database Schema

### User

- `id`: UUID
- `email`: string (unique)
- `password`: string (hashed)
- `phone_number`: string
- `role`: string (admin, user)
- `otp`: string
- `verified`: boolean
- `created_at`: timestamp
- `updated_at`: timestamp

### Book

- `id`: UUID
- `isbn`: string (unique)
- `title`: string
- `author`: string
- `description`: string
- `genre`: string
- `price`: float64
- `quantity`: int
- `cover_images`: string (JSON array of image URLs)
- `user_id`: UUID (user who added the book)

### Order

- `id`: UUID
- `user_id`: UUID (user who placed the order)
- `cart_id`: UUID
- `total_price`: float64
- `status`: string (pending, completed, cancelled)
- `created_at`: timestamp
- `updated_at`: timestamp

### Cart

- `id`: UUID
- `user_id`: UUID
- `active`: boolean
- `created_at`: timestamp
- `updated_at`: timestamp

### Review

- `id`: UUID
- `user_id`: UUID
- `book_id`: UUID
- `username`: string
- `rating`: float64
- `comment`: string
- `created_at`: timestamp
- `updated_at`: timestamp

---

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

## Contact

For questions, contact me at:

- Email: your_email@example.com
- GitHub: https://github.com/yourusername
