CREATE TABLE "Users"
(
    "id"            serial       NOT NULL UNIQUE,
    "username"      VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    CONSTRAINT "Users_pk" PRIMARY KEY ("id")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Forms"
(
    "id"        serial       NOT NULL UNIQUE,
    "author_id" integer      NOT NULL,
    "title"     VARCHAR(255) NOT NULL,
    CONSTRAINT "Forms_pk" PRIMARY KEY ("id")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "FormsInputs"
(
    "id"      serial       NOT NULL,
    "form_id" integer      NOT NULL,
    "type_id" integer      NOT NULL,
    "name"    VARCHAR(255) NOT NULL,
    CONSTRAINT "FormsInputs_pk" PRIMARY KEY ("id")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Constraints"
(
    "id"   serial       NOT NULL,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    CONSTRAINT "Constraints_pk" PRIMARY KEY ("id")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "FormInputsConstraints"
(
    "id"            serial  NOT NULL,
    "form_input_id" integer NOT NULL,
    "constraint_id" integer NOT NULL,
    "value"         VARCHAR(json
),
	CONSTRAINT "FormInputsConstraints_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);



CREATE TABLE "FormsAnswers"
(
    "id"        serial       NOT NULL,
    "author_id" integer      NOT NULL,
    "form_id"   integer      NOT NULL,
    "input_id"  integer      NOT NULL,
    "value"     VARCHAR(255) NOT NULL,
    CONSTRAINT "FormsAnswers_pk" PRIMARY KEY ("id")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "FormInputTypes"
(
    "id"   serial       NOT NULL,
    "name" VARCHAR(255) NOT NULL UNIQUE,
    CONSTRAINT "FormInputTypes_pk" PRIMARY KEY ("id")
) WITH (
      OIDS = FALSE
      );


ALTER TABLE "Forms"
    ADD CONSTRAINT "Forms_fk0" FOREIGN KEY ("author_id") REFERENCES "Users" ("id");

ALTER TABLE "FormsInputs"
    ADD CONSTRAINT "FormsInputs_fk0" FOREIGN KEY ("form_id") REFERENCES "Forms" ("id");
ALTER TABLE "FormsInputs"
    ADD CONSTRAINT "FormsInputs_fk1" FOREIGN KEY ("type_id") REFERENCES "FormInputTypes" ("id");


ALTER TABLE "FormInputsConstraints"
    ADD CONSTRAINT "FormInputsConstraints_fk0" FOREIGN KEY ("form_input_id") REFERENCES "FormsInputs" ("id");
ALTER TABLE "FormInputsConstraints"
    ADD CONSTRAINT "FormInputsConstraints_fk1" FOREIGN KEY ("constraint_id") REFERENCES "Constraints" ("id");

ALTER TABLE "FormsAnswers"
    ADD CONSTRAINT "FormsAnswers_fk0" FOREIGN KEY ("author_id") REFERENCES "Users" ("id");
ALTER TABLE "FormsAnswers"
    ADD CONSTRAINT "FormsAnswers_fk1" FOREIGN KEY ("form_id") REFERENCES "Forms" ("id");
ALTER TABLE "FormsAnswers"
    ADD CONSTRAINT "FormsAnswers_fk2" FOREIGN KEY ("input_id") REFERENCES "FormsInputs" ("id");