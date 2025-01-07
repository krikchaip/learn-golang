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
