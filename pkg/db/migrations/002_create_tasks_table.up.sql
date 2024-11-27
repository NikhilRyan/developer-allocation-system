CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    estimation FLOAT DEFAULT 0,
    required_skills TEXT,
    required_knowledge FLOAT DEFAULT 0,
    priority VARCHAR(50),
    delivery_date TIMESTAMP,
    status VARCHAR(50) DEFAULT 'Not Started',
    dependencies TEXT,
    stakeholders TEXT,
    risk_assessment VARCHAR(50),
    assigned_developer_id INTEGER REFERENCES developers(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
