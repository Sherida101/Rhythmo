-- This script initialises the database for the Habit Tracker application.
-- It creates the necessary tables and populates them with initial data.
-- Add permission: chmod 775 ./scripts/init.sql

-- Create priority table
CREATE TABLE IF NOT EXISTS priority (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);

-- Create habit categories table
CREATE TABLE IF NOT EXISTS habit_categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE,
  colour VARCHAR(7) DEFAULT '#FFFFFF',
  icon VARCHAR(50),
);

-- Define status enum type
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'habit_status') THEN
        CREATE TYPE habit_status AS ENUM ('active', 'paused', 'completed', 'abadoned');
    END IF;
END
$$;

-- Create habits table
CREATE TABLE IF NOT EXISTS habits (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    lastDone DATE,
    streak INTEGER DEFAULT 0,
    frequency VARCHAR(50) NOT NULL,
    startDate DATE NOT NULL,
    endDate DATE,
    description TEXT,
    category_id INTEGER REFERENCES habit_categories(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    status habit_status DEFAULT 'active',
    reminder_time TIME,
    priority INTEGER REFERENCES priority(id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO priority (name) VALUES
('Low'),
('Medium'),
('High');

/* Populate categories */
INSERT INTO habit_categories (name, colour, icon) VALUES
('Health', '#FF5733', 'health-icon'),
('Fitness', '#33FF57', 'fitness-icon'),
('Productivity', '#3357FF', 'productivity-icon'),
('Wellness', '#FF33A1', 'wellness-icon'),
('Learning', '#FF8C33', 'learning-icon'),
('Finance', '#33FFF5', 'finance-icon'),
('Social', '#F533FF', 'social-icon'),
('Hobbies', '#F5FF33', 'hobbies-icon'),
('Personal Development', '#33FF8C', 'personal-development-icon');

/* Populate users */
/* Note: Passwords should be hashed in a real application */
INSERT INTO users (username, password_hash, email) VALUES
('janeDoe101', 'password', 'jane.doe@example.com');

/* Populate habits */
INSERT INTO habits (name, lastDone, streak, frequency, startDate, endDate, description, category_id, user_id, status, reminder_time, priority) VALUES
('Drink Water', '2023-10-01', 5, 'daily', '2023-09-01', NULL, 'Stay hydrated!', 1, 1, 'active', '08:00:00', 1),
('Exercise', '2023-10-02', 10, 'weekly', '2023-09-15', NULL, 'Stay fit and healthy!', 2, 1, 'active', '07:00:00', 2),
('Read a Book', '2023-10-03', 7, 'weekly', '2023-09-20', NULL, 'Expand your knowledge!', 5, 1, 'active', '20:00:00', 3),
('Meditate', '2023-10-01', 3, 'daily', '2023-09-10', NULL, 'Clear your mind!', 4, 1, 'active', '06:00:00', 1),
('Learn a Language', '2023-10-02', 2, 'weekly', '2023-09-25', NULL, 'Expand your horizons!', 5, 1, 'active', '19:00:00', 2);
