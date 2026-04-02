CREATE TABLE mothers (
    mother_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,
    blood_type VARCHAR(5),
    full_name VARCHAR(255) NOT NULL,
    birth_date DATE,
    medical_history TEXT,
    height DECIMAL(5,2)
);

