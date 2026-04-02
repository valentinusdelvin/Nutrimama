CREATE TABLE nutrition_tracking (
    track_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    mother_id BIGINT UNSIGNED REFERENCES mothers(mother_id) ON DELETE CASCADE,
    child_id BIGINT UNSIGNED REFERENCES children(child_id) ON DELETE CASCADE,
    tracking_type VARCHAR(20) CHECK (tracking_type IN ('daily', 'weekly', 'monthly')),
    
    -- Generic JSON Payload securely holding 0-100% computed metrics dynamically
    scores_data JSON NOT NULL,
    
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (mother_id IS NOT NULL OR child_id IS NOT NULL)
);