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
   (1, 'Created', 'Deferred', '{"id": "1", "object": "object1", "owner": "owner1"}', false, false, '2021-11-11 23:00:00.000'),
   (2, 'Created', 'Deferred', '{"id": "2", "object": "object2", "owner": "owner2"}', false, false, '2021-11-11 23:00:00.000'),
   (3, 'Created', 'Deferred', '{"id": "3", "object": "object3", "owner": "owner2"}', false, false, '2021-11-11 23:00:00.000'),
   (4, 'Created', 'Deferred', '{"id": "4", "object": "object4", "owner": "owner3"}', false, false, '2021-11-11 23:00:00.000'),
   (5, 'Created', 'Deferred', '{"id": "5", "object": "object5", "owner": "owner4"}', false, false, '2021-11-11 23:00:00.000'),
   (6, 'Created', 'Deferred', '{"id": "6", "object": "object6", "owner": "owner6"}', true, false, '2021-11-11 23:00:00.000'),
   (7, 'Created', 'Deferred', '{"id": "7", "object": "object7", "owner": "owner7"}', false, false, '2021-11-11 23:00:00.000');

-- +goose Down
TRUNCATE apartments;
TRUNCATE apartments_events;
