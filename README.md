# SpyCat REST API

This repository contains REST API for managing Spy Cats data using Golang programming language. Tools: Gin (web framework written in Golang) for HTTP request handling, GORM to interact with the database, Docker to containerise the project.
Database: PostgreSQL. 
# SpyCatAgency

## Installation and Running

### Clone the repository:

```bash
git clone https://github.com/rvxt21/spycats-app.git
cd spycats-app
```

### Build the project:
```bash
docker-compose up --build
```

## Description of API routes

### Routes for **SpyCats**

- **POST** `/spycats/` — Create new spy cat;
- **GET** `/spycats/` — Get all spy cats.
- **DELETE** `/spycats/:id` — Delete spy cat by ID.
- **PATCH** `/spycats/:id` — Update spy cat salary.
- **GET** `/spycats/:id` — Get spy cat information by ID.

### Routes for **Missions**

- **POST** `/missions/` — Create new mission.
- **GET** `/missions/` — Get all missions info.
- **DELETE** `/missions/:id` — Delete mission by ID.
- **PATCH** `/missions/:id/updatestatus` — Updating the mission status (completed/uncompleted).
- **PATCH** `/missions/:id/assigncat` — Assigning a cat to a mission.
- **GET** `/missions/:id` — Retrieve mission information by ID.

### Routes for **Targets**

- **POST** `/missions/:id/targets/` — Adding a goal to a mission.
- **DELETE** `/missions/:id/targets/deletetarget` — Removing a target from a mission by ID (passed in the request body).
- **PATCH** `/missions/:id/targets/updatenotes` — Updating notes for a target.
- **PATCH** `/missions/:id/targets/updatestatus` — Updating the status of the goal (completed/uncompleted).


