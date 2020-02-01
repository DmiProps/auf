package database

import (
	"context"
	"time"

	"github.com/DmiProps/auf/settings"
	"github.com/jackc/pgx/v4"
)

func createActualSchema(conn *pgx.Conn) error {

	// Add info about actual version database schema
	dbSchemaDate, _ := time.Parse(settings.ShortDateForm, settings.DbSchemaDate)

	_, err := conn.Exec(
		context.Background(),
		`insert into version(id, schema_version, schema_date) values (0, $1, $2)
		on conflict (id) do update set schema_version = $1, schema_date = $2`,
		settings.DbSchemaVersion,
		dbSchemaDate)
	if err != nil {
		return err
	}

	// Create table 'accounts' if not exists
	_, err = conn.Exec(
		context.Background(),
		`create table if not exists accounts (
			id				serial primary key,		-- account id
			username		varchar(100) not null,	-- user name
			email			varchar(100) not null,	-- e-mail
			password_hash	varchar(80) not null,	-- password hash
			phone			varchar(30) not null default '', -- phone number
			email_confirmed	boolean not null default false, -- email confirmation flag
			phone_confirmed	boolean not null default false, -- phone confirmation flag
			creation_date	timestamp				-- creation date of account 
		)`)
	if err != nil {
		return err
	}
	_, err = conn.Exec(
		context.Background(),
		`create index idx_accounts_username on accounts (lower (username))`)
	if err != nil {
		return err
	}
	_, err = conn.Exec(
		context.Background(),
		`create index idx_accounts_email on accounts (lower (email))`)
	if err != nil {
		return err
	}

	// Create table 'email_confirmations' if not exists
	_, err = conn.Exec(
		context.Background(),
		`create table if not exists email_confirmations (
			account_id		integer references accounts(id) on delete cascade, -- reference to account id
			ref				uuid,					-- fragment of the email confirmation ref
			actual_date		timestamp,				-- date when the link expires
			primary key (account_id)
		)`)
	if err != nil {
		return err
	}

	// Create table 'phone_confirmations' if not exists
	_, err = conn.Exec(
		context.Background(),
		`create table if not exists phone_confirmations (
			account_id		integer references accounts(id) on delete cascade, -- reference to account id
			code			varchar(8),				-- phone confirmation code
			actual_date		timestamp,				-- date when the code expires
			primary key (account_id)
		)`)
	if err != nil {
		return err
	}

	return nil

}
