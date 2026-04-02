CREATE TABLE consultation_sessions (
    consultation_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    mother_id INTEGER NOT NULL REFERENCES mothers(mother_id) ON DELETE CASCADE,
    consultant_id INTEGER NOT NULL REFERENCES consultants(consultant_id) ON DELETE CASCADE,
    session_date DATE NOT NULL,
    time_start TIME NOT NULL,
    hour_end TIME NOT NULL,
    total_price DECIMAL(10,2) NOT NULL DEFAULT 0,
    payment_status VARCHAR(20) DEFAULT 'unpaid',
    status VARCHAR(20) DEFAULT 'scheduled',
    CHECK (hour_end > time_start)
);

ALTER TABLE messages ADD COLUMN consultation_id INTEGER REFERENCES consultation_sessions(consultation_id) ON DELETE CASCADE;