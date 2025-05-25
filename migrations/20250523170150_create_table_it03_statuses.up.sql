CREATE TABLE IF NOT EXISTS it03_statuses (
    id SERIAL PRIMARY KEY,
    status varchar(50) NOT NULL
);

BEGIN;
INSERT INTO it03_statuses (id, status) 
VALUES 
    (1, 'รออนุมัติ'),
    (2, 'อนุมัติ'),
    (3, 'ไม่อนุมัติ')
ON CONFLICT (id) DO NOTHING;
COMMIT;