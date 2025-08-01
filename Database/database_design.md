# MongoDB Database Design for Blog Platform

## Database: `blog_platform`

### 1. Users Collection

**Collection Name:** `users`

```json
{
  "_id": "ObjectId",
  "username": "string (unique, required)",
  "email": "string (unique, required)",
  "password": "string (hashed with bcrypt, required)",
  "firstName": "string",
  "lastName": "string",
  "bio": "string",
  "profilePicture": "string (URL)",
  "role": "string (enum: 'user', 'admin', default: 'user')",
  "isActive": "boolean (default: true)",
  "isEmailVerified": "boolean (default: false)",
  "emailVerificationToken": "string",
  "emailVerificationExpires": "Date",
  "passwordResetToken": "string",
  "passwordResetExpires": "Date",
  "lastLogin": "Date",
  "createdAt": "Date (default: now)",
  "updatedAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "email": 1 }` (unique)
- `{ "username": 1 }` (unique)
- `{ "role": 1 }`
- `{ "isActive": 1 }`

### 2. Blog Posts Collection

**Collection Name:** `blog_posts`

```json
{
  "_id": "ObjectId",
  "title": "string (required)",
  "content": "string (required)",
  "excerpt": "string",
  "authorId": "ObjectId (ref: users, required)",
  "authorName": "string (denormalized for performance)",
  "tags": ["string"],
  "category": "string",
  "status": "string (enum: 'draft', 'published', 'archived', default: 'draft')",
  "viewCount": "number (default: 0)",
  "likeCount": "number (default: 0)",
  "dislikeCount": "number (default: 0)",
  "commentCount": "number (default: 0)",
  "featuredImage": "string (URL)",
  "slug": "string (unique, URL-friendly version of title)",
  "readingTime": "number (estimated minutes)",
  "publishedAt": "Date",
  "createdAt": "Date (default: now)",
  "updatedAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "authorId": 1 }`
- `{ "status": 1 }`
- `{ "publishedAt": -1 }`
- `{ "viewCount": -1 }`
- `{ "likeCount": -1 }`
- `{ "tags": 1 }`
- `{ "category": 1 }`
- `{ "slug": 1 }` (unique)
- `{ "title": "text", "content": "text" }` (text search)

### 3. Comments Collection

**Collection Name:** `comments`

```json
{
  "_id": "ObjectId",
  "postId": "ObjectId (ref: blog_posts, required)",
  "authorId": "ObjectId (ref: users, required)",
  "authorName": "string (denormalized)",
  "content": "string (required)",
  "parentCommentId": "ObjectId (ref: comments, for nested comments)",
  "likeCount": "number (default: 0)",
  "dislikeCount": "number (default: 0)",
  "isEdited": "boolean (default: false)",
  "isDeleted": "boolean (default: false)",
  "createdAt": "Date (default: now)",
  "updatedAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "postId": 1 }`
- `{ "authorId": 1 }`
- `{ "parentCommentId": 1 }`
- `{ "createdAt": -1 }`

### 4. User Interactions Collection

**Collection Name:** `user_interactions`

```json
{
  "_id": "ObjectId",
  "userId": "ObjectId (ref: users, required)",
  "postId": "ObjectId (ref: blog_posts, required)",
  "interactionType": "string (enum: 'like', 'dislike', 'view', required)",
  "createdAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "userId": 1, "postId": 1, "interactionType": 1 }` (unique compound)
- `{ "postId": 1, "interactionType": 1 }`
- `{ "userId": 1 }`

### 5. Authentication Tokens Collection

**Collection Name:** `auth_tokens`

```json
{
  "_id": "ObjectId",
  "userId": "ObjectId (ref: users, required)",
  "tokenType": "string (enum: 'access', 'refresh', required)",
  "token": "string (required)",
  "isRevoked": "boolean (default: false)",
  "expiresAt": "Date (required)",
  "createdAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "userId": 1, "tokenType": 1 }`
- `{ "token": 1 }` (unique)
- `{ "expiresAt": 1 }` (TTL index)

### 6. Tags Collection

**Collection Name:** `tags`

```json
{
  "_id": "ObjectId",
  "name": "string (unique, required)",
  "description": "string",
  "postCount": "number (default: 0)",
  "createdAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "name": 1 }` (unique)
