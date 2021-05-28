INSERT INTO "public"."teams" (id, name) VALUES
(1, 'FooBar inc.'),
(2, 'Hoge Fuga org.'),
(3, 'ABC inc.')
ON CONFLICT DO NOTHING;

INSERT INTO "public"."team_invitations" (team_id, host_user_id, invitee_email, invitee_name) VALUES
(1, 'taro_user_id', 'jiro@example.com', 'Jiro'),
(2, 'taro_user_id', 'jiro@example.com', 'jiro'),
(3, 'taro_user_id', 'jiro@example.com', '__jiro__'),
(1, 'taro_user_id', 'hanako@example.com', 'Hanako')
ON CONFLICT DO NOTHING;

INSERT INTO "public"."team_members"(team_id, user_id, name) VALUES
(1, 'taro_user_id', 'Taro'),
(2, 'taro_user_id', 'taro-san')
ON CONFLICT DO NOTHING;

INSERT INTO "public"."users"(id, auth_id, email) VALUES
('uuid_user_id', 'firebase_taro_user_id', 'taro@example.com')
ON CONFLICT DO NOTHING;



