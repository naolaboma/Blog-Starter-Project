# MongoDB Atlas Setup Guide for Blog Platform

## 1. MongoDB Atlas Account Setup

### 1.1 Create MongoDB Atlas Account

1. Go to [MongoDB Atlas](https://www.mongodb.com/atlas)
2. Click "Try Free" or "Sign Up"
3. Fill in your details and create an account
4. Verify your email address

### 1.2 Create a New Cluster

1. Click "Build a Database"
2. Choose "FREE" tier (M0 Sandbox)
3. Select your preferred cloud provider (AWS, Google Cloud, or Azure)
4. Choose a region close to your target users
5. Click "Create"

### 1.3 Configure Database Access

1. Go to "Database Access" in the left sidebar
2. Click "Add New Database User"
3. Create a username and password (save these securely)
4. Select "Read and write to any database"
5. Click "Add User"

### 1.4 Configure Network Access

1. Go to "Network Access" in the left sidebar
2. Click "Add IP Address"
3. For development: Click "Allow Access from Anywhere" (0.0.0.0/0)
4. For production: Add specific IP addresses
5. Click "Confirm"

## 2. Database Configuration

### 2.1 Get Connection String

1. Click "Connect" on your cluster
2. Choose "Connect your application"
3. Select "Go" as your driver
4. Copy the connection string

### 2.2 Connection String Format

```
mongodb+srv://<username>:<password>@<cluster>.mongodb.net/<database>?retryWrites=true&w=majority
```

Replace:

- `<username>`: Your database username
- `<password>`: Your database password
- `<cluster>`: Your cluster name
- `<database>`: `blog_platform`

## 3. Environment Configuration

### 3.1 Create .env File

Create a `.env` file in your project root:

```env
# MongoDB Configuration
MONGODB_URI=mongodb+srv://your_username:your_password@your_cluster.mongodb.net/blog_platform?retryWrites=true&w=majority
MONGODB_DATABASE=blog_platform

# JWT Configuration
JWT_SECRET=your_super_secret_jwt_key_here_make_it_long_and_random
JWT_REFRESH_SECRET=your_super_secret_refresh_key_here_make_it_long_and_random
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
```

### 3.2 Security Best Practices

1. **Never commit .env files** to version control
2. Use strong, unique passwords for database users
3. Rotate JWT secrets regularly
4. Use environment-specific configurations

## 4. Database Indexes Setup

### 4.1 Create Indexes via MongoDB Atlas

1. Go to your cluster in MongoDB Atlas
2. Click "Browse Collections"
3. Select your database and collections
4. Create the following indexes:

#### Users Collection Indexes

```javascript
// Email index (unique)
db.users.createIndex({ email: 1 }, { unique: true });

// Username index (unique)
db.users.createIndex({ username: 1 }, { unique: true });

// Role index
db.users.createIndex({ role: 1 });

// Active users index
db.users.createIndex({ isActive: 1 });
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

## 5. Database Initialization

### 5.1 Create Initial Data

Run these commands in MongoDB Atlas shell:

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

## 6. Monitoring and Alerts

### 6.1 Set Up Monitoring

1. Go to "Metrics" in MongoDB Atlas
2. Monitor key metrics:
   - Operations per second
   - Connection count
   - Query performance
   - Storage usage

### 6.2 Configure Alerts

1. Go to "Alerts" in MongoDB Atlas
2. Set up alerts for:
   - High CPU usage (>80%)
   - High memory usage (>80%)
   - Connection count limits
   - Storage thresholds

## 7. Backup Strategy

### 7.1 Automated Backups

1. Go to "Backup" in MongoDB Atlas
2. Enable "Cloud Backup"
3. Configure retention period (7 days for free tier)
4. Set backup schedule

### 7.2 Manual Backups

```bash
# Export database
mongodump --uri="mongodb+srv://username:password@cluster.mongodb.net/blog_platform"

# Import database
mongorestore --uri="mongodb+srv://username:password@cluster.mongodb.net/blog_platform" dump/
```

## 8. Security Checklist

- [ ] Database user has minimal required permissions
- [ ] Network access is restricted to necessary IPs
- [ ] JWT secrets are strong and unique
- [ ] Environment variables are properly configured
- [ ] .env file is in .gitignore
- [ ] Database indexes are created
- [ ] Monitoring and alerts are configured
- [ ] Backup strategy is in place

## 9. Performance Optimization

### 9.1 Connection Pooling

```go
// In your database configuration
clientOptions := options.Client().ApplyURI(uri)
clientOptions.SetMaxPoolSize(100)
clientOptions.SetMinPoolSize(5)
```

### 9.2 Query Optimization

- Use indexes for all frequently queried fields
- Implement pagination for large datasets
- Use projection to limit returned fields
- Implement caching for frequently accessed data

### 9.3 Monitoring Queries

```javascript
// Enable query profiling
db.setProfilingLevel(1, { slowms: 100 });

// View slow queries
db.system.profile.find().sort({ ts: -1 }).limit(10);
```

## 10. Troubleshooting

### 10.1 Common Issues

1. **Connection timeout**: Check network access settings
2. **Authentication failed**: Verify username/password
3. **Index creation failed**: Check if index already exists
4. **Slow queries**: Review and optimize indexes

### 10.2 Useful Commands

```javascript
// Check database stats
db.stats();

// Check collection stats
db.blog_posts.stats();

// List all indexes
db.blog_posts.getIndexes();

// Check current operations
db.currentOp();
```

## 11. Production Considerations

### 11.1 Scaling

- Upgrade to paid tier for better performance
- Use dedicated clusters for production
- Implement read replicas for high read loads
- Consider sharding for very large datasets

### 11.2 Security

- Use VPC peering for network isolation
- Implement encryption at rest
- Use encryption in transit (TLS)
- Regular security audits

### 11.3 Monitoring

- Set up comprehensive monitoring
- Configure automated alerts
- Regular performance reviews
- Capacity planning
