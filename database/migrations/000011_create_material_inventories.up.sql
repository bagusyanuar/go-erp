-- Enable UUID support (jika belum)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create material_inventories table
CREATE TABLE material_inventories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),       -- UUID sebagai primary key
    material_id UUID,                                    -- Foreign key untuk material, nullable
    unit_id UUID,                                        -- Foreign key untuk unit, nullable
    quantity NUMERIC(15, 2) DEFAULT 0,  
    modified_by UUID,                                    -- Foreign key untuk modificator, nullable
    created_at TIMESTAMPTZ DEFAULT now(),                -- Timestamp otomatis
    updated_at TIMESTAMPTZ DEFAULT now(),
    deleted_at TIMESTAMPTZ,                              -- Untuk soft delete, nullable

    -- Foreign key constraint
    CONSTRAINT fk_materials_inventories_material FOREIGN KEY (material_id)
        REFERENCES materials(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_materials_inventories_unit FOREIGN KEY (unit_id)
        REFERENCES units(id)
        ON DELETE SET NULL,
    CONSTRAINT fk_materials_inventories_modified_by FOREIGN KEY (modified_by)
        REFERENCES users(id)
        ON DELETE SET NULL
);

-- Optional: Index untuk soft delete
CREATE INDEX idx_material_inventories_deleted_at ON material_inventories(deleted_at);
