-- This is the SQL script that will be used to initialize the database schema.

-- THIS IS SCRIPT FOR CREATING USERS TABLE
CREATE TABLE users (
	uuid UUID PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password TEXT NOT NULL,
    is_premium BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT unique_username UNIQUE (username),
    CONSTRAINT unique_email UNIQUE (email)
);

-- THIS IS SCRIPT FOR CREATING SWIPE TABLE
CREATE TABLE swipes (
    uuid UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    target_id UUID NOT NULL,
    direction VARCHAR(10) NOT NULL CHECK (direction IN ('left', 'right')),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    UNIQUE (user_id, target_id)
);


-- FUNCTION FOR LIMITING SWIPE IF USER IS NOT PREMIUM
CREATE OR REPLACE FUNCTION limit_swipes_per_day()
RETURNS TRIGGER AS $$
DECLARE
    swipe_count INTEGER;
    is_user_premium BOOLEAN;
BEGIN
    SELECT is_premium INTO is_user_premium FROM users WHERE uuid = NEW.user_id;
    IF NOT is_user_premium THEN
        SELECT COUNT(*) INTO swipe_count
        FROM swipes
        WHERE user_id = NEW.user_id
          AND DATE(created_at) = CURRENT_DATE;
        IF swipe_count >= 10 THEN
            RAISE EXCEPTION 'Non-premium users can only swipe 10 times per day.';
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- TRIGGER FUNCTION
CREATE TRIGGER check_swipe_limit
BEFORE INSERT ON swipes
FOR EACH ROW
EXECUTE FUNCTION limit_swipes_per_day();

CREATE INDEX idx_swipes_user_created_at ON swipes(user_id, created_at);

-- INSERT 20 DUMMY DATA
INSERT INTO users (uuid, email, username, password, is_premium, created_at)
VALUES
    ('1f5ed1f0-91e6-4f39-a726-b7e48cb8f4e2', 'alice.smith@example.com', 'alicesmith', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('db4c207b-f26f-48a3-8252-bd38bfa5e158', 'bob.jones@example.com', 'bobjones', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('8a6878b4-6ac4-4324-85ed-8d6b55960d4f', 'charlie.brown@example.com', 'charliebrown', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('a5c942fa-e8f3-4427-b1f1-51007e25e684', 'david.green@example.com', 'davidgreen', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('1234f598-e503-4b2e-bd02-7c321bc84e10', 'emma.white@example.com', 'emmawhite', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('c84019ad-5b52-4f8b-a567-e441649fbf8b', 'frank.black@example.com', 'frankblack', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('9b91c299-ef62-4b96-9123-d5132020507b', 'grace.kim@example.com', 'gracekim', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('6da0b2b3-fb89-4173-b28d-16b482132d19', 'hank.miller@example.com', 'hankmiller', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('3471bb9e-8f88-43d7-8fdb-f9d47a4266ff', 'isabel.jones@example.com', 'isabeljones', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('d56890be-bb2b-4fbb-a110-0c8f3b580c83', 'jack.smith@example.com', 'jacksmith', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('16f9f6d0-40c2-4ea9-b7ac-60d5f83fa5cf', 'karen.taylor@example.com', 'karentaylor', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('70e3f4cb-5e69-4aaf-99ed-4ed3a8e228ec', 'lisa.martin@example.com', 'lisamartin', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('ed20e3c3-0334-4049-8e23-5f659e68dbb0', 'mike.davis@example.com', 'mikedavis', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('4c2399f0-9e7f-4b27-87fd-d87ccfc31cba', 'nina.williams@example.com', 'ninawilliams', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('a1c2f129-8195-44c4-8196-590f60c0f3ad', 'olivia.roberts@example.com', 'oliviaroberts', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('bba4f4f4-9481-4a19-b237-1f7bc859b99d', 'paul.johnson@example.com', 'pauljohnson', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('49fbe501-2c61-4530-a1a5-dcb1499029e9', 'quincy.lee@example.com', 'quincylee', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW()),
    ('6f3ff8e4-036f-4f9a-853f-bb7f8580e1de', 'rachel.adams@example.com', 'racheladams', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', FALSE, NOW()),
    ('1c1a4387-2196-442f-a073-44146f6e396a', 'steve.taylor@example.com', 'stevetaylor', '$2a$10$PX3f2W62e515xHqgY7TX6Oq6lUEA/58OedkgBBR4v3HbY8/BWNIlW', TRUE, NOW());