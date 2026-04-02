CREATE TABLE foods (
    food_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    carbohydrates DECIMAL(6,2),
    fat DECIMAL(6,2),
    category VARCHAR(100),
    nutrition_id INTEGER REFERENCES nutrition_tracking(track_id) ON DELETE SET NULL
);