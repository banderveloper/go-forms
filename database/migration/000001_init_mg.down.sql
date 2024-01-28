ALTER TABLE "Forms" DROP CONSTRAINT IF EXISTS "Forms_fk0";

ALTER TABLE "FormsInputs" DROP CONSTRAINT IF EXISTS "FormsInputs_fk0";

ALTER TABLE "FormsInputs" DROP CONSTRAINT IF EXISTS "FormsInputs_fk1";

ALTER TABLE "FormInputsConstraints" DROP CONSTRAINT IF EXISTS "FormInputsConstraints_fk0";

ALTER TABLE "FormInputsConstraints" DROP CONSTRAINT IF EXISTS "FormInputsConstraints_fk1";

ALTER TABLE "FormsAnswers" DROP CONSTRAINT IF EXISTS "FormsAnswers_fk0";

ALTER TABLE "FormsAnswers" DROP CONSTRAINT IF EXISTS "FormsAnswers_fk1";

ALTER TABLE "FormsAnswers" DROP CONSTRAINT IF EXISTS "FormsAnswers_fk2";
