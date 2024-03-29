CREATE TABLE "violations"
(
    "violation_id"   UUID         NOT NULL,
    "violation_name" VARCHAR(255) NOT NULL,
    "fine_amount"    BIGINT       NOT NULL
);
ALTER TABLE
    "violations"
    ADD PRIMARY KEY ("violation_id");
CREATE TABLE "cases"
(
    "case_id"         UUID         NOT NULL,
    "transport_id"    UUID         NOT NULL,
    "camera_id"       UUID         NOT NULL,
    "violation_id"    UUID         NOT NULL,
    "violation_value" VARCHAR(255) NOT NULL,
    "required_skill"  INTEGER      NOT NULL,
    "case_date"       DATE         NOT NULL,
    "is_solved"       BOOLEAN      NOT NULL,
    "fine_decision"   BOOLEAN      NOT NULL,
    "solved_at"       TIMESTAMP
);
ALTER TABLE
    "cases"
    ADD PRIMARY KEY ("case_id");
CREATE TABLE "camera_types"
(
    "camera_type_id"   UUID         NOT NULL,
    "camera_type_name" VARCHAR(255) NOT NULL
);
ALTER TABLE
    "camera_types"
    ADD PRIMARY KEY ("camera_type_id");
CREATE TABLE "transports"
(
    "transport_id"    UUID        NOT NULL,
    "transport_chars" VARCHAR(10) NOT NULL,
    "transport_nums"  VARCHAR(10) NOT NULL,
    "region"          VARCHAR(10) NOT NULL,
    "person_id"       UUID        NOT NULL
);
ALTER TABLE
    "transports"
    ADD PRIMARY KEY ("transport_id");
CREATE TABLE "cameras"
(
    "camera_id"        UUID             NOT NULL,
    "camera_type_id"   UUID             NOT NULL,
    "camera_latitude"  DOUBLE PRECISION NOT NULL,
    "camera_longitude" DOUBLE PRECISION NOT NULL,
    "short_desc"       VARCHAR(255)     NOT NULL
);
ALTER TABLE
    "cameras"
    ADD PRIMARY KEY ("camera_id");
CREATE TABLE "persons"
(
    "id"        UUID         NOT NULL,
    "phone_num" VARCHAR(255) NOT NULL,
    "email"     VARCHAR(255) NULL,
    "vk_id"     VARCHAR(255) NULL,
    "tg_id"     VARCHAR(255) NULL
);
ALTER TABLE
    "persons"
    ADD PRIMARY KEY ("id");
ALTER TABLE
    "cameras"
    ADD CONSTRAINT "cameras_camera_type_id_foreign" FOREIGN KEY ("camera_type_id") REFERENCES "camera_types" ("camera_type_id");
ALTER TABLE
    "cases"
    ADD CONSTRAINT "cases_camera_id_foreign" FOREIGN KEY ("camera_id") REFERENCES "cameras" ("camera_id");
ALTER TABLE
    "cases"
    ADD CONSTRAINT "cases_violation_id_foreign" FOREIGN KEY ("violation_id") REFERENCES "violations" ("violation_id");
ALTER TABLE
    "cases"
    ADD CONSTRAINT "cases_transport_id_foreign" FOREIGN KEY ("transport_id") REFERENCES "transports" ("transport_id");
ALTER TABLE
    "transports"
    ADD CONSTRAINT "transports_person_id_foreign" FOREIGN KEY ("person_id") REFERENCES "persons" ("id");