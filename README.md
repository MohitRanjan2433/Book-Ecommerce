BookEcom - A Book E-commerce Application
Description
BookEcom is a complete e-commerce application for buying and selling books, built using Go Fiber for the backend and integrated with PostgreSQL. It features user authentication, OTP verification, JWT token generation, and book management functionalities. The application also supports cart management, reviews, and order placements.

Features
User Authentication & Authorization: Signup, login, and JWT token management with refresh tokens.
OTP Verification: One-Time Password (OTP) for account verification.
Book Management: Add, edit, and delete books, along with search functionality by title or author.
Cart Management: Add/remove items in the cart and place orders.
Reviews: Users can leave and update reviews for books.
Admin Features: Admins can manage users, books, orders, and reviews.
Technologies Used
Backend: Go (Golang), Fiber
Database: PostgreSQL
Authentication: JWT, bcrypt
Email Service: SendGrid (or custom SMTP)
Other: GORM, UUID, Golang packages for utilities, etc.
Setup Instructions
Prerequisites
Before running the application, ensure you have the following installed:

Go (version 1.20 or later)
PostgreSQL

Clone the Repository
bash
Copy code
git clone https://github.com/yourusername/bookecom.git
cd bookecom
