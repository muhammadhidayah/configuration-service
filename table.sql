CREATE TABLE public.configuration_client (
	config_client_id serial NOT NULL,
	config_client_uuid varchar(255) NULL,
	multiple_language_id int NULL,
	appname varchar(255) NULL,
	report_title varchar(255) NULL,
	company_subs_id varchar(255) NULL,
	is_config_deleted bool NULL,
	CONSTRAINT configuration_client_pk PRIMARY KEY (config_client_id)
);

CREATE TABLE public.configuration_global (
	config_global_id serial NOT NULL,
	footertext text NULL,
	server_smpt text NULL,
	ssl bool NULL,
	port int8 NULL,
	is_auth bool NULL,
	username varchar(255) NULL,
	"password" varchar(255) NULL,
	is_active bool NULL,
	CONSTRAINT configuration_global_pk PRIMARY KEY (config_global_id)
);