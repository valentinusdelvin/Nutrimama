CREATE TABLE children (
    child_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    pregnancy_id BIGINT UNSIGNED REFERENCES pregnancies(pregnancy_id),
    mother_id BIGINT UNSIGNED REFERENCES mothers(mother_id) ON DELETE CASCADE,
    birth_date DATE NOT NULL,
    birth_weight DECIMAL(5,2),
    birth_length DECIMAL(5,2),
    gender VARCHAR(10),
    apgar_score INTEGER,
    delivery_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE 
);