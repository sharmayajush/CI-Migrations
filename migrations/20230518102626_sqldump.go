package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upSqldump, downSqldump)
}

func upSqldump(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	if _, err := tx.Exec(`
--
-- PostgreSQL database dump
--
CREATE SCHEMA cwpp;
	ALTER SCHEMA cwpp OWNER TO yajush;

--
-- Name: channel_cloudwatch_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_cloudwatch_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_cloudwatch_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_cloudwatch_settings (
    id integer DEFAULT nextval('cwpp.channel_cloudwatch_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    integration_name character varying(50) NOT NULL,
    access_key character varying(255) NOT NULL,
    secret_key character varying(255) NOT NULL,
    log_group_name character varying(100) DEFAULT NULL::character varying,
    region character varying(25) NOT NULL,
    tenant_id character varying(25) NOT NULL,
    status character varying(20) DEFAULT NULL::character varying,
    created_at text,
    updated_at text
);



--
-- Name: channel_elasticsearch_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_elasticsearch_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_elasticsearch_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_elasticsearch_settings (
    id integer DEFAULT nextval('cwpp.channel_elasticsearch_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    integration_name character varying(50) NOT NULL,
    api_url character varying(255) NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255) NOT NULL,
    mount_path character varying(255) DEFAULT NULL::character varying,
    index_name character varying(100) DEFAULT NULL::character varying,
    index_type character varying(50) DEFAULT NULL::character varying,
    tenant_id character varying(25) NOT NULL,
    status character varying(20) DEFAULT NULL::character varying,
    created_at text,
    updated_at text
);



--
-- Name: channel_jira_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_jira_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_jira_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_jira_settings (
    id integer DEFAULT nextval('cwpp.channel_jira_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    integration_name character varying(50) NOT NULL,
    user_email character varying(255) DEFAULT NULL::character varying,
    token character varying(255) DEFAULT NULL::character varying,
    user_id character varying(100) DEFAULT NULL::character varying,
    issue_summary character varying(50) DEFAULT NULL::character varying,
    site character varying(255) NOT NULL,
    project character varying(50) DEFAULT NULL::character varying,
    issue_type character varying(50) DEFAULT NULL::character varying,
    tags_to_be_sent character varying(50) DEFAULT NULL::character varying,
    tenant_id character varying(25) NOT NULL,
    status character varying(20) DEFAULT NULL::character varying,
    created_at text,
    updated_at text
);



--
-- Name: channel_rsyslog_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_rsyslog_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_rsyslog_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_rsyslog_settings (
    id integer DEFAULT nextval('cwpp.channel_rsyslog_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    integration_name character varying(50) NOT NULL,
    server_address character varying(500) NOT NULL,
    path character varying(100) DEFAULT NULL::character varying,
    port integer NOT NULL,
    tenant_id character varying(25) NOT NULL,
    transport character varying(50) NOT NULL,
    status character varying(20) DEFAULT NULL::character varying,
    created_at text,
    updated_at text
);



--
-- Name: channel_slack_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_slack_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_slack_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_slack_settings (
    id integer DEFAULT nextval('cwpp.channel_slack_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    integration_name character varying(50) NOT NULL,
    sender_name character varying(50) NOT NULL,
    webhook_url character varying(255) NOT NULL,
    channel_name character varying(50) NOT NULL,
    title character varying(200) DEFAULT NULL::character varying,
    tenant_id character varying(25) NOT NULL,
    status character varying(20) DEFAULT NULL::character varying,
    created_at text,
    updated_at text
);



--
-- Name: channel_splunk_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_splunk_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_splunk_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_splunk_settings (
    id integer DEFAULT nextval('cwpp.channel_splunk_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    integration_name character varying(50) NOT NULL,
    url character varying(255) NOT NULL,
    token character varying(200) NOT NULL,
    source character varying(50) NOT NULL,
    splunk_index character varying(50) NOT NULL,
    source_type character varying(50) NOT NULL,
    tenant_id character varying(25) NOT NULL,
    status character varying(20) DEFAULT NULL::character varying,
    created_at text,
    updated_at text
);



--
-- Name: channel_splunkapp_settings_seq; Type: SEQUENCE; Schema: cwpp; Owner: postgres
--

CREATE SEQUENCE cwpp.channel_splunkapp_settings_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;



--
-- Name: channel_splunkapp_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channel_splunkapp_settings (
    id integer DEFAULT nextval('cwpp.channel_splunkapp_settings_seq'::regclass) NOT NULL,
    channel_id integer NOT NULL,
    tenant_id character varying(25) NOT NULL,
    splunk_app_name character varying(50) NOT NULL,
    integration_name character varying(100) NOT NULL,
    splunk_http_url character varying(255) NOT NULL,
    splunk_app_token character varying(255) NOT NULL,
    splunk_app_source character varying(50) NOT NULL,
    splunk_app_index character varying(50) NOT NULL,
    splunk_app_sourcetype character varying(50) NOT NULL
);



--
-- Name: channels; Type: TABLE; Schema: cwpp; Owner: postgres
--

CREATE TABLE cwpp.channels (
    channel_id integer NOT NULL,
    channel_name character varying(25) NOT NULL,
    status character varying(50) DEFAULT NULL::character varying,
    tenant_id character varying(25) NOT NULL,
    created_by character varying(25) DEFAULT NULL::character varying
);




ALTER TABLE ONLY cwpp.channels
    ADD CONSTRAINT ch_name UNIQUE (channel_id, tenant_id);


--
-- Name: channel_cloudwatch_settings channel_cloudwatch_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_cloudwatch_settings
    ADD CONSTRAINT channel_cloudwatch_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_elasticsearch_settings channel_elasticsearch_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_elasticsearch_settings
    ADD CONSTRAINT channel_elasticsearch_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_jira_settings channel_jira_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_jira_settings
    ADD CONSTRAINT channel_jira_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_rsyslog_settings channel_rsyslog_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_rsyslog_settings
    ADD CONSTRAINT channel_rsyslog_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_slack_settings channel_slack_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_slack_settings
    ADD CONSTRAINT channel_slack_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_splunk_settings channel_splunk_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_splunk_settings
    ADD CONSTRAINT channel_splunk_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_splunkapp_settings channel_splunkapp_settings_pkey; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_splunkapp_settings
    ADD CONSTRAINT channel_splunkapp_settings_pkey PRIMARY KEY (id);


--
-- Name: channel_splunkapp_settings integration_name; Type: CONSTRAINT; Schema: cwpp; Owner: postgres
--

ALTER TABLE ONLY cwpp.channel_splunkapp_settings
    ADD CONSTRAINT integration_name UNIQUE (integration_name);

--
-- PostgreSQL database dump complete
--

--
-- inserting data
--
INSERT INTO cwpp.channels (channel_id, channel_name, status, tenant_id, created_by)
VALUES 
  (1, 'Slack', 'Active', 'Tenant_1', 'User_1'),
  (1, 'Slack', 'Active', 'Tenant_2', 'User_2'),
  (1, 'Slack', 'Active', 'Tenant_3', 'User_3');


INSERT INTO cwpp.channel_slack_settings (channel_id, integration_name, sender_name, webhook_url, channel_name, title, tenant_id, status, created_at, updated_at)
VALUES 
  (1, 'Integration 1', 'Sender 1', 'https://webhook-url-1', 'Channel 1', 'Title_1', 'Tenant_1', 'Active', '2023-03-17 12:13:30.670761', '2023-03-17 12:13:30.670761'),
  (1, 'Integration 2', 'Sender 2', 'https://webhook-url-2', 'Channel 2', 'Title_2', 'Tenant_2', 'Inactive', '2023-03-17 12:13:30.670761', '2023-03-17 12:13:30.670761'),
  (1, 'Integration 3', 'Sender 3', 'https://webhook-url-3', 'Channel_3', 'Title_3', 'Tenant_3', 'Active', '2023-03-17 12:13:30.670761', '2023-03-17 12:13:30.670761');
  
  
  INSERT INTO cwpp.channels (channel_id, channel_name, status, tenant_id, created_by)
  VALUES 
    (2, 'Splunk', 'Active', 'Tenant_1', 'User_1'),
    (2, 'Splunk', 'Active', 'Tenant_2', 'User_2'),
    (2, 'Splunk', 'Active', 'Tenant_3', 'User_3');
  

  INSERT INTO cwpp.channel_splunk_settings (channel_id, integration_name, url, token, source, splunk_index, source_type, tenant_id, status, created_at, updated_at)
  VALUES 
    (2, 'Integration 1', 'https://example.com', 'splunk_token_1', 'source_1', 'index_1', 'source_type_1', 'Tenant_1', 'Active', '2023-05-23 10:00:00', '2023-05-23 10:00:00'),
    (2, 'Integration 2', 'https://example.com', 'splunk_token_2', 'source_2', 'index_2', 'source_type_2', 'Tenant_2', 'Active', '2023-05-23 11:00:00', '2023-05-23 11:00:00'),
    (2, 'Integration 3', 'https://example.com', 'splunk_token_3', 'source_3', 'index_3', 'source_type_3', 'Tenant_3', 'Active', '2023-05-23 12:00:00', '2023-05-23 12:00:00');
  
    INSERT INTO cwpp.channels (channel_id, channel_name, status, tenant_id, created_by)
    VALUES 
      (3, 'Cloudwatch', 'Active', 'Tenant_1', 'User_1'),
      (3, 'Cloudwatch', 'Active', 'Tenant_2', 'User_2'),
      (3, 'Cloudwatch', 'Active', 'Tenant_3', 'User_3');
    

      INSERT INTO cwpp.channel_cloudwatch_settings (channel_id, integration_name, access_key, secret_key, log_group_name, region, tenant_id, status, created_at, updated_at)
      VALUES 
        (3, 'Integration 1', 'access_key_1', 'secret_key_1', 'log_group_1', 'us-west-1', 'Tenant_1', 'Active', '2023-05-23 10:00:00', '2023-05-23 10:00:00'),
        (3, 'Integration 2', 'access_key_2', 'secret_key_2', 'log_group_2', 'us-east-1', 'Tenant_2', 'Active', '2023-05-23 11:00:00', '2023-05-23 11:00:00'),
        (3, 'Integration 3', 'access_key_3', 'secret_key_3', 'log_group_3', 'eu-central-1', 'Tenant_3', 'Active', '2023-05-23 12:00:00', '2023-05-23 12:00:00');
  
        INSERT INTO cwpp.channels (channel_id, channel_name, status, tenant_id, created_by)
        VALUES 
          (4, 'Jira', 'Active', 'Tenant_1', 'User_1'),
          (4, 'Jira', 'Active', 'Tenant_2', 'User_2'),
          (4, 'Jira', 'Active', 'Tenant_3', 'User_3');
  
        INSERT INTO cwpp.channel_jira_settings (channel_id, integration_name, user_email, token, user_id, issue_summary, site, project, issue_type, tags_to_be_sent, tenant_id, status, created_at, updated_at)
        VALUES 
          (4, 'Integration 1', 'user1@example.com', 'jira_token_1', 'user_id_1', 'Summary 1', 'https://example.jira.com', 'Project 1', 'Bug', 'tag1,tag2', 'Tenant_1', 'Active', '2023-05-23 10:00:00', '2023-05-23 10:00:00'),
          (4, 'Integration 2', 'user2@example.com', 'jira_token_2', 'user_id_2', 'Summary 2', 'https://example.jira.com', 'Project 2', 'Task', 'tag3,tag4', 'Tenant_2', 'Active', '2023-05-23 11:00:00', '2023-05-23 11:00:00'),
          (4, 'Integration 3', 'user3@example.com', 'jira_token_3', 'user_id_3', 'Summary 3', 'https://example.jira.com', 'Project 3', 'Story', 'tag5,tag6', 'Tenant_3', 'Active', '2023-05-23 12:00:00', '2023-05-23 12:00:00');
        
          INSERT INTO cwpp.channels (channel_id, channel_name, status, tenant_id, created_by)
          VALUES 
            (5, 'Rsyslog', 'Active', 'Tenant_1', 'User_1'),
            (5, 'Rsyslog', 'Active', 'Tenant_2', 'User_2'),
            (5, 'Rsyslog', 'Active', 'Tenant_3', 'User_3');
          INSERT INTO cwpp.channel_rsyslog_settings (channel_id, integration_name, server_address, path, port, tenant_id, transport, status, created_at, updated_at)
          VALUES 
            (5, 'Integration 1', 'server1.example.com', '/var/log/syslog', 514, 'Tenant_1', 'tcp', 'Active', '2023-05-23 10:00:00', '2023-05-23 10:00:00'),
            (5, 'Integration 2', 'server2.example.com', '/var/log/messages', 514, 'Tenant_2', 'udp', 'Active', '2023-05-23 11:00:00', '2023-05-23 11:00:00'),
            (5, 'Integration 3', 'server3.example.com', '/var/log/auth.log', 514, 'Tenant_3', 'tcp', 'Active', '2023-05-23 12:00:00', '2023-05-23 12:00:00');
	`); err != nil {
		return err
	}
	return nil
}

func downSqldump(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	if _, err := tx.Exec(`
	DROP SCHEMA cwpp CASCADE
	`); err != nil {
		return err
	}
	return nil
}
