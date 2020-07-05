# B"H


- [URL](https://devcenter.heroku.com/articles/heroku-postgresql#using-the-cli)
- [](https://gist.github.com/Kartones/dd3ff5ec5ea238d4c546)



---

#### `heroku pg:info` 
- Show all **PostgreSQL** databases provisioned by your application and the identifying characteristics of each (such as database size, status, number of tables, and PG version).

**Sample Output:**
```
Plan:                  Hobby-dev
Status:                Available
Connections:           1/20
PG Version:            12.3
Created:               2020-02-19 12:36 UTC
Data Size:             8.4 MB
Tables:                4
Rows:                  17/10000 (In compliance)
Fork/Follow:           Unsupported
Rollback:              Unsupported
Continuous Protection: Off
Add-on:                postgresql-corrugated-45675
```


---

#### `heroku pg:credentials:url`
- Heroku Postgres provides convenient access to the credentials and location of your database should you want to use a GUI to access your instance.
- https://devcenter.heroku.com/articles/heroku-postgresql#pg-credentials




---

#### `heroku pg:psql`
- Establishes a `psql` session with your remote database.
- You must already have PostgreSQL installed on your system.
- If you have multiple databases, see URL above (simple - just pass db as param to CLI command...)


--- 

#### Common commands with `psql`

```
\q: Quit/Exit

\c __database__: Connect to a database

TABLE INFO:
\d __table__  : Show table definition (columns, etc.) including triggers
\d+ __table__ : More detailed table definition including description and physical disk size
\dt *.*       : List tables from all schemas (if *.* is omitted will only show SEARCH_PATH ones)

DB INFO:
\l: List databases

QUERIES:
\copy (SELECT * FROM __table_name__) TO 'file_path_and_name.csv' WITH CSV: Export a table as CSV

OTHER:
\dy: List events
\df: List functions
\di: List indexes
\dn: List schemas
\dT+: List all data types
\dv: List views
\dx: List all extensions installed
\df+ __function__ : Show function SQL code.
\x: Pretty-format query results instead of the not-so-useful ASCII tables

\des+: List all foreign servers
\dE[S+]: List all foreign tables
```