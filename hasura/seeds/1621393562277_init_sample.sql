INSERT INTO "public"."users" (id, name, email) VALUES
('foo', 'Ichiro', 'ichiro@example.com'),
('bar', 'Jiro', 'jiro@example.com'),
('baz', 'Hanako', 'hanako@example.com');

INSERT INTO "public"."todos" (title, user_id) VALUES
('Buy bread', 'foo'),
('Buy milk', 'foo'),
('Cook curry', 'bar'),
('Dish washing', 'baz');

