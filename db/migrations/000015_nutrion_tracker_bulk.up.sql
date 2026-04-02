CREATE TABLE questions (
    question_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    question_text VARCHAR(500) NOT NULL,
    question_key VARCHAR(100) UNIQUE NOT NULL, -- e.g., "protein_animal", "water_intake"
    category VARCHAR(50), -- e.g.,daily, weekly, montly
    input_type VARCHAR(50) CHECK (input_type IN (
        'single_choice', 
        'multiple_choice', 
        'number', 
        'boolean', 
        'text',
        'image_upload'
    )),
    unit VARCHAR(50), -- e.g., "glasses", "grams", "times"
    is_required BOOLEAN DEFAULT TRUE,
    display_order INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE question_options (
    option_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    question_id BIGINT UNSIGNED REFERENCES questions(question_id) ON DELETE CASCADE,
    option_value VARCHAR(100) NOT NULL, 
    option_label VARCHAR(100) NOT NULL, 
    icon_emoji VARCHAR(50), 
    display_order INTEGER,
    is_default BOOLEAN DEFAULT FALSE
);

CREATE TABLE user_tracking_logs (
    tracking_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    mother_id BIGINT UNSIGNED REFERENCES mothers(mother_id),
    pregnancy_id BIGINT UNSIGNED REFERENCES pregnancies(pregnancy_id),
    child_id BIGINT UNSIGNED REFERENCES children(child_id),
    tracking_date DATE NOT NULL,
    frequency VARCHAR(20) CHECK (frequency IN ('daily', 'weekly', 'monthly')),
    
    -- The JSON Object holds all dynamically fetched answers
    -- e.g., {"water_glasses": 5, "iron_pill_taken": true, "hemoglobin_level": 12.5}
    answers_data JSON NOT NULL, 
    
    completion_status VARCHAR(20),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(mother_id, pregnancy_id, child_id, tracking_date, frequency)
);