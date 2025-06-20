

# Habit Tracker

A simple RESTful API service built with Go, Gin, and SQLite for tracking daily habits.

## Overview

This project implements a backend API to help users create, view, and manage habits. It allows you to:

* Create new habits with a name, description, goal duration, and creation timestamp.
* Retrieve all habits stored in the database.
* Fetch a specific habit by its ID.

The backend uses:

* **Go** for application logic
* **Gin** as the HTTP web framework
* **SQLite** as the lightweight embedded database

---

## Features

* **Create Habit:** `POST /habits` — add a new habit by sending JSON with habit details.
* **Get All Habits:** `GET /allhabits` — retrieve a list of all habits.
* **Get Habit By ID:** `GET /habitbyid/:id` — retrieve a specific habit by its unique ID.

---

## Getting Started

### Prerequisites

* Go 1.18+ installed
* Git

### Installation

1. Clone the repository

```bash
git clone https://github.com/santoridev/habit-tracker.git
cd habit-tracker
```

2. Install dependencies

```bash
go mod download
```

3. Run the server

```bash
go run .
```

The server will start on `http://localhost:8080`.

---

## API Usage Examples

### Create Habit

```bash
curl -X POST http://localhost:8080/habits \
  -H "Content-Type: application/json" \
  -d '{"name": "Morning Run", "description": "Jog 30 mins daily", "goal_days": 30}'
```

### Get All Habits

```bash
curl http://localhost:8080/allhabits
```

### Get Habit By ID

```bash
curl http://localhost:8080/habitbyid/1
```

---

## Project Structure

* `/handlers` — HTTP request handlers for creating and fetching habits.
* `/database` — database initialization and connection handling.
* `/models` — data models defining habit structure.
* `main.go` — application entry point, sets up routes and starts the server.

---

## Notes

* SQLite database file (`habits.db`) is created automatically on startup.
* The app uses prepared statements to safely interact with the database.
* Error handling ensures appropriate HTTP status codes are returned on failure.




