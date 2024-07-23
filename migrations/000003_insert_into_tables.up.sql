-- 002_insert_data.sql
INSERT INTO users (passport_series, passport_number, surname, name, patronymic, address, created_at, updated_at) VALUES
    ('1234', '123456', 'Ivanov', 'Ivan', 'Ivanovich', '123 Main St', NOW(), NOW()),
    ('2345', '234567', 'Petrov', 'Petr', 'Petrovich', '456 Elm St', NOW(), NOW()),
    ('3456', '345678', 'Sidorov', 'Sidr', 'Sidorovich', '789 Oak St', NOW(), NOW()),
    ('5678', '456789', 'Kuznetsov', 'Nikolai', 'Nikolaevich', '101 Pine St', NOW(), NOW()),
    ('7890', '567890', 'Smirnov', 'Sergei', 'Sergeevich', '202 Maple St', NOW(), NOW()),
    ('9012', '678901', 'Popov', 'Alexander', 'Alexandrovich', '303 Birch St', NOW(), NOW()),
    ('1234', '789012', 'Vasiliev', 'Vasiliy', 'Vasilievich', '404 Cedar St', NOW(), NOW()),
    ('3456', '890123', 'Mikhailov', 'Mikhail', 'Mikhailovich', '505 Cherry St', NOW(), NOW()),
    ('5678', '901234', 'Fedorov', 'Fedor', 'Fedorovich', '606 Walnut St', NOW(), NOW()),
    ('7890', '012345', 'Nikolaev', 'Nikolai', 'Nikolaevich', '707 Willow St', NOW(), NOW());

INSERT INTO tasks (user_id, description, start_time, end_time, duration, created_at, updated_at) VALUES
    (1, 'Task 1', NOW(), NOW() + INTERVAL '1 hour', '360', NOW(), NOW()),
    (2, 'Task 2', NOW(), NOW() + INTERVAL '2 hours', '7200', NOW(), NOW()),
    (3, 'Task 3', NOW(), NOW() + INTERVAL '3 hours', '10800', NOW(), NOW()),
    (4, 'Task 4', NOW(), NOW() + INTERVAL '4 hours', '14400', NOW(), NOW()),
    (5, 'Task 5', NOW(), NOW() + INTERVAL '5 hours', '18000', NOW(), NOW()),
    (6, 'Task 6', NOW(), NOW() + INTERVAL '6 hours', '21600', NOW(), NOW()),
    (7, 'Task 7', NOW(), NOW() + INTERVAL '7 hours', '25200', NOW(), NOW()),
    (8, 'Task 8', NOW(), NOW() + INTERVAL '8 hours', '28800', NOW(), NOW()),
    (9, 'Task 9', NOW(), NOW() + INTERVAL '9 hours', '32400', NOW(), NOW()),
    (10, 'Task 10', NOW(), NOW() + INTERVAL '10 hours', '36000', NOW(), NOW()),
    (1, 'Task 11', NOW(), NOW() + INTERVAL '1 hour', '3600', NOW(), NOW()),
    (2, 'Task 12', NOW(), NOW() + INTERVAL '2 hours', '720', NOW(), NOW()),
    (3, 'Task 13', NOW(), NOW() + INTERVAL '3 hours', '10800', NOW(), NOW()),
    (4, 'Task 14', NOW(), NOW() + INTERVAL '4 hours', '14400', NOW(), NOW()),
    (5, 'Task 15', NOW(), NOW() + INTERVAL '5 hours', '18000', NOW(), NOW());

