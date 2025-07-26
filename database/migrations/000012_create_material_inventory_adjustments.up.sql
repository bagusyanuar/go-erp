-- Enable UUID support (jika belum)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Buat enum type untuk type field
DO $$ BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'adjustment_type') THEN
        CREATE TYPE adjustment_type AS ENUM ('in', 'out');
    END IF;
END $$;

-- Create material_inventory_adjustments table
CREATE TABLE material_inventory_adjustments ( 
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    material_id UUID,                                    -- Foreign key untuk material, nullable
    unit_id UUID,
    date DATE NOT NULL,
    type adjustment_type NOT NULL,
    quantity NUMERIC(15, 2) DEFAULT 0,
    author_id UUID,
    created_at TIMESTAMPTZ DEFAULT now(),                -- Timestamp otomatis
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    -- Foreign key constraint
    CONSTRAINT fk_material_inventory_adjustments_material FOREIGN KEY (material_id)
        REFERENCES materials(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_material_inventory_adjustments_unit FOREIGN KEY (unit_id)
        REFERENCES units(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_material_inventory_adjustments_author FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE SET NULL
);

-- Optional: Index untuk soft delete
CREATE INDEX idx_material_inventory_adjustments_deleted_at ON material_inventory_adjustments(deleted_at);