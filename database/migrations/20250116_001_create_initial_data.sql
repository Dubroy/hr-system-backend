-- 創建部門數據
INSERT INTO departments (name, code, created_at, updated_at) VALUES
    ('Engineering', 'ENG', NOW(), NOW()),
    ('Human Resources', 'HR', NOW(), NOW()),
    ('Marketing', 'MKT', NOW(), NOW()),
    ('Finance', 'FIN', NOW(), NOW()),
    ('Sales', 'SAL', NOW(), NOW());

-- 創建職位數據
INSERT INTO positions (title, level, created_at, updated_at) VALUES
    ('Software Engineer', 3, NOW(), NOW()),
    ('Senior Software Engineer', 4, NOW(), NOW()),
    ('HR Manager', 4, NOW(), NOW()),
    ('Marketing Specialist', 3, NOW(), NOW()),
    ('Financial Analyst', 3, NOW(), NOW()),
    ('Sales Representative', 2, NOW(), NOW());

-- 創建員工數據
INSERT INTO employees (
    employee_code,
    name,
    email,
    phone_number,
    department_id,
    position_id,
    status,
    join_date,
    created_at,
    updated_at
) VALUES
    ('ENG001', 'John Doe', 'john.doe@company.com', '0912345678', 1, 2, 'active', '2023-01-01', NOW(), NOW()),
    ('ENG002', 'Jane Smith', 'jane.smith@company.com', '0923456789', 1, 1, 'active', '2023-02-01', NOW(), NOW()),
    ('HR001', 'Mary Johnson', 'mary.johnson@company.com', '0934567890', 2, 3, 'active', '2023-01-15', NOW(), NOW()),
    ('MKT001', 'Robert Wilson', 'robert.wilson@company.com', '0945678901', 3, 4, 'active', '2023-03-01', NOW(), NOW()),
    ('FIN001', 'Lisa Brown', 'lisa.brown@company.com', '0956789012', 4, 5, 'active', '2023-02-15', NOW(), NOW()),
    ('SAL001', 'Michael Lee', 'michael.lee@company.com', '0967890123', 5, 6, 'active', '2023-04-01', NOW(), NOW()); 