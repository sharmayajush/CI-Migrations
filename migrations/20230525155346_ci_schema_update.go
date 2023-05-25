package migrations

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCiSchemaUpdate, downCiSchemaUpdate)
}

func upCiSchemaUpdate(tx *sql.Tx) error {
	// This code is executed when the migration is applied.

	

	if _, err := tx.Exec(`
--
-- Name: channels; Type: TABLE; Schema: cwpp; Owner: postgres
--

	CREATE TABLE cwpp.channels (
		id serial NOT NULL,
		settings_id int NOT NULL,
		channel_type varchar NOT NULL,
		integration_name varchar(100) NOT NULL,
		status varchar(50) DEFAULT NULL,
	  	tenant_id varchar(255) DEFAULT NULL,
	  	created_by varchar(255) NOT NULL,
		created_at timestamp DEFAULT NULL,
		updated_at timestamp DEFAULT NULL,
		PRIMARY KEY (id)
  	) ;

	`); err != nil {
		return err
	}

	if _, err := tx.Exec(`
--
-- Name: channel_slack_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--
	CREATE TABLE cwpp.channel_slack_settings (
		channels_id int NOT NULL,
		webhook_url varchar(255) NOT NULL,
	  	sender_name varchar(50) NOT NULL,
	  	channel_name varchar(50) NOT NULL,
	  	title varchar(200),
		PRIMARY KEY (channels_id),
	 	FOREIGN KEY (channels_id) REFERENCES cwpp.channels (id)
  	) ;
	

--
-- Name: channel_splunk_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--
	  
	CREATE TABLE cwpp.channel_splunk_settings (
	channels_id int NOT NULL,
	url varchar(2000) NOT NULL,
	token varchar(1000) NOT NULL,
	source varchar(100) NOT NULL,
	source_type varchar(100),
	splunk_index varchar(100),
	PRIMARY KEY (channels_id), FOREIGN KEY (channels_id) REFERENCES cwpp.channels (id)
	) ;
	  
	
--
-- Name: channel_cloudwatch_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--
	CREATE TABLE cwpp.channel_cloudwatch_settings (
		channels_id int NOT NULL,
		access_key varchar(2000) NOT NULL,
	  	secret_key varchar(1000) NOT NULL,
	  	region varchar(100) NOT NULL,
	  	log_group_name varchar(100),
		PRIMARY KEY (channels_id),
	  	FOREIGN KEY (channels_id) REFERENCES cwpp.channels (id)
  	) ;
	
		
--
-- Name: channel_jira_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--
	
	CREATE TABLE cwpp.channel_jira_settings (
		channels_id int NOT NULL,
		issue_summary varchar(2000) NOT NULL,
	  	site varchar(1000) NOT NULL,
	  	project varchar(100) NOT NULL,
	  	issue_type varchar(100),
	  	user_email varchar(100),
	  	token varchar(1000),
	  	user_id varchar(100),
		PRIMARY KEY (channels_id),
	  	FOREIGN KEY (channels_id) REFERENCES cwpp.channels (id)
  	) ;

--
-- Name: channel_rsyslog_settings; Type: TABLE; Schema: cwpp; Owner: postgres
--
	  
	CREATE TABLE cwpp.channel_rsyslog_settings (
		channels_id int NOT NULL,
		server_address varchar(2000) NOT NULL,
		port int NOT NULL,
		transport varchar(100) NOT NULL,
		PRIMARY KEY (channels_id),
		FOREIGN KEY (channels_id) REFERENCES cwpp.channels (id)
	) ;
	  


	`); err != nil {
		return err
	}

	return nil
}

func downCiSchemaUpdate(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.

	if _, err := tx.Exec(`
	DROP table cwpp.channel_rsyslog_settings;
	DROP table cwpp.channel_jira_settings;
	DROP table cwpp.channel_cloudwatch_settings;
	DROP table cwpp.channel_splunk_settings;
	DROP table cwpp.channel_slack_settings;
	DROP table cwpp.channels;

	
	`); err != nil {
		return err
	}
	return nil
}