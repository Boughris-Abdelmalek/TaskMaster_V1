-- LOAD DATA FOR TODOS
INSERT INTO todos(title, description, completed)
VALUES
    ('Complete Project Proposal', 'Write and finalize the project proposal document.', false),
    ('Research Literature', 'Gather relevant research papers and articles for the project.', true),
    ('Design UI Mockups', 'Create mockups for the user interface of the application.', false),
    ('Set Up Development Environment', 'Install and configure the required development tools.', false),
    ('Implement User Authentication', 'Develop user authentication functionality using secure methods.', true),
    ('Write API Documentation', 'Document the endpoints.ts and usage of the applications API.', false),
    ('Test Application Components', 'Perform unit and integration testing of key components.', false),
    ('Bug Fixes and Optimization', 'Address any identified bugs and optimize the performance.', false),
    ('Create User Guide', 'Prepare a comprehensive guide on how to use the application.', false),
    ('Deploy to Production', 'Deploy the application to the production server for public use.', false);

-- CREATE PREDEFINED USERS
INSERT INTO users(first_name, last_name, user_name, password, user_type, email, created_at)
VALUES
    ('John', 'Doe', 'johndoe', 'password123', 'admin', 'johndoe@example.com', NOW()),
    ('Alice', 'Smith', 'alicesmith', 'pass456', 'regular', 'alice@example.com', NOW()),
    ('Bob', 'Johnson', 'bobj', 'mysecret', 'regular', 'bob@example.com', NOW()),
    ('Eva', 'Williams', 'evaw', 'secure123', 'admin', 'eva@example.com', NOW()),
    ('Mike', 'Brown', 'mikeb', 'topsecret', 'regular', 'mike@example.com', NOW());