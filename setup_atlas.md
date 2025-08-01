# MongoDB Atlas Setup for Your Cluster

## Your Atlas Connection Details

**Cluster:** `cluster0.mijqs0y.mongodb.net`  
**Username:** `naolaboma`  
**Database:** `blog_platform`

## Step 1: Create .env File

Create a `.env` file in your project root with the following content:

```env
# MongoDB Configuration
MONGODB_URI=mongodb+srv://naolaboma:<db_password>@cluster0.mijqs0y.mongodb.net/blog_platform?retryWrites=true&w=majority&appName=Cluster0
MONGODB_DATABASE=blog_platform

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here_make_it_long_and_random_at_least_32_characters
JWT_REFRESH_SECRET=your_super_secret_refresh_key_here_make_it_long_and_random_at_least_32_characters
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=7d

# Server Configuration
PORT=8080
ENVIRONMENT=development

# Email Configuration (for password reset and verification)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password
FROM_EMAIL=noreply@yourblog.com

# AI Configuration (if using external AI services)
OPENAI_API_KEY=your_openai_api_key

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:8080

# Rate Limiting
RATE_LIMIT_REQUESTS=100
RATE_LIMIT_DURATION=1m
```

**Important:** Replace `<db_password>` with your actual MongoDB Atlas password.

## Step 2: Generate Secure JWT Secrets

Generate secure JWT secrets using this command:

```bash
# Generate JWT secret
openssl rand -base64 32

# Generate JWT refresh secret
openssl rand -base64 32
```

Replace the JWT_SECRET and JWT_REFRESH_SECRET values in your .env file with the generated secrets.

## Step 3: Set Up Database Collections and Indexes

### 3.1 Access MongoDB Atlas Shell

1. Go to your MongoDB Atlas dashboard
2. Click on your cluster
3. Click "Browse Collections"
4. Create a new database called `blog_platform`

### 3.2 Create Collections

Create these collections in your `blog_platform` database:

- `users`
- `blog_posts`
- `comments`
- `user_interactions`
- `auth_tokens`
- `tags`
- `categories`
- `ai_suggestions`
- `user_sessions`

### 3.3 Create Indexes

Run these commands in the MongoDB Atlas shell:

#### Users Collection Indexes

```javascript
use blog_platform

// Email index (unique)
db.users.createIndex({ "email": 1 }, { unique: true })

// Username index (unique)
db.users.createIndex({ "username": 1 }, { unique: true })

// Role index
db.users.createIndex({ "role": 1 })

// Active users index
db.users.createIndex({ "isActive": 1 })
```

#### Blog Posts Collection Indexes

```javascript
// Author index
db.blog_posts.createIndex({ authorId: 1 });

// Status index
db.blog_posts.createIndex({ status: 1 });

// Published date index
db.blog_posts.createIndex({ publishedAt: -1 });

// View count index
db.blog_posts.createIndex({ viewCount: -1 });

// Like count index
db.blog_posts.createIndex({ likeCount: -1 });

// Tags index
db.blog_posts.createIndex({ tags: 1 });

// Category index
db.blog_posts.createIndex({ category: 1 });

// Slug index (unique)
db.blog_posts.createIndex({ slug: 1 }, { unique: true });

// Text search index
db.blog_posts.createIndex({
  title: "text",
  content: "text",
  tags: "text",
});
```

#### Comments Collection Indexes

```javascript
// Post ID index
db.comments.createIndex({ postId: 1 });

// Author index
db.comments.createIndex({ authorId: 1 });

// Parent comment index
db.comments.createIndex({ parentCommentId: 1 });

// Created date index
db.comments.createIndex({ createdAt: -1 });
```

#### User Interactions Collection Indexes

```javascript
// Compound unique index
db.user_interactions.createIndex(
  { userId: 1, postId: 1, interactionType: 1 },
  { unique: true }
);

// Post interactions index
db.user_interactions.createIndex({ postId: 1, interactionType: 1 });

// User interactions index
db.user_interactions.createIndex({ userId: 1 });
```

#### Auth Tokens Collection Indexes

```javascript
// User and token type index
db.auth_tokens.createIndex({ userId: 1, tokenType: 1 });

// Token index (unique)
db.auth_tokens.createIndex({ token: 1 }, { unique: true });

// TTL index for automatic expiration
db.auth_tokens.createIndex({ expiresAt: 1 }, { expireAfterSeconds: 0 });
```

#### Tags Collection Indexes

```javascript
// Name index (unique)
db.tags.createIndex({ name: 1 }, { unique: true });

// Post count index
db.tags.createIndex({ postCount: -1 });
```

#### Categories Collection Indexes

```javascript
// Name index (unique)
db.categories.createIndex({ name: 1 }, { unique: true });

// Slug index (unique)
db.categories.createIndex({ slug: 1 }, { unique: true });

// Post count index
db.categories.createIndex({ postCount: -1 });
```

#### AI Suggestions Collection Indexes

```javascript
// User index
db.ai_suggestions.createIndex({ userId: 1 });

// Post index
db.ai_suggestions.createIndex({ postId: 1 });

// Suggestion type index
db.ai_suggestions.createIndex({ suggestionType: 1 });

// Used index
db.ai_suggestions.createIndex({ isUsed: 1 });
```

#### User Sessions Collection Indexes

```javascript
// User index
db.user_sessions.createIndex({ userId: 1 });

// Session ID index (unique)
db.user_sessions.createIndex({ sessionId: 1 }, { unique: true });

// TTL index for automatic expiration
db.user_sessions.createIndex({ expiresAt: 1 }, { expireAfterSeconds: 0 });
```

## Step 4: Initialize Default Data

Run these commands to create initial categories and tags:

```javascript
// Create default categories
db.categories.insertMany([
  {
    name: "Technology",
    description: "Technology related articles",
    slug: "technology",
    postCount: 0,
    createdAt: new Date(),
  },
  {
    name: "Lifestyle",
    description: "Lifestyle and personal development",
    slug: "lifestyle",
    postCount: 0,
    createdAt: new Date(),
  },
  {
    name: "Business",
    description: "Business and entrepreneurship",
    slug: "business",
    postCount: 0,
    createdAt: new Date(),
  },
]);

// Create default tags
db.tags.insertMany([
  {
    name: "golang",
    description: "Go programming language",
    postCount: 0,
    createdAt: new Date(),
  },
  {
    name: "mongodb",
    description: "MongoDB database",
    postCount: 0,
    createdAt: new Date(),
  },
  {
    name: "api",
    description: "Application Programming Interface",
    postCount: 0,
    createdAt: new Date(),
  },
]);
```

## Step 5: Test Connection

Run the setup script to test your connection:

```bash
chmod +x scripts/setup.sh
./scripts/setup.sh
```

## Step 6: Start the Application

```bash
go run main.go
```

## Verification Checklist

- [ ] .env file created with correct MongoDB URI
- [ ] JWT secrets generated and configured
- [ ] All collections created in blog_platform database
- [ ] All indexes created successfully
- [ ] Default categories and tags inserted
- [ ] Application starts without errors
- [ ] Database connection successful

## Troubleshooting

### Connection Issues

- Verify your password is correct in the MONGODB_URI
- Check that your IP address is whitelisted in MongoDB Atlas
- Ensure the database user has read/write permissions

### Index Creation Issues

- Some indexes might already exist, that's okay
- Check the MongoDB Atlas logs for any errors
- Verify you're in the correct database (`use blog_platform`)

### Application Issues

- Make sure all Go dependencies are installed: `go mod download`
- Check that the .env file is in the project root
- Verify all environment variables are set correctly
