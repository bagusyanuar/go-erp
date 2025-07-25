-- Enable UUID support
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create material_categories table
CREATE TABLE material_categories (
    material_id UUID NOT NULL,
    category_id UUID NOT NULL,

    PRIMARY KEY (material_id, category_id),

    CONSTRAINT fk_material FOREIGN KEY (material_id) REFERENCES materials(id) ON DELETE CASCADE,
    CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);