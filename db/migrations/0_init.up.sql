 CREATE TYPE "type" AS ENUM (
  'expense',
  'holiday',
  'absence'
);

CREATE TABLE "user" (
     "id" INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
     "first_name" varchar,
     "last_name" varchar,
     "email" varchar NOT NULL,
     "password" varchar NOT NULL,
     "birth_date" timestamp,
     "created_at" timestamp NOT NULL DEFAULT now()
);

CREATE TABLE expense (
     id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
     type varchar,
     userId varchar,
     name varchar,
     amount decimal,
     status varchar,
     category varchar,
     "group" varchar,
     comment varchar,
     date date,
     creation date
);

CREATE TABLE holiday (
   id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
   type varchar,
   userId varchar,
   name varchar,
   status varchar,
   fromDate date,
   toDate date,
   creation date
);

CREATE TABLE absence (
   id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
   type varchar,
   userId varchar,
   name varchar,
   amount decimal,
   status varchar,
   fromDate date,
   toDate date,
   creation date
);
