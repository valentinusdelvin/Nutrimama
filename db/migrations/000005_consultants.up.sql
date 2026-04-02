CREATE TABLE consultants (
    consultant_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    facility_name VARCHAR(255),
    rating DECIMAL(3,2) CHECK (rating >= 0 AND rating <= 5),
    hourly_rate DECIMAL(10,2) DEFAULT 15000
);
