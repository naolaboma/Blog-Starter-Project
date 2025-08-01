# Blog Platform - Backend API

A comprehensive blog platform backend API built with Go, featuring user authentication, blog post management, AI integration, and MongoDB Atlas as the database.

## 🚀 Features

- **User Management**: Registration, login, authentication, and role-based access control
- **Blog Posts**: CRUD operations with rich content management
- **Comments System**: Nested comments with like/dislike functionality
- **Search & Filtering**: Advanced search with tags, categories, and popularity metrics
- **AI Integration**: Content suggestions and improvements
- **Security**: JWT authentication, password hashing, rate limiting
- **Performance**: Optimized queries, pagination, and caching
- **Scalability**: MongoDB Atlas with proper indexing and monitoring

## 🏗️ Architecture

```
Mini_PRD/
├── Domain/           # Domain models and DTOs
├── Repository/       # Data access interfaces
├── Usecase/         # Business logic
├── Delivery/        # HTTP handlers and routing
├── Infrastructure/  # Database and external services
├── Database/        # Database design and setup
└── main.go         # Application entry point
```

## 📋 Prerequisites

- Go 1.24.4 or higher
- MongoDB Atlas account
- Git

## 🛠️ Setup Instructions

### 1. Clone the Repository

```bash
git clone <repository-url>
cd Mini_PRD
```

### 2. Install Dependencies

```bash
go mod download
```

### 3. MongoDB Atlas Setup

Follow the detailed setup guide in `Database/mongodb_atlas_setup.md`:

1. Create MongoDB Atlas account
2. Create a new cluster
3. Configure database access
4. Set up network access
5. Get your connection string

### 4. Environment Configuration

1. Copy the example environment file:

```bash
cp env.example .env
```

2. Update `.env` with your MongoDB Atlas credentials:

```env
MONGODB_URI=mongodb+srv://your_username:your_password@your_cluster.mongodb.net/blog_platform?retryWrites=true&w=majority
MONGODB_DATABASE=blog_platform
JWT_SECRET=your_super_secret_jwt_key_here_make_it_long_and_random
JWT_REFRESH_SECRET=your_super_secret_refresh_key_here_make_it_long_and_random
```

### 5. Database Initialization

1. Follow the MongoDB Atlas setup guide to create indexes
2. Run the initialization scripts in MongoDB Atlas shell
3. Verify all collections and indexes are created

### 6. Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

## 📊 Database Design

### Collections

1. **users** - User accounts and profiles
2. **blog_posts** - Blog posts with metadata
3. **comments** - Post comments with threading
4. **user_interactions** - Likes, dislikes, and views
5. **auth_tokens** - JWT tokens for authentication
6. **tags** - Blog post tags
7. **categories** - Blog post categories
8. **ai_suggestions** - AI-generated content suggestions
9. **user_sessions** - User session management

### Key Features

- **Denormalization**: Author names stored in posts for performance
- **Compound Indexes**: Optimized for common query patterns
- **TTL Indexes**: Automatic token expiration
- **Text Search**: Full-text search on posts
- **Aggregation Pipelines**: Popular posts and analytics

## 🔐 Security Features

- **Password Hashing**: bcrypt with salt rounds
- **JWT Authentication**: Access and refresh tokens
- **Role-Based Access**: Admin and User roles
- **Rate Limiting**: API request throttling
- **Input Validation**: Comprehensive request validation
- **CORS Configuration**: Cross-origin resource sharing

## 🚀 API Endpoints

### Authentication

- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login
- `POST /api/auth/logout` - User logout
- `POST /api/auth/refresh` - Refresh access token
- `POST /api/auth/forgot-password` - Password reset request
- `POST /api/auth/reset-password` - Password reset

### Users

- `GET /api/users/profile` - Get user profile
- `PUT /api/users/profile` - Update user profile
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/admin/users/{id}/role` - Update user role (Admin only)

### Blog Posts

- `POST /api/posts` - Create blog post
- `GET /api/posts` - Get all posts (with pagination)
- `GET /api/posts/{id}` - Get post by ID
- `PUT /api/posts/{id}` - Update post
- `DELETE /api/posts/{id}` - Delete post
- `GET /api/posts/search` - Search posts
- `GET /api/posts/popular` - Get popular posts

### Comments

- `POST /api/posts/{id}/comments` - Add comment
- `GET /api/posts/{id}/comments` - Get post comments
- `PUT /api/comments/{id}` - Update comment
- `DELETE /api/comments/{id}` - Delete comment

### Interactions

- `POST /api/posts/{id}/like` - Like post
- `POST /api/posts/{id}/dislike` - Dislike post
- `DELETE /api/posts/{id}/like` - Remove like
- `DELETE /api/posts/{id}/dislike` - Remove dislike

### AI Integration

- `POST /api/ai/suggest` - Get AI suggestions
- `GET /api/ai/suggestions` - Get user suggestions

## 📈 Performance Optimizations

- **Connection Pooling**: Optimized MongoDB connections
- **Indexing Strategy**: Comprehensive database indexes
- **Pagination**: Efficient data retrieval
- **Caching**: Redis integration (planned)
- **Query Optimization**: Aggregation pipelines
- **Goroutines**: Concurrent request handling

## 🔍 Monitoring and Analytics

### MongoDB Atlas Monitoring

- Real-time performance metrics
- Query performance analysis
- Connection monitoring
- Storage usage tracking

### Application Metrics

- Request/response times
- Error rates
- User engagement metrics
- Popular content tracking

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./Usecase -v
```

## 📦 Deployment

### Development

```bash
go run main.go
```

### Production

```bash
# Build the application
go build -o blog-api main.go

# Run with environment variables
./blog-api
```

### Docker (Coming Soon)

```bash
docker build -t blog-api .
docker run -p 8080:8080 blog-api
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License.

## 🆘 Support

For support and questions:

- Check the MongoDB Atlas setup guide
- Review the API documentation
- Open an issue on GitHub

## 🔄 Version History

- **v1.0.0** - Initial release with core blog functionality
- **v1.1.0** - Added AI integration and advanced search
- **v1.2.0** - Performance optimizations and monitoring

---

**Note**: Make sure to follow the security checklist in the MongoDB Atlas setup guide before deploying to production.
