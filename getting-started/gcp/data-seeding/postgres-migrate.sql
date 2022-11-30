CREATE TABLE IF NOT EXISTS tasks (
    task_id INT PRIMARY KEY NOT NULL,
    title TEXT NOT NULL,
    start_date DATE,
    due_date DATE,
    status INT NOT NULL,
    priority INT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 
