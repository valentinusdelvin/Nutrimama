CREATE TABLE pregnancies (
    pregnancy_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    mother_id BIGINT UNSIGNED NOT NULL REFERENCES mothers(mother_id) ON DELETE CASCADE,
    lpm DATE,
    starting_weight DECIMAL(6,2),
    fetus_count INTEGER DEFAULT 1,
    status VARCHAR(50) DEFAULT 'active'
);

