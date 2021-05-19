
CREATE TABLE "public"."users"(
  "id" text NOT NULL,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "last_seen" timestamptz,
  PRIMARY KEY ("id")
);

CREATE TABLE "public"."todos"(
  "id" serial NOT NULL,
  "title" text NOT NULL,
  "is_completed" boolean NOT NULL DEFAULT false,
  "is_public" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "user_id" text NOT NULL, PRIMARY KEY ("id")
);

CREATE OR REPLACE VIEW "public"."online_users" AS
 SELECT users.id,
    users.last_seen
   FROM users
  WHERE (users.last_seen >= (now() - '00:00:30'::interval));


