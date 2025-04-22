# Rhythmo — Build better habits with rhythm

📦 **File:** [api.md](https://github.com/Sherida101/Rhythmo/docs/md/api.md)

🔗 Repo: [https://github.com/Sherida101/Rhythmo](https://github.com/Sherida101/Rhythmo)

---

# 📡 Rhythmo REST API

This is the REST API for the Rhythmo Spring Boot backend. It exposes endpoints to manage user habits.

## 🔗 Base URL

```bash
http://localhost:8080/api/habits
```

To set the base API URL using flags or a config file:

```bash
./habit-cli --api=http://localhost:8080/api/habits
```
---

## 📘 Endpoints

### 🧾 Get All Users

**GET** `/api/users`

**Example**
``curl http://localhost:8080/api/users``

**Response**
```json
[
  {
    "id": 1,
    "name": "Jane Doe",
    "email": "jane.doe@example.com",
    "password": "hashed_password",
    "createdAt": "2025-04-17",
    "updatedAt": "2025-04-17",
    "lastLogin": "2025-04-17"
  }
]
```

### 🧾 Get All Habits

**GET** `/api/habits`

**Example**

``curl http://localhost:8080/api/habits``

**Response**
```json
[
  {
    "id": 1,
    "name": "Read Book",
    "lastDone": "2025-04-17",
    "streak": 4
  }
]
```

### ➕ Create Habit

**POST** `/api/habits`

**Example**

``curl http://localhost:8080/api/habits -X POST -H "Content-Type: application/json" -d '{"name": "Exercise", "lastDone": "2025-04-18", "streak": 1}'``

**Request Body**
```json
[
  {
  "name": "Exercise",
  "lastDone": "2025-04-18",
  "streak": 1
 }
]
```

**Response**
```json

{
  "id": 2,
  "name": "Exercise",
  "lastDone": "2025-04-18",
  "streak": 1
}
```

### ✏️ Update Habit

**PUT** `/api/habits/{id}`

**Example**

``curl http://localhost:8080/api/habits/1 -X PUT -H "Content-Type: application/json" -d '{"name": "Exercise", "lastDone": "2025-04-19", "streak": 2}'``

**Request Body**
```json
[
  {
  "name": "Exercise",
  "lastDone": "2025-04-19",
  "streak": 2
 }
]
```

### ❌ Delete Habit

**DELETE** `/api/habits/{id}`

**Example**: ``curl http://localhost:8080/api/habits/1 -X DELETE``

**Response**

```json
204 No Content
```
---

## 🔐 Future Improvements
### 🛠️ Core Functionality (Foundation)

- **Data Management**
  - Store data in PostgreSQL
  - Use JPA (Hibernate) for models

  - Add validation methods to ensure that the habit data is valid (e.g., check if the name is not empty, the frequency is valid, etc.)

  - Add methods to convert the habit to a DTO (Data Transfer Object) for API responses or to a different format for storage in a different database

- **User Management**
  - Spring Security for authentication (JWT or session)
  - User-specific habits

- **Habit Tracking**
  - Handle user interactions, such as marking a habit as done, skipping a habit, etc.
  - Calculate progress, check if a habit is due, etc.
  - Handle habit history, such as tracking the progress over time, tracking the streak, etc.

---

### 📈 Enhanced User Experience (Growth)

- **Analytics & Visualisation**
  - Habit completion history
  - Track time spent per habit

  - Handle habit analytics, such as tracking the success rate, tracking the completion rate, etc.

  - Add a calendar heatmap like GitHub streaks

- **Reminders & Notifications**
  - Handle reminders, notifications, etc.

  - Handle habit reminders, such as sending reminders for a habit, snoozing reminders for a habit, etc.

  - Handle habit notifications, such as sending notifications for a habit, snoozing notifications for a habit, etc.

---

### 🚀 Expansion & Engagement (Scaling)

- **Gamification & Social Features**
  - Handle habit gamification, such as adding points for completing a habit, adding badges for completing a habit, etc.

  - Handle habit challenges, such as creating a challenge for a habit, joining a challenge for a habit, etc.

  - Handle habit rewards, such as adding rewards for completing a habit, redeeming rewards for completing a habit, etc.

  - Handle habit penalties, such as adding penalties for skipping a habit, redeeming penalties for skipping a habit, etc.

  - Handle habit sharing, such as sharing a habit with a friend, sharing a habit with a group, etc.

- **Integration & Synchronisation**
  - Handle habit synchronisation, such as syncing a habit with a calendar, syncing a habit with a task manager, etc.
  - Handle habit integration, such as integrating a habit with a fitness tracker, integrating a habit with a health app, etc.

---

### 📱 Platform Expansion (Future Development)
- **Mobile Development**
  - Build a mobile app later using Kotlin or Flutter and sync with backend

- **Feature Expansion**
  - Expansion from habits to lifestyle, journaling, routines, etc.

- **Performance Enhancements**
  - Add concurrency later to process multiple habits or reminders.

---

### Front end
- Add/Edit/Delete Habits
- Mark as done
- View streaks visually

### Planned endpoints
- Fetch all habits from the server
- Post new habits
- Mark habits as done
- Push/pull sync logic
