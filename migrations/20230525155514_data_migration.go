package migrations

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type ChannelSlackSettings struct {
	ID              int       `json:"id"`
	ChannelID       int       `gorm:"not null;" json:"channel_id"`
	IntegrationName string    `gorm:"size:100;not null;unique" json:"integration_name"`
	WebhookUrl      string    `gorm:"size:2000;not null;" json:"webhook_url"`
	SenderName      string    `gorm:"size:100;not null;" json:"sender_name"`
	ChannelName     string    `gorm:"size:100;not null;" json:"channel_name"`
	TenantID        string    `gorm:"size:100;not null;" json:"tenant_id"`
	Status          string    `gorm:"size:20;not null;" json:"status"`
	CreatedAt       time.Time `gorm:"size:100;not null;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"size:100;not null;" json:"updated_at"`
	Title           string    `gorm:"size:100;" json:"title"`
}

// Model for channels table.
type Channels struct {
	ChannelID   int    `gorm:"size:255;not null;" json:"channel_id"`
	ChannelName string `gorm:"size:255;not null;" json:"channel_name"`
	Status      string `gorm:"size:255;not null;" json:"status"`
	TenantID    string `gorm:"size:255;not null;" json:"tenant_id"`
	CreatedBy   string `gorm:"size:255;not null;" json:"created_by"`
}

// Model for channel_cloudwatch_settings table.
type ChannelCloudwatchSettings struct {
	ID              int       `json:"id"`
	ChannelID       int       `gorm:"not null;" json:"channel_id"`
	TenantID        string    `gorm:"size:50;not null;" json:"tenant_id"`
	IntegrationName string    `gorm:"size:50;not null;unique" json:"integration_name"`
	AccessKey       string    `gorm:"size:255;not null;" json:"access_key"`
	SecretKey       string    `gorm:"size:255;not null;" json:"secret_key"`
	Region          string    `gorm:"size:50;not null;" json:"region"`
	LogGroupName    string    `gorm:"size:50;not null;" json:"log_group_name"`
	Status          string    `gorm:"size:50;not null;" json:"status"`
	CreatedAt       time.Time `gorm:"size:50;not null;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"size:50;not null;" json:"updated_at"`
}

// Model for channel_jira_settings table.
type ChannelJiraSettings struct {
	ID              int       `json:"id"`
	ChannelID       int       `gorm:"not null;" json:"channel_id"`
	IntegrationName string    `gorm:"size:50;not null;" json:"integration_name"`
	IssueSummary    string    `gorm:"size:255;not null;" json:"issue_summary"`
	Site            string    `gorm:"size:255;not null;" json:"site"`
	Project         string    `gorm:"size:50;not null;" json:"project"`
	IssueType       string    `gorm:"size:50;not null;" json:"issue_type"`
	TenantID        string    `gorm:"size:50;not null;" json:"tenant_id"`
	UserEmail       string    `gorm:"size:255;not null;" json:"user_email"`
	Token           string    `gorm:"size:500;not null;" json:"token"`
	UserID          string    `gorm:"size:100;not null;" json:"user_id"`
	Status          string    `gorm:"size:50;not null;" json:"status"`
	CreatedAt       time.Time `gorm:"size:50;not null;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"size:50;not null;" json:"updated_at"`
}

// Model for channel_rsyslog_settings table.
type ChannelRsyslogSettings struct {
	ID              int       `json:"id"`
	ChannelID       int       `gorm:"not null;" json:"channel_id"`
	TenantID        string    `gorm:"size:50;not null;" json:"tenant_id"`
	IntegrationName string    `gorm:"size:50;not null;" json:"integration_name"`
	ServerAddress   string    `gorm:"size:100;not null;" json:"server_address"`
	Port            int       `gorm:"size:50;not null;" json:"port"`
	Transport       string    `gorm:"size:50;not null;" json:"transport"`
	Status          string    `gorm:"size:50;not null;" json:"status"`
	CreatedAt       time.Time `gorm:"size:50;not null;" json:"created_at"`
	UpdatedAt       time.Time `gorm:"size:50;not null;" json:"updated_at"`
}