- `{ "postCount": -1 }`

### 7. Categories Collection

**Collection Name:** `categories`

```json
{
  "_id": "ObjectId",
  "name": "string (unique, required)",
  "description": "string",
  "slug": "string (unique, URL-friendly)",
  "postCount": "number (default: 0)",
  "createdAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "name": 1 }` (unique)
- `{ "slug": 1 }` (unique)
- `{ "postCount": -1 }`

### 8. AI Suggestions Collection

**Collection Name:** `ai_suggestions`

```json
{
  "_id": "ObjectId",
  "userId": "ObjectId (ref: users, required)",
  "postId": "ObjectId (ref: blog_posts, optional)",
  "suggestionType": "string (enum: 'content', 'title', 'tags', 'improvement', required)",
  "originalContent": "string",
  "suggestedContent": "string (required)",
  "keywords": ["string"],
  "confidence": "number (0-1)",
  "isUsed": "boolean (default: false)",
  "createdAt": "Date (default: now)"
}
```

**Indexes:**

- `{ "userId": 1 }`
- `{ "postId": 1 }`
- `{ "suggestionType": 1 }`
- `{ "isUsed": 1 }`

### 9. User Sessions Collection

**Collection Name:** `user_sessions`

```json
{
  "_id": "ObjectId",
  "userId": "ObjectId (ref: users, required)",
  "sessionId": "string (unique, required)",
  "ipAddress": "string",
  "userAgent": "string",
  "isActive": "boolean (default: true)",
  "lastActivity": "Date",
  "createdAt": "Date (default: now)",
  "expiresAt": "Date (required)"
}
```

**Indexes:**

- `{ "userId": 1 }`
- `{ "sessionId": 1 }` (unique)
- `{ "expiresAt": 1 }` (TTL index)

## Database Configuration

### Connection String Format

```
mongodb+srv://<username>:<password>@<cluster>.mongodb.net/<database>?retryWrites=true&w=majority
```

### Environment Variables

```env
MONGODB_URI=mongodb+srv://username:password@cluster.mongodb.net/blog_platform
MONGODB_DATABASE=blog_platform
JWT_SECRET=your_jwt_secret_key
JWT_REFRESH_SECRET=your_jwt_refresh_secret_key
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=7d
```

## Data Relationships

### One-to-Many Relationships:

- User → Blog Posts (one user can have many posts)
- User → Comments (one user can have many comments)
- Blog Post → Comments (one post can have many comments)
- User → AI Suggestions (one user can have many suggestions)

### Many-to-Many Relationships:

- Users ↔ Blog Posts (through interactions - likes, dislikes, views)
- Blog Posts ↔ Tags (many posts can have many tags)

## Performance Considerations

### Aggregation Pipelines for Popular Posts:

```javascript
db.blog_posts.aggregate([
  { $match: { status: "published" } },
  { $sort: { viewCount: -1, likeCount: -1 } },
  { $limit: 10 },
]);
```

### Text Search Index:

```javascript
db.blog_posts.createIndex({
  title: "text",
  content: "text",
  tags: "text",
});
```

### Compound Indexes for Filtering:

```javascript
db.blog_posts.createIndex({
  status: 1,
  publishedAt: -1,
  category: 1,
});
```

## Security Considerations

1. **Password Hashing**: Use bcrypt with salt rounds of 12
2. **JWT Tokens**: Store refresh tokens in database for revocation
3. **Input Validation**: Validate all user inputs
4. **Rate Limiting**: Implement rate limiting for authentication endpoints
5. **CORS**: Configure CORS properly for web clients

## Backup Strategy

1. **Automated Backups**: Enable MongoDB Atlas automated backups
2. **Point-in-Time Recovery**: Configure for 7-day retention
3. **Manual Exports**: Regular data exports for development

## Monitoring and Analytics

### Key Metrics to Track:

- User registration and login rates
- Blog post creation and view counts
- Popular tags and categories
- User engagement (likes, comments, shares)
- API response times
- Error rates

### MongoDB Atlas Monitoring:

- Enable MongoDB Atlas monitoring
- Set up alerts for:
  - High CPU/Memory usage
  - Slow queries
  - Connection limits
  - Storage thresholds
