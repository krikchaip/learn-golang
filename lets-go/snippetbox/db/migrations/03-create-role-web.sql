CREATE ROLE web WITH LOGIN PASSWORD 'secret';

GRANT INSERT, SELECT, UPDATE, DELETE
ON ALL TABLES
IN SCHEMA "public"
TO web;

-- the user also requires permissions for SEQUENCE data type
-- ref: https://stackoverflow.com/questions/9325017/error-permission-denied-for-sequence-cities-id-seq-using-postgres
GRANT SELECT, USAGE
ON ALL SEQUENCES
IN SCHEMA "public"
TO web;

-- granting permissions for future created postgres objects (tables, sequences, ...)
-- refs: https://stackoverflow.com/questions/76952543/postgres-granting-all-permissions-but-still-denied-on-tables
--       https://www.postgresql.org/docs/current/sql-alterdefaultprivileges.html
ALTER DEFAULT PRIVILEGES
IN SCHEMA "public"
GRANT INSERT, SELECT, UPDATE, DELETE
ON TABLES
TO web;

ALTER DEFAULT PRIVILEGES
IN SCHEMA "public"
GRANT SELECT, USAGE
ON SEQUENCES
TO web;
