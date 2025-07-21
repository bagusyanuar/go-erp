-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create features table
CREATE TABLE features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);


-- Trigger function to update updated_at
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger on update
CREATE TRIGGER set_timestamp
BEFORE UPDATE ON features
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();