// Model for channel_splunk_settings table.
type ChannelSplunkSettings struct {
	ID              int       `json:"id"`
	IntegrationName string    `gorm:"size:50;not null;" json:"integration_name,omitempty"`
	ChannelID       int       `gorm:"not null;" json:"channel_id,omitempty"`
	Url             string    `gorm:"size:255;not null;" json:"url,omitempty"`
	Token           string    `gorm:"size:255;not null;" json:"token,omitempty"`
	Source          string    `gorm:"size:50;not null;" json:"source,omitempty"`
	SourceType      string    `gorm:"size:50;not null;" json:"source_type,omitempty"`
	SplunkIndex     string    `gorm:"size:50;not null;" json:"splunk_index,omitempty"`
	TenantID        string    `gorm:"size:50;not null;" json:"tenant_id,omitempty"`
	Status          string    `gorm:"size:50;not null;" json:"status,omitempty"`
	CreatedAt       time.Time `gorm:"size:50;not null;" json:"created_at,omitempty"`
	UpdatedAt       time.Time `gorm:"size:50;not null;" json:"updated_at,omitempty"`
}

var id int

func init() {
	goose.AddMigration(upDataMigration, downDataMigration)
}

func upDataMigration(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	id = 1
	stmt1, err := tx.Prepare("INSERT INTO cwpp.channels (settings_id, channel_type, integration_name, status, tenant_id, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt1.Close()

	//slack
	cache := []ChannelSlackSettings{}
	rows, err := tx.Query("SELECT * FROM cwpp_0.channel_slack_settings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for rows.Next() {
		var settings ChannelSlackSettings
		created_at := ""
		updated_at := ""
		err := rows.Scan(&settings.ID, &settings.ChannelID, &settings.IntegrationName, &settings.WebhookUrl, &settings.SenderName, &settings.ChannelName, &settings.Title, &settings.TenantID, &settings.Status, &created_at, &updated_at)
		if err != nil {
			fmt.Println("5----")
			fmt.Println(err)
			return err
		}
		const layout = "2006-01-02 15:04:05"
		settings.UpdatedAt, err = time.Parse(layout, updated_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}
		settings.CreatedAt, err = time.Parse(layout, created_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}

		cache = append(cache, settings)
	}

	for _, settings := range cache {
		var channel Channels

		err = tx.QueryRow("SELECT * FROM cwpp_0.channels WHERE channel_id=$1 AND tenant_id=$2", settings.ChannelID, settings.TenantID).Scan(&channel.ChannelID, &channel.ChannelName, &channel.Status, &channel.TenantID, &channel.CreatedBy)
		if err != nil {
			fmt.Println("4----")
			return err
		}
		_, err = stmt1.Exec(settings.ChannelID, channel.ChannelName, settings.IntegrationName, settings.Status, settings.TenantID, channel.CreatedBy, settings.CreatedAt, settings.UpdatedAt)
		if err != nil {
			fmt.Println("1---")
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(`INSERT INTO cwpp.channel_slack_settings (channels_id, webhook_url, sender_name, channel_name, title) VALUES ($1, $2, $3, $4, $5)`, id, settings.WebhookUrl, settings.SenderName, settings.ChannelName, settings.Title)
		if err != nil {
			fmt.Println("2---")
			fmt.Println(err)
			return err
		}
		id++
	}

	// splunk
	splunkCache := []ChannelSplunkSettings{}
	rows, err = tx.Query("SELECT * FROM cwpp_0.channel_splunk_settings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for rows.Next() {
		var settings ChannelSplunkSettings
		created_at := ""
		updated_at := ""
		err := rows.Scan(&settings.ID, &settings.ChannelID, &settings.IntegrationName, &settings.Url, &settings.Token, &settings.Source, &settings.SplunkIndex, &settings.SourceType, &settings.TenantID, &settings.Status, &created_at, &updated_at)
		if err != nil {
			fmt.Println("5----")
			fmt.Println(err)
			return err
		}
		const layout = "2006-01-02 15:04:05"
		settings.UpdatedAt, err = time.Parse(layout, updated_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}
		settings.CreatedAt, err = time.Parse(layout, created_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}

		splunkCache = append(splunkCache, settings)
	}

	for _, settings := range splunkCache {
		var channel Channels
		err = tx.QueryRow("SELECT * FROM cwpp_0.channels WHERE channel_id=$1 AND tenant_id=$2", settings.ChannelID, settings.TenantID).Scan(&channel.ChannelID, &channel.ChannelName, &channel.Status, &channel.TenantID, &channel.CreatedBy)
		if err != nil {
			fmt.Println("4----")
			return err
		}
		_, err = stmt1.Exec(settings.ChannelID, channel.ChannelName, settings.IntegrationName, settings.Status, settings.TenantID, channel.CreatedBy, settings.CreatedAt, settings.UpdatedAt)
		if err != nil {
			fmt.Println("1---")
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(`INSERT INTO cwpp.channel_splunk_settings (channels_id, url, token, source, source_type, splunk_index) VALUES ($1, $2, $3, $4, $5, $6)`, id, settings.Url, settings.Token, settings.Source, settings.SourceType, settings.SplunkIndex)
		if err != nil {
			fmt.Println("2---")
			fmt.Println(err)
			return err
		}
		id++
	}

	// cloudwatch
	cloudwatchCache := []ChannelCloudwatchSettings{}
	rows, err = tx.Query("SELECT * FROM cwpp_0.channel_cloudwatch_settings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for rows.Next() {
		var settings ChannelCloudwatchSettings
		created_at := ""
		updated_at := ""
		err := rows.Scan(&settings.ID, &settings.ChannelID, &settings.IntegrationName, &settings.AccessKey, &settings.SecretKey, &settings.LogGroupName, &settings.Region, &settings.TenantID, &settings.Status, &created_at, &updated_at)
		if err != nil {
			fmt.Println("5----")
			fmt.Println(err)
			return err
		}
		const layout = "2006-01-02 15:04:05"
		settings.UpdatedAt, err = time.Parse(layout, updated_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}
		settings.CreatedAt, err = time.Parse(layout, created_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}

		cloudwatchCache = append(cloudwatchCache, settings)
	}

	for _, settings := range cloudwatchCache {
		var channel Channels
		err = tx.QueryRow("SELECT * FROM cwpp_0.channels WHERE channel_id=$1 AND tenant_id=$2", settings.ChannelID, settings.TenantID).Scan(&channel.ChannelID, &channel.ChannelName, &channel.Status, &channel.TenantID, &channel.CreatedBy)
		if err != nil {
			fmt.Println("4----")
			return err
		}
		_, err = stmt1.Exec(settings.ChannelID, channel.ChannelName, settings.IntegrationName, settings.Status, settings.TenantID, channel.CreatedBy, settings.CreatedAt, settings.UpdatedAt)
		if err != nil {
			fmt.Println("1---")
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(`INSERT INTO cwpp.channel_cloudwatch_settings (channels_id, access_key, secret_key, region, log_group_name) VALUES ($1, $2, $3, $4, $5)`, id, settings.AccessKey, settings.SecretKey, settings.Region, settings.LogGroupName)
		if err != nil {
			fmt.Println("2---")
			fmt.Println(err)
			return err
		}
		id++
	}

	// jira
	jiraCache := []ChannelJiraSettings{}
	rows, err = tx.Query("SELECT * FROM cwpp_0.channel_jira_settings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for rows.Next() {
		var settings ChannelJiraSettings
		created_at := ""
		updated_at := ""
		tags := ""

		err := rows.Scan(&settings.ID, &settings.ChannelID, &settings.IntegrationName, &settings.UserEmail, &settings.Token, &settings.UserID, &settings.IssueSummary, &settings.Site, &settings.Project, &settings.IssueType, &tags, &settings.TenantID, &settings.Status, &created_at, &updated_at)
		if err != nil {
			fmt.Println("5----")
			fmt.Println(err)
			return err
		}
		const layout = "2006-01-02 15:04:05"
		settings.UpdatedAt, err = time.Parse(layout, updated_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}
		settings.CreatedAt, err = time.Parse(layout, created_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}

		jiraCache = append(jiraCache, settings)
	}

	for _, settings := range jiraCache {
		var channel Channels
		err = tx.QueryRow("SELECT * FROM cwpp_0.channels WHERE channel_id=$1 AND tenant_id=$2", settings.ChannelID, settings.TenantID).Scan(&channel.ChannelID, &channel.ChannelName, &channel.Status, &channel.TenantID, &channel.CreatedBy)
		if err != nil {
			fmt.Println("4----")
			return err
		}
		_, err = stmt1.Exec(settings.ChannelID, channel.ChannelName, settings.IntegrationName, settings.Status, settings.TenantID, channel.CreatedBy, settings.CreatedAt, settings.UpdatedAt)
		if err != nil {
			fmt.Println("1---")
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(`INSERT INTO cwpp.channel_jira_settings (channels_id, issue_summary, site, project, issue_type, user_email, token, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, id, settings.IssueSummary, settings.Site, settings.Project, settings.IssueType, settings.UserEmail, &settings.Token, &settings.UserID)
		if err != nil {
			fmt.Println("2---")
			fmt.Println(err)
			return err
		}
		id++
	}

	// rsyslog
	rsyslogCache := []ChannelRsyslogSettings{}
	rows, err = tx.Query("SELECT * FROM cwpp_0.channel_rsyslog_settings")
	if err != nil {
		fmt.Println(err)
		return err
	}

	for rows.Next() {
		var settings ChannelRsyslogSettings
		created_at := ""
		updated_at := ""
		path := ""
		err := rows.Scan(&settings.ID, &settings.ChannelID, &settings.IntegrationName, &settings.ServerAddress, &path, &settings.Port, &settings.TenantID, &settings.Transport, &settings.Status, &created_at, &updated_at)
		if err != nil {
			fmt.Println("5----")
			fmt.Println(err)
			return err
		}
		const layout = "2006-01-02 15:04:05"
		settings.UpdatedAt, err = time.Parse(layout, updated_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}
		settings.CreatedAt, err = time.Parse(layout, created_at)
		if err != nil {
			fmt.Println("Error parsing time:", err)
			return err
		}

		rsyslogCache = append(rsyslogCache, settings)
	}

	for _, settings := range rsyslogCache {
		var channel Channels
		err = tx.QueryRow("SELECT * FROM cwpp_0.channels WHERE channel_id=$1 AND tenant_id=$2", settings.ChannelID, settings.TenantID).Scan(&channel.ChannelID, &channel.ChannelName, &channel.Status, &channel.TenantID, &channel.CreatedBy)
		if err != nil {
			fmt.Println("4----")
			return err
		}
		_, err = stmt1.Exec(settings.ChannelID, channel.ChannelName, settings.IntegrationName, settings.Status, settings.TenantID, channel.CreatedBy, settings.CreatedAt, settings.UpdatedAt)
		if err != nil {
			fmt.Println("1---")
			fmt.Println(err)
			return err
		}
		_, err = tx.Exec(`INSERT INTO cwpp.channel_rsyslog_settings (channels_id, server_address, port, transport) VALUES ($1, $2, $3, $4)`, id, settings.ServerAddress, settings.Port, settings.Transport)
		if err != nil {
			fmt.Println("2---")
			fmt.Println(err)
			return err
		}
		id++
	}

	return nil
}

func downDataMigration(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	id = 1
	return nil
}