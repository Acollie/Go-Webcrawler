package sqlRelational

const (
	updateWebsite   = ``
	checkForWebsite = ``

	addWebsite = `CREATE TABLE website (
    baseurl TEXT NOT NULL UNIQUE, -- UNIQUE constraint for baseurl
    promancevalue INT, -- Adjusted integer type to INT
    uuid CHAR(36) NOT NULL DEFAULT (UUID()), -- MariaDB uses UUID() for generating UUIDs
    PRIMARY KEY (uuid) -- Primary key defined as UUID
);

ALTER TABLE website
    OWNER TO 'root'; -- Adjusted the owner assignment syntax for MariaDB compatibility

CREATE INDEX website_baseurl_index
    ON website (baseurl); -- Index for baseurl`

	addPage = `create table page
	(
		pageurl       varchar                            not null,
		title         varchar,
		body          text,
		baseurl       varchar,
		uuid          uuid    default uuid_generate_v4() not null
			primary key,
		promancevalue integer default 1                  not null
	);

	alter table page
		owner to root;

	create index page_pageurl_index
		on page (pageurl);`

	dropWebsite = `drop table website;`
	dropPage    = `drop table page;`
)
