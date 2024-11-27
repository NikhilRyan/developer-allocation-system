# Developer Allocation System

## **Overview**
The Developer Allocation System is a comprehensive tool designed to streamline resource allocation and task management in agile development teams. The system optimizes task assignments by recommending the best-matched developers based on skills, availability, experience, and other key factors. This solution ensures efficient project delivery while maintaining balanced workloads for developers.

---

## **Key Features**
- **Developer Recommendation System**: Recommends developers for tasks with a ranking score (0-100) based on matching criteria.
- **Task Management**: Handles task creation, assignment, and tracking with priority, dependencies, and estimation.
- **Developer Management**: Tracks developer availability, skills, experience, workload, and responsibilities.
- **On-Call Rotation Support**: Automates on-call developer scheduling.
- **Database Migration Support**: Handles schema updates seamlessly.
- **Scalable and Portable**: Fully Dockerized to run on any platform.

---

## **Tech Stack**

- **Backend**: Go (Golang)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Containerization**: Docker
- **Frameworks**: Gin (for HTTP APIs), GORM (for ORM)
- **Authentication**: JWT-based authentication

---

## **Project Setup**

### **Prerequisites**
1. [Docker](https://www.docker.com/) installed on your machine.
2. [Postman](https://www.postman.com/) or any HTTP client for testing APIs.

### **Steps to Run**

1. Clone the repository:
   ```bash
   git clone https://github.com/nikhilryan/developer-allocation-system.git
   cd developer-allocation-system
   ```

2. Start the services using Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the application at:
   - **Base URL**: `http://localhost:8080`

---

## **API Endpoints**

### **Developers**

- **GET** `/api/v1/developers` - Fetch all developers.
- **GET** `/api/v1/developers/:id` - Fetch a developer by ID.
- **POST** `/api/v1/developers` - Add a new developer.
- **PUT** `/api/v1/developers/:id` - Update a developer's details.
- **DELETE** `/api/v1/developers/:id` - Delete a developer.
- **PATCH** `/api/v1/developers/:id/availability` - Update a developer's availability.
- **GET** `/api/v1/developers/recommendations/:taskID` - Fetch recommended developers for a task.

### **Tasks**

- **GET** `/api/v1/tasks` - Fetch all tasks.
- **GET** `/api/v1/tasks/:id` - Fetch a task by ID.
- **POST** `/api/v1/tasks` - Add a new task.
- **PUT** `/api/v1/tasks/:id` - Update a task's details.
- **DELETE** `/api/v1/tasks/:id` - Delete a task.

---

## **How It Works**

1. **Add Developers and Tasks**: Use the provided API endpoints to add developers and tasks to the system.
2. **Developer Recommendation**: Call the `/developers/recommendations/:taskID` endpoint to fetch a ranked list of developers for a specific task.
3. **Task Assignment**: Assign tasks to developers based on recommendations or manually using the update APIs.

---

## **Sample Data**
### **Developer**
```json
{
  "name": "Alice Johnson",
  "email": "alice.johnson@example.com",
  "availability": 35.0,
  "start_date": "2021-01-15T00:00:00Z",
  "skill_level": {
    "Go": 5,
    "Python": 4,
    "Docker": 4
  },
  "system_knowledge": 80.0,
  "level": "Senior",
  "current_workload": 20.0,
  "on_call_rotation_week": false,
  "responsibilities": ["Backend Development", "API Design"],
  "is_available": true,
  "time_zone": "UTC-5",
  "language_proficiency": ["English", "Spanish"]
}
```

### **Task**
```json
{
  "title": "Implement User Authentication",
  "description": "Develop a user authentication system with JWT.",
  "estimation": 20.0,
  "required_skills": {
    "Go": 4,
    "JWT": 3
  },
  "required_knowledge": 70.0,
  "priority": "High",
  "delivery_date": "2024-12-15T00:00:00Z",
  "status": "Open",
  "dependencies": [],
  "stakeholders": ["Product Manager"],
  "risk_assessment": "Medium"
}
```

---

## **Contributing**

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Push to your branch.
5. Open a pull request.

---
