CREATE TABLE IF NOT EXISTS tasks (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR,
                                     description VARCHAR,
                                     assigned_date DATE,
                                     is_completed BOOLEAN
);