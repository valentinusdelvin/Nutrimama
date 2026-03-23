CREATE TABLE children (
    child_id SERIAL PRIMARY KEY,
    pregnancy_id INT REFERENCES pregnancies(pregnancy_id),
    mother_id INT REFERENCES mothers(mother_id) ON DELETE CASCADE,
    birth_date DATE NOT NULL,
    birth_weight DECIMAL(5,2),
    birth_length DECIMAL(5,2),
    gender VARCHAR(10),
    apgar_score INTEGER,
    delivery_type VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_active BOOLEAN DEFAULT TRUE 
);