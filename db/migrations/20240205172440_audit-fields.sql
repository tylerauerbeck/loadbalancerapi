-- +goose Up
-- modify "load_balancers" table
ALTER TABLE "load_balancers" ADD COLUMN "created_by" character varying NULL, ADD COLUMN "updated_by" character varying NULL;
-- create index "loadbalancer_created_by" to table: "load_balancers"
CREATE INDEX "loadbalancer_created_by" ON "load_balancers" ("created_by");
-- create index "loadbalancer_updated_by" to table: "load_balancers"
CREATE INDEX "loadbalancer_updated_by" ON "load_balancers" ("updated_by");
-- modify "providers" table
ALTER TABLE "providers" ADD COLUMN "created_by" character varying NULL, ADD COLUMN "updated_by" character varying NULL;
-- create index "provider_created_by" to table: "providers"
CREATE INDEX "provider_created_by" ON "providers" ("created_by");
-- create index "provider_updated_by" to table: "providers"
CREATE INDEX "provider_updated_by" ON "providers" ("updated_by");

-- +goose Down
-- reverse: create index "provider_updated_by" to table: "providers"
DROP INDEX "provider_updated_by";
-- reverse: create index "provider_created_by" to table: "providers"
DROP INDEX "provider_created_by";
-- reverse: modify "providers" table
ALTER TABLE "providers" DROP COLUMN "updated_by", DROP COLUMN "created_by";
-- reverse: create index "loadbalancer_updated_by" to table: "load_balancers"
DROP INDEX "loadbalancer_updated_by";
-- reverse: create index "loadbalancer_created_by" to table: "load_balancers"
DROP INDEX "loadbalancer_created_by";
-- reverse: modify "load_balancers" table
ALTER TABLE "load_balancers" DROP COLUMN "updated_by", DROP COLUMN "created_by";
