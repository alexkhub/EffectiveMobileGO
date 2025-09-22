CREATE TABLE IF NOT EXISTS subscription(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name TEXT NOT NULL, 
    price INT NOT NULL,
    user_id UUID NULL,
    start_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);