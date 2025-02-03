# **Planify Dynamic Scheduler - Backend API**
🚀 Scalable, secure, and efficient **appointment scheduling and team management** backend.

---

## **📌 Project Overview**
The **Planify Dynamic Scheduler Backend API** provides a robust backend system to manage **appointments, users, teams, payments, inventory, communication, and AI-powered assistance**.

This backend is built using **microservices architecture** with **Node.js, Express.js, PostgreSQL, and Docker**.

---

## **🚀 Features**
✅ **User Management** - Signup, Login, Role-based Access Control (RBAC), JWT authentication  
✅ **Appointment Scheduling** - Book, reschedule, cancel appointments  
✅ **Team Management** - Manage staff members and permissions  
✅ **Reminders & Notifications** - Automated email/SMS notifications  
✅ **Online Booking System** - Public booking portal for clients  
✅ **Reports & Analytics** - Generate insights from booking data  
✅ **Client Forms & Management** - Store client details & feedback  
✅ **Inventory Management** - Manage products and stock levels  
✅ **Centralized Communication** - Internal messaging & notifications  
✅ **AI Assistant** - Smart recommendations and automation  
✅ **Secure Payments** - Integration with Stripe and PayPal  
✅ **Calendar Sync** - Sync with Google Calendar  
✅ **Review & Ratings** - Clients can leave feedback  
✅ **Product Sales** - Sell items directly from the platform

---

## **🛠️ Technologies Used**
| Technology | Purpose |
|------------|---------|
| **Node.js + Express.js** | Backend Framework |
| **PostgreSQL** | Relational Database |
| **Docker & Docker Compose** | Containerization |
| **JWT (jsonwebtoken)** | Authentication & Security |
| **bcrypt** | Password Hashing |
| **Redis** | Caching for performance optimization |
| **RabbitMQ** | Message queue for asynchronous tasks |
| **Swagger** | API Documentation |
| **Winston** | Logging |
| **Jest** | Unit Testing |

---

## **📂 Folder Structure**
```
backend/
│── src/
│   ├── configs/            # App & database configuration
│   ├── domains/            # All the domains logics here
│   │   ├── controllers/        # Business logic for API endpoints
│   │   ├── middleware/         # JWT auth, error handling, validation
│   │   ├── models/             # Database models & schemas
│   │   ├── services/           # Core business logic & external API calls
│   ├── routes/             # Express.js route handlers
│   ├── utils/              # Helpers (password hashing, validation)
│── tests/                  # Unit & integration tests
│── docker-compose.yml      # Containerization setup
│── package.json            # Dependencies & scripts
│── .env                    # Environment variables
│── README.md               # Documentation
```

---

## **📦 Installation & Setup**

### **1️⃣ Prerequisites**
Ensure you have installed:
- **Node.js** (v18+)
- **Docker** (for PostgreSQL)
- **PostgreSQL** (if not using Docker)

### **2️⃣ Clone Repository**
```bash
git clone https://github.com/yourusername/planify-backend.git
cd planify-backend
```

### **3️⃣ Install Dependencies**
```bash
npm install
```

### **4️⃣ Configure Environment Variables**
Create a `.env` file:
```
DATABASE_URL=postgres://admin:secret@localhost:5432/planify_db
JWT_SECRET=supersecretkey
PORT=5000
```

### **5️⃣ Run the Server**
```bash
npm start
```

### **6️⃣ Run with Docker**
```bash
docker-compose up --build
```

---

## **🚀 API Endpoints**
### **📌 Authentication**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/auth/signup` | Register new user |
| `POST` | `/api/auth/login` | Authenticate user |
| `GET` | `/api/auth/me` | Get logged-in user details |

### **📌 Appointments**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/appointments` | Create new appointment |
| `GET` | `/api/appointments` | Get all appointments |
| `PUT` | `/api/appointments/:id` | Update appointment |
| `DELETE` | `/api/appointments/:id` | Cancel appointment |

### **📌 Reviews**
| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/reviews` | Add a review |
| `GET` | `/api/reviews` | Get all reviews |
| `DELETE` | `/api/reviews/:id` | Delete a review |

📜 **Full API documentation is available via Swagger.**  
Start the server and visit:  
📌 `http://localhost:5000/api-docs`

---

## **🧪 Running Tests**
Run unit tests:
```bash
npm test
```

Run test coverage:
```bash
npm run coverage
```

---

## **📜 Contribution Guidelines**
We welcome contributions! Follow these steps:
1. Fork the repository
2. Create a new feature branch (`git checkout -b feature-name`)
3. Commit changes (`git commit -m "Added new feature"`)
4. Push to your branch (`git push origin feature-name`)
5. Open a Pull Request

---

## **🛠 Future Enhancements**
- **GraphQL Support**
- **WebSockets for real-time updates**
- **Multi-language Support**
- **AI-driven appointment recommendations**
- **More integrations with third-party services**

---

## **📞 Need Help?**
📩 **Email**: support@planify.com  
🌐 **Website**: [planify.com](https://planify.com)  
💬 **Slack Community**: [Join Here](https://slack.planify.com)

---

🔥 **Let's build an intelligent and scalable appointment management system! 🚀**
