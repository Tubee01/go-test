CREATE TABLE IF NOT EXISTS "user" (
	"id" serial PRIMARY KEY NOT NULL,
	"first_name" varchar(256) NOT NULL,
	"last_name" varchar(256) NOT NULL,
	"email" varchar(256) NOT NULL,
	"password" varchar(256) NOT NULL,
	"country_num_code" varchar(4) DEFAULT NULL,
	"created_at" timestamp DEFAULT now(),
	"updated_at" timestamp,
	"deleted_at" timestamp DEFAULT NULL,
	CONSTRAINT "uq_user_email" UNIQUE("email")
);