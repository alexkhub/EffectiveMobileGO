CREATE TABLE IF NOT EXISTS subsription(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    service_name TEXT, 
    price INT,
    user_id TEXT,
    start_date TIMESTAMP 
);