CREATE TABLE IF NOT EXISTS it03 (
    id  SERIAL PRIMARY KEY,
    name varchar(200) NOT NULL,
    reason TEXT NOT NULL,
    status_id int2 NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (status_id) REFERENCES it03_statuses(id)
);

CREATE TRIGGER update_timestamp
BEFORE UPDATE ON it03
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

BEGIN;
INSERT INTO it03 (name, reason, status_id) 
VALUES 
    ('รายการที่ 1', 'ไม่อนุมัติสาเหตุที่ 1', 3),
    ('รายการที่ 2', 'ไม่อนุมัติสาเหตุที่ 1', 3),
    ('รายการที่ 3', 'อยุมัติสาเหตุที่ 1', 2),
    ('รายการที่ 4', 'อยุมัติสาเหตุที่ 1', 2),
    ('รายการที่ 5', 'ไม่อนุมัติสาเหตุที่ 2', 3),
    ('รายการที่ 6', '', 1),
    ('รายการที่ 7', '', 1),
    ('รายการที่ 8', '', 1),
    ('รายการที่ 9', '', 1),
    ('รายการที่ 10', '', 1),
    ('รายการที่ 11', '', 1)
ON CONFLICT (id) DO NOTHING;
COMMIT;