CREATE TABLE messages (
    message_id INT AUTO_INCREMENT PRIMARY KEY,
    consultant_id INTEGER REFERENCES consultants(consultant_id) ON DELETE SET NULL,
    mother_id INTEGER REFERENCES mothers(mother_id) ON DELETE SET NULL,
    message TEXT NOT NULL,
    status VARCHAR(50) DEFAULT 'unread',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


