-- Drop old tables
DROP TABLE IF EXISTS car CASCADE;
DROP TABLE IF EXISTS engine CASCADE;

-- Create engine table
CREATE TABLE engine (
    id UUID PRIMARY KEY,
    displacement INT NOT NULL,
    noofcylinders INT NOT NULL,
    carrange INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create car table
CREATE TABLE car (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    year VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    fuel_type VARCHAR(50) NOT NULL,
    engine_id UUID NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_engine_id FOREIGN KEY (engine_id) REFERENCES engine(id) ON DELETE CASCADE
);

-- Insert dummy data
INSERT INTO engine (id, displacement, noofcylinders, carrange)
VALUES
    ('e1f86b1a-0873-4c19-bae2-fc60329d0140'::uuid, 2000, 4, 600),
    ('f4a9c66b-8e38-419b-93c4-215d5cefb318'::uuid, 1600, 4, 550),
    ('cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c'::uuid, 3000, 6, 700),
    ('9746be12-07b7-42a3-b8ab-7d1f209b63d7'::uuid, 1800, 4, 500);

INSERT INTO car (id, name, year, brand, fuel_type, engine_id, price)
VALUES
    ('c7c1a6d5-1ec4-4c64-a59a-8a2f6f3d2bf3'::uuid, 'Honda Civic', '2023', 'Honda', 'Gasoline', 'e1f86b1a-0873-4c19-bae2-fc60329d0140'::uuid, 25000.00),
    ('9d6a56f8-79c3-4931-a5c0-6b290c84ba2f'::uuid, 'Toyota Corolla', '2022', 'Toyota', 'Gasoline', 'f4a9c66b-8e38-419b-93c4-215d5cefb318'::uuid, 22000.00),
    ('9b9437c4-3ed1-45a5-b240-0fe3e24e0e4e'::uuid, 'Ford Mustang', '2024', 'Ford', 'Gasoline', 'cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c'::uuid, 40000.00),
    ('5e9df51a-8d7a-4d84-9c58-4ccfe5c7db06'::uuid, 'BMW 3 Series', '2023', 'BMW', 'Gasoline', '9746be12-07b7-42a3-b8ab-7d1f209b63d7'::uuid, 35000.00);
