# GoNotes

## Overview

GoNotes is a full-stack note-taking application with a Next.js frontend and a Go backend API. Users can register, log in, and manage personal notes with full CRUD (Create, Read, Update, Delete) functionality. The application uses JWT-based authentication with cookie-based token storage.

## User Preferences

Preferred communication style: Simple, everyday language.

## System Architecture

### Frontend Architecture

- **Framework**: Next.js 16 with App Router
- **Language**: TypeScript
- **Styling**: Tailwind CSS v4 with PostCSS
- **Icons**: Lucide React
- **State Management**: React hooks (useState, useEffect)
- **Routing**: Next.js file-based routing with dynamic routes (e.g., `/notes/[noteid]`)

The frontend runs on port 5000 and communicates with the backend API. Authentication state is managed through HTTP-only cookies, and protected routes redirect unauthenticated users to the login page.

**Key Pages**:
- `/` - Home page displaying all notes with add note functionality
- `/login` - User authentication
- `/register` - New user registration
- `/notes/[noteid]` - Individual note view with edit/delete capabilities

### Backend Architecture

- **Language**: Go (Golang)
- **API Version**: v1 (all endpoints prefixed with `/api/v1`)
- **Authentication**: JWT tokens stored in cookies

**API Endpoints**:
- `POST /auth/register` - Create new user account
- `POST /auth/login` - Authenticate user
- `POST /auth/logout` - End user session
- `GET /notes` - Fetch all user notes
- `POST /notes` - Create new note
- `GET /notes/:id` - Fetch single note
- `PUT /notes/:id` - Update note
- `DELETE /notes/:id` - Delete note

### Data Models

**Note**:
```typescript
{
  id: string;
  title: string;
  content: string;
  author: string;
  created_at: string;
  updated_at: string;
}
```

**User** (inferred from API):
```typescript
{
  username: string;
  email: string;
  password: string;
}
```

### Authentication Flow

1. User registers or logs in via the frontend
2. Backend validates credentials and returns JWT token in `access_token` cookie
3. Frontend includes credentials in all API requests (`credentials: "include"`)
4. Protected routes check authentication by attempting to fetch notes
5. Failed authentication redirects to login page

## External Dependencies

### Frontend Dependencies
- **next** (16.1.6) - React framework
- **react** / **react-dom** (19.2.3) - UI library
- **lucide-react** - Icon library
- **tailwindcss** (v4) - CSS framework

### Backend Dependencies
- Go backend (separate server, referenced via `API_BASE_URL`)
- JWT library for token generation/validation

### External Services
- The backend API is hosted externally and configured via `client/app/config.ts`
- The API base URL points to a Replit deployment