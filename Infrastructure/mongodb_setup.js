// MongoDB Setup Script for Blog API
// Run this script to set up the database locally
// Usage: mongo < mongodb_setup.js

// Switch to blog_db database (creates it if it doesn't exist)
use blog_db;

print("Setting up Blog API Database...");

// Create users collection with indexes
db.createCollection("users");
db.users.createIndex({ "username": 1 }, { unique: true });
db.users.createIndex({ "email": 1 }, { unique: true });
db.users.createIndex({ "role": 1 });

print("Users collection created with indexes");

// Create blogs collection with indexes
db.createCollection("blogs");
db.blogs.createIndex({ "author_id": 1 });
db.blogs.createIndex({ "author_username": 1 });
db.blogs.createIndex({ "tags": 1 });
db.blogs.createIndex({ "created_at": -1 });
db.blogs.createIndex({ "view_count": -1 });
db.blogs.createIndex({ "title": "text", "content": "text" });

print("Blogs collection created with indexes");

// Create sessions collection with indexes
db.createCollection("sessions");
db.sessions.createIndex({ "username": 1 });
db.sessions.createIndex({ "refresh_token": 1 });
db.sessions.createIndex({ "verification_token": 1 });
db.sessions.createIndex({ "password_reset_token": 1 });
db.sessions.createIndex({ "expires_at": 1 });

print("Sessions collection created with indexes");

// Create reactions collection with indexes
db.createCollection("reactions");
db.reactions.createIndex({ "blog_id": 1, "user_id": 1 }, { unique: true });
db.reactions.createIndex({ "blog_id": 1, "reaction_type": 1 });
db.reactions.createIndex({ "user_id": 1 });

print("Reactions collection created with indexes");

// Create password reset tokens collection with indexes
db.createCollection("password_reset_tokens");
db.password_reset_tokens.createIndex({ "user_id": 1 });
db.password_reset_tokens.createIndex({ "token": 1 }, { unique: true });
db.password_reset_tokens.createIndex({ "expires_at": 1 });

print("Password reset tokens collection created with indexes");

// Create tags collection with indexes
db.createCollection("tags");
db.tags.createIndex({ "name": 1 }, { unique: true });

print("Tags collection created with indexes");

// Insert sample data (optional)
print("Inserting sample data...");

// Sample user
db.users.insertOne({
    username: "admin",
    email: "admin@example.com",
    password: "$2a$10$hashedpassword", // This will be hashed by the app
    role: "admin",
    profile_picture: {
        filename: "",
        file_path: "",
        public_id: "",
        uploaded_at: new Date()
    },
    bio: "System Administrator",
    created_at: new Date(),
    updated_at: new Date()
});

// Sample blog
db.blogs.insertOne({
    title: "Welcome to Blog API",
    content: "This is a sample blog post to test the API.",
    author_id: db.users.findOne({username: "admin"})._id,
    author_username: "admin",
    tags: ["welcome", "sample"],
    view_count: 0,
    like_count: 0,
    comment_count: 0,
    likes: [],
    dislikes: [],
    comments: [],
    created_at: new Date(),
    updated_at: new Date()
});

// Insert sample tags
db.tags.insertMany([
    { name: "technology", created_at: new Date() },
    { name: "programming", created_at: new Date() },
    { name: "golang", created_at: new Date() },
    { name: "web-development", created_at: new Date() },
    { name: "database", created_at: new Date() },
    { name: "api", created_at: new Date() },
    { name: "tutorial", created_at: new Date() },
    { name: "best-practices", created_at: new Date() }
]);

print("Sample data inserted");

// Show collections
print("\nDatabase Collections:");
db.getCollectionNames().forEach(function(collection) {
    print("  - " + collection);
});

print("\nMongoDB setup completed successfully!");
print("You can now run the Blog API application."); 