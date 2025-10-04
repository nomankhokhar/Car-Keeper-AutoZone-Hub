-- Drop old tables
DROP TABLE IF EXISTS cars CASCADE;
DROP TABLE IF EXISTS engines CASCADE;

-- Create engines table
CREATE TABLE engines (
    engine_id UUID PRIMARY KEY,
    displacement BIGINT NOT NULL,
    no_of_cylinders BIGINT NOT NULL,
    car_range BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL
);

CREATE INDEX idx_engines_deleted_at ON engines(deleted_at);

-- Create cars table
CREATE TABLE cars (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    fuel_type VARCHAR(50) NOT NULL,
    engine_id UUID NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ NULL,
    CONSTRAINT fk_cars_engine FOREIGN KEY (engine_id) 
        REFERENCES engines(engine_id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);

CREATE INDEX idx_cars_deleted_at ON cars(deleted_at);

-- Insert dummy engine data
INSERT INTO engines (engine_id, displacement, no_of_cylinders, car_range, created_at, updated_at)
VALUES
    ('e1f86b1a-0873-4c19-bae2-fc60329d0140'::uuid, 2000, 4, 600, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('f4a9c66b-8e38-419b-93c4-215d5cefb318'::uuid, 1600, 4, 550, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c'::uuid, 3000, 6, 700, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('9746be12-07b7-42a3-b8ab-7d1f209b63d7'::uuid, 1800, 4, 500, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert dummy car data
INSERT INTO cars (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at)
VALUES
    ('c7c1a6d5-1ec4-4c64-a59a-8a2f6f3d2bf3'::uuid, 'Honda Civic', '2023', 'Honda', 'Gasoline', 'e1f86b1a-0873-4c19-bae2-fc60329d0140'::uuid, 25000.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('9d6a56f8-79c3-4931-a5c0-6b290c84ba2f'::uuid, 'Toyota Corolla', '2022', 'Toyota', 'Gasoline', 'f4a9c66b-8e38-419b-93c4-215d5cefb318'::uuid, 22000.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('9b9437c4-3ed1-45a5-b240-0fe3e24e0e4e'::uuid, 'Ford Mustang', '2024', 'Ford', 'Gasoline', 'cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c'::uuid, 40000.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
    ('5e9df51a-8d7a-4d84-9c58-4ccfe5c7db06'::uuid, 'BMW 3 Series', '2023', 'BMW', 'Gasoline', '9746be12-07b7-42a3-b8ab-7d1f209b63d7'::uuid, 35000.00, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
