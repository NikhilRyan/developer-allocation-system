CREATE TABLE developers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    availability FLOAT DEFAULT 0,
    start_date TIMESTAMP NOT NULL,
    skill_level TEXT,
    system_knowledge FLOAT DEFAULT 0,
    level VARCHAR(50),
    current_workload FLOAT DEFAULT 0,
    on_call_rotation_week BOOLEAN DEFAULT FALSE,
    responsibilities TEXT,
    is_available BOOLEAN DEFAULT TRUE,
    time_zone VARCHAR(100),
    language_proficiency TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
