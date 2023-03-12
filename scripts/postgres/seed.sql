
INSERT INTO users (id, name)
VALUES (1, 'Joao das Dores'),
       (2, 'Marciano');

INSERT INTO cost_centers (id, title)
VALUES (1, 'health'),
       (2, 'food');

INSERT INTO activities (id, cost_center_id, title)
VALUES (11, 1,'remedies'),
       (12, 1,'appointment'),
       (13, 2,'breakfast'),
       (14, 2,'lunch');

INSERT INTO expenses (id, activity_id, user_id, cost, description, billing_date)
VALUES (1112, 11, 1, 50.95, 'drugstore','2022-10-11'),
       (1113, 11, 1, 50, 'drugstore','2022-10-11'),
       (1114, 12, 1, 100, 'lungs exam','2022-10-13'),
       (1115, 13, 1, 40.5, 'Eggs & Cheese','2022-10-10'),
       (1116, 14, 1, 59.5, 'Domino','2022-10-08'),
       (1117, 14, 1, 160.2, 'Corleones','2022-09-18');
