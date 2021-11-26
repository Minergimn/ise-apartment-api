-- +goose Up
INSERT INTO apartments ("object","owner",status)
VALUES
    ('object1','owner1',0),
    ('object2','owner2',0),
    ('object3','owner2',0),
    ('object4','owner3',0),
    ('object5','owner3',0),
    ('object6','owner6',1),
    ('object7','owner7',0);

INSERT INTO apartments_events
(apartment_id, "type", status, payload, is_deleted, is_locked, updated)
VALUES
   (1, 0, 0, '{"ID": 1, "Owner": "owner1", "Object": "object1", "Status": 0}', false, false, '2021-11-11 23:00:00.000'),
   (2, 0, 0, '{"ID": 2, "Owner": "owner2", "Object": "object2", "Status": 0}', false, false, '2021-11-11 23:00:00.000'),
   (3, 0, 0, '{"ID": 3, "Owner": "owner2", "Object": "object3", "Status": 0}', false, false, '2021-11-11 23:00:00.000'),
   (4, 0, 0, '{"ID": 4, "Owner": "owner3", "Object": "object4", "Status": 0}', false, false, '2021-11-11 23:00:00.000'),
   (5, 0, 0, '{"ID": 5, "Owner": "owner4", "Object": "object5", "Status": 0}', false, false, '2021-11-11 23:00:00.000'),
   (6, 0, 0, '{"ID": 6, "Owner": "owner6", "Object": "object6", "Status": 0}', true, false, '2021-11-11 23:00:00.000'),
   (7, 0, 0, '{"ID": 7, "Owner": "owner7", "Object": "object7", "Status": 0}', false, false, '2021-11-11 23:00:00.000');

-- +goose Down
TRUNCATE apartments;
TRUNCATE apartments_events;
