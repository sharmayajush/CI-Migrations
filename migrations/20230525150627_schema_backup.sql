-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

-- Create a new schema
-- CREATE SCHEMA cwpp_0;

-- Duplicate tables from the existing schema to the new schema
DO LANGUAGE plpgsql
$body$
DECLARE
   old_schema NAME = 'cwpp';
   new_schema NAME = 'cwpp_0';
   tbl TEXT;
   sql TEXT;
BEGIN
  EXECUTE format('CREATE SCHEMA IF NOT EXISTS %I', new_schema);

  FOR tbl IN
    SELECT table_name
    FROM information_schema.tables
    WHERE table_schema=old_schema
  LOOP
    sql := format(
            'CREATE TABLE IF NOT EXISTS %I.%I '
            '(LIKE %I.%I INCLUDING INDEXES INCLUDING CONSTRAINTS)'
            , new_schema, tbl, old_schema, tbl);

    EXECUTE sql;

    sql := format(
            'INSERT INTO %I.%I '
            'SELECT * FROM %I.%I'
            , new_schema, tbl, old_schema, tbl);

    EXECUTE sql;
  END LOOP;
END
$body$;

DROP TABLE cwpp.channel_slack_settings, cwpp.channel_splunk_settings, cwpp.channel_cloudwatch_settings, cwpp.channel_jira_settings, cwpp.channel_rsyslog_settings;
DROP TABLE cwpp.channels;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP SCHEMA cwpp CASCADE;
ALTER SCHEMA cwpp_0 RENAME TO cwpp;
-- +goose StatementEnd
