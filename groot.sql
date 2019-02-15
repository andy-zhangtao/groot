-- ----------------------------
-- Table structure for Groot
-- ----------------------------
DROP TABLE IF EXISTS "public"."groot_account";
CREATE TABLE "public"."groot_account"
(
  "time"    varchar(8) COLLATE "pg_catalog"."default",
  "b_icbc"  float DEFAULT 0,
  "b_abc"   float DEFAULT 0,
  "b_bocom" float DEFAULT 0,
  "b_cmb"   float DEFAULT 0,
  "b_citic" float DEFAULT 0,
  "b_ccb"   float DEFAULT 0,
  "b_bj"    float DEFAULT 0,
  "b_ali"   float DEFAULT 0,
  "b_oth"   float DEFAULT 0
);

ALTER TABLE "public"."groot_account"
  OWNER TO "postgres";