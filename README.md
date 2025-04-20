Weather Alert Service project

Project Structure:
WeatherApp

├── controllers/

├── initializers/

├── models/

├── services/

├── utils/

├── tests/

├── .env

├── Dockerfile

├── docker-compose.yml

├── go.mod

├── go.sum

├── main.go

└── README.md


### Description:
A Go-based web service that allows users to subscribe for weather notifications based on customizable conditions (e.g., "temperature < 0"). If the specified weather condition is met in the selected city, a daily email notification is sent.

###  Layers

- **Controllers**: Handle HTTP requests (`/weather`, `/users`, `/subscriptions`).
- **Models**: Define data structures and relationships (`User`, `Subscription`, `WeatherCache`, `NotificationHistory`).
- **Initializers**: Setup for the database connection and automatic migrations.
- **Utils**: Business logic helpers like condition parsing and evaluation.
- **Router**: Gin router setup with all routes and middleware.
- **Tests**: Unit and integration tests covering business logic and API.
- **Docker** — Containerization.



## Application Logic

### 1. **User & Subscription Management**
- A user provides their email and subscribes to weather conditions (e.g., `"temperature < 5"` in `"Kyiv"`).
- This data is stored in PostgreSQL using GORM models.

### 2. **Daily Weather Check**
- A scheduler (via `cron/v3`) runs once daily:
    - Checks the latest weather from OpenWeatherMap API.
    - Stores fetched results in the `weather_cache` table.
    - Evaluates all active subscriptions against current conditions.
    - Sends an email if the condition is met.
    - Logs the notification to `notification_history` (to prevent duplicate emails in the same day).

### 3. **APIs**
| Method | Endpoint           | Description                          |
|--------|--------------------|--------------------------------------|
| GET    | `/weather?city=X`  | Get current weather for a city       |
| POST   | `/users`           | Create a new user                    |
| POST   | `/subscriptions`   | Subscribe a user to a weather alert |

##  Choice of Technologies
- **GORM**: The most popular ORM service, provides migrations, and easy model handling.
- **PostgreSQL**: Advanced features, flexibility, scalability, and a robust ecosystem of tools and extensions
- **Gin Web Framework**: The most popular, fast and minimal web framework perfect for API projects.
- **cron/v3**: Lightweight and precise for scheduling daily checks.
- **OpenWeatherMap API**: Free and reliable source for weather data.
- **Email Notifications**: Sent using `gomail.v2`, fast and easy service, supports Gmail SMTP.
- **Docker & docker-compose** — Fast and easy way of containerization.

## Testing

### Includes:
- **Unit tests** for:
    - `CheckCondition()` logic.
    - Model validation (`WeatherCache`, `NotificationHistory`).
- **Integration tests** for:
    - `/subscriptions` API.

PS: This is my first time writing something using GO language. Turns out GO is fun and easy to learn, thanks to SKELAR for the possibility of learning a new technology in a fun way.
