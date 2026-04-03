CREATE TABLE educational_tools (
    edu_tools_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    publisher VARCHAR(255),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    category VARCHAR(100),
    thumbnail VARCHAR(255),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);