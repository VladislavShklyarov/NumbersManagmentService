CREATE TABLE IF NOT EXISTS phone_numbers (
    number TEXT PRIMARY KEY,
    country TEXT,
    region TEXT,
    provider TEXT,
    source TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_country ON phone_numbers(country);
CREATE INDEX idx_region ON phone_numbers(region);
CREATE INDEX idx_provider ON phone_numbers(provider);

