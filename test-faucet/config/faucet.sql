/*
 Navicat Premium Data Transfer

 Source Server         : localhost-postgre
 Source Server Type    : PostgreSQL
 Source Server Version : 90602
 Source Host           : localhost:5432
 Source Catalog        : test-faucet
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 90602
 File Encoding         : 65001

 Date: 23/02/2018 14:54:13
*/


-- ----------------------------
-- Sequence structure for faucet_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."faucet_id_seq";
CREATE SEQUENCE "public"."faucet_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for faucet
-- ----------------------------
DROP TABLE IF EXISTS "public"."faucet";
CREATE TABLE "public"."faucet" (
  "id" int4 NOT NULL DEFAULT nextval('faucet_id_seq'::regclass),
  "address" varchar(64) COLLATE "pg_catalog"."default",
  "amount" int4,
  "time" timestamp(6),
  "coin" varchar(16) COLLATE "pg_catalog"."default"
)
;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."faucet_id_seq"
OWNED BY "public"."faucet"."id";
SELECT setval('"public"."faucet_id_seq"', 61, true);

-- ----------------------------
-- Primary Key structure for table faucet
-- ----------------------------
ALTER TABLE "public"."faucet" ADD CONSTRAINT "faucet_pkey" PRIMARY KEY ("id");
