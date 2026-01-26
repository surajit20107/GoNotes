# Notes API Documentation

Welcome to the official documentation for the Notes API. This API provides a robust and secure way to manage your personal notes, with full authentication and CRUD capabilities.

---

## Base URL
All API requests should be made to:
`http://<your-domain>/api/v1`

---

## Authentication

The API uses **JWT (JSON Web Tokens)** for authentication. Tokens are returned upon successful login/registration and should be included in the `access_token` cookie or as a Bearer token in the `Authorization` header (depending on client implementation).

### 1. Register User
Create a new user account.

- **URL:** `/auth/register`
- **Method:** `POST`
- **Body:**
  ```json
  {
    "username": "johndoe",
    "email": "john@example.com",
    "password": "securepassword123"
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "success": true,
    "message": "User created successfully",
    "user": { ... },
    "token": "eyJhbG..."
  }
  ```

### 2. Login
Authenticate an existing user.

- **URL:** `/auth/login`
- **Method:** `POST`
- **Body:**
  ```json
  {
    "email": "john@example.com",
    "password": "securepassword123"
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "User logged in successfully",
    "user": { ... },
    "token": "eyJhbG..."
  }
  ```

### 3. Logout
Clear the authentication session.

- **URL:** `/auth/logout`
- **Method:** `POST`
- **Auth Required:** Yes
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "User logged out successfully"
  }
  ```

---

## Notes Management

*All notes endpoints require authentication.*

### 1. Get All Notes
Retrieve all notes belonging to the authenticated user.

- **URL:** `/notes/`
- **Method:** `GET`
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "notes": [
      {
        "id": "uuid",
        "title": "Meeting Notes",
        "content": "Discuss project roadmap",
        "created_at": "..."
      }
    ]
  }
  ```

### 2. Create Note
Create a new note.

- **URL:** `/notes/`
- **Method:** `POST`
- **Body:**
  ```json
  {
    "title": "New Note",
    "content": "This is my note content."
  }
  ```
- **Response (201 Created):**
  ```json
  {
    "success": true,
    "message": "Note created successfully",
    "note": { ... }
  }
  ```

### 3. Get Note by ID
Retrieve a specific note by its UUID.

- **URL:** `/notes/:id`
- **Method:** `GET`
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "note": { ... }
  }
  ```

### 4. Update Note
Update an existing note (partial updates supported).

- **URL:** `/notes/:id`
- **Method:** `PUT`
- **Body:**
  ```json
  {
    "title": "Updated Title",
    "content": "New content here."
  }
  ```
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Note updated successfully",
    "note": { ... }
  }
  ```

### 5. Delete Note
Remove a note.

- **URL:** `/notes/:id`
- **Method:** `DELETE`
- **Response (200 OK):**
  ```json
  {
    "success": true,
    "message": "Note deleted successfully"
  }
  ```

---

## Error Handling

The API returns consistent error messages in the following format:

```json
{
  "success": false,
  "error": "Detailed error message here"
}
```

Common HTTP Status Codes:
- `400 Bad Request`: Invalid input data
- `401 Unauthorized`: Authentication failed or missing
- `404 Not Found`: Resource not found
- `409 Conflict`: User already exists (during registration)
- `500 Internal Server Error`: Something went wrong on our end
