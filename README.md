# Acuware: Inventory Management System

## Overview
Acuware is an inventory management system built with a microservices architecture, featuring Golang, Ruby on Rails, MongoDB, PostgreSQL, and Docker.

## Project Structure
- **Golang Service**: Provides the backend API for managing inventory.
- **Rails Service**: Manages product data and API.
- **MongoDB**: Stores metadata for fast access.
- **Postgres**: Stores primary inventory data.

## Setup Instructions
1. **Clone the repository**.
2. **Run `docker-compose up --build`** to start all services.
3. Access the services:
   - Golang API: `http://localhost:8080/inventory`
   - Rails API: `http://localhost:3000/products`
