
CREATE TABLE "public"."teams"(
  "id" serial NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "public"."team_members"(
  "id" serial NOT NULL,
  "team_id" INT NOT NULL,
  "user_id" text NOT NULL,
  "name" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE TABLE "public"."team_invitations"(
  "id" serial NOT NULL,
  "team_id" INT NOT NULL,
  "host_user_id" text NOT NULL,
  "invitee_email" text NOT NULL,
  "invitee_name" text NOT NULL,
  "expires_at" timestamptz NOT NULL DEFAULT NOW() + INTERVAL '1 DAY',
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  UNIQUE ("team_id", "invitee_email")
);

CREATE TABLE "public"."users"(
  "id" text NOT NULL,
  "auth_id" text NOT NULL,
  "email" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "last_seen" timestamptz,
  PRIMARY KEY ("id")
);

CREATE TABLE "public"."todos"(
  "id" serial NOT NULL,
  "team_id" INT NOT NULL,
  "user_id" text NOT NULL,
  "title" text NOT NULL,
  "is_completed" boolean NOT NULL DEFAULT false,
  "is_public" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id")
);

CREATE OR REPLACE VIEW "public"."online_users" AS
 SELECT users.id,
    users.last_seen
   FROM users
  WHERE (users.last_seen >= (now() - '00:00:30'::interval));


