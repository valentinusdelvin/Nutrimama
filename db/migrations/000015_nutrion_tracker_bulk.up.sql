CREATE TABLE questions (
    question_id INT AUTO_INCREMENT PRIMARY KEY,
    schedule_id INT REFERENCES tracking_schedules(schedule_id) ON DELETE CASCADE,
    question_text VARCHAR(500) NOT NULL,
    question_key VARCHAR(100) UNIQUE NOT NULL, -- e.g., "protein_animal", "water_intake"
    category VARCHAR(50), -- e.g., "nutrition", "supplement", "symptom"
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
    option_id INT AUTO_INCREMENT PRIMARY KEY,
    question_id INT REFERENCES questions(question_id) ON DELETE CASCADE,
    option_value VARCHAR(100) NOT NULL, 
    option_label VARCHAR(100) NOT NULL, 
    icon_emoji VARCHAR(50), 
    display_order INTEGER,
    is_default BOOLEAN DEFAULT FALSE
);

CREATE TABLE daily_tracking (
    tracking_id INT AUTO_INCREMENT PRIMARY KEY,
    mother_id INT REFERENCES mothers(mother_id),
    pregnancy_id INT REFERENCES pregnancies(pregnancy_id),
    child_id INT REFERENCES children(child_id), 
    tracking_date DATE NOT NULL,
    
    -- Quick access fields (denormalized for reporting)
    protein_animal BOOLEAN DEFAULT FALSE,
    protein_sources VARCHAR(255), 
    water_glasses INTEGER DEFAULT 0,
    iron_pill_taken BOOLEAN DEFAULT FALSE,
    folic_acid_taken BOOLEAN DEFAULT FALSE,
    vitamin_d3_taken BOOLEAN DEFAULT FALSE,
    clinical_symptoms VARCHAR(255),
    
    -- For children
    exclusive_breastfeeding BOOLEAN DEFAULT FALSE,
    mpasi_protein BOOLEAN DEFAULT FALSE,
    added_fats BOOLEAN DEFAULT FALSE,
    carbs_protein_plant BOOLEAN DEFAULT FALSE,
    vegetables_fruits BOOLEAN DEFAULT FALSE,
    iodized_salt BOOLEAN DEFAULT FALSE,
    iron_drops BOOLEAN DEFAULT FALSE,
    bowel_urine_normal BOOLEAN DEFAULT FALSE,
    
    -- Metadata
    completion_status VARCHAR(20) CHECK (completion_status IN ('incomplete', 'partial', 'complete')),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(mother_id, pregnancy_id, child_id, tracking_date)
);

CREATE TABLE weekly_tracking (
    tracking_id INT AUTO_INCREMENT PRIMARY KEY,
    mother_id INT REFERENCES mothers(mother_id),
    pregnancy_id INT REFERENCES pregnancies(pregnancy_id),
    child_id INT REFERENCES children(child_id),
    week_start_date DATE NOT NULL,
    week_number INTEGER NOT NULL,
    
    -- Pregnant women
    omega3_consumed BOOLEAN DEFAULT FALSE,
    omega3_source VARCHAR(255),
    self_weighing_done BOOLEAN DEFAULT FALSE,
    weight_recorded DECIMAL(5,2),
    light_exercise_done BOOLEAN DEFAULT FALSE,
    exercise_minutes INTEGER,
    
    -- Children
    food_texture_evaluated BOOLEAN DEFAULT FALSE,
    texture_level VARCHAR(20), -- 'smooth', 'mashed', 'chopped', 'family'
    motoric_stimulation_done BOOLEAN DEFAULT FALSE,
    cognitive_stimulation_done BOOLEAN DEFAULT FALSE,
    
    completion_status VARCHAR(20),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(mother_id, pregnancy_id, child_id, week_start_date)
);

CREATE TABLE monthly_tracking (
    tracking_id INT AUTO_INCREMENT PRIMARY KEY,
    mother_id INT REFERENCES mothers(mother_id),
    pregnancy_id INT REFERENCES pregnancies(pregnancy_id),
    child_id INT REFERENCES children(child_id),
    month_date DATE NOT NULL,
    
    -- Pregnant women
    calcium_supplement BOOLEAN DEFAULT FALSE,
    lila_measured DECIMAL(5,2), 
    hemoglobin_level DECIMAL(4,2), 
    anc_visit_done BOOLEAN DEFAULT FALSE,
    blood_pressure_systolic INTEGER,
    blood_pressure_diastolic INTEGER,
    
    -- Children
    vitamin_a_given BOOLEAN DEFAULT FALSE, 
    weight_measured DECIMAL(5,2), 
    height_measured DECIMAL(5,2), 
    head_circumference DECIMAL(5,2), 
    immunization_up_to_date BOOLEAN DEFAULT FALSE,
    immunizations_given VARCHAR(255),
    child_hemoglobin DECIMAL(4,2),
    
    completion_status VARCHAR(20),
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    UNIQUE(mother_id, pregnancy_id, child_id, month_date)
);