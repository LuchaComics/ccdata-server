# ccdata-server

**PROJECT IS STILL IN DEVELOPMENT AND NOT READY FOR USAGE**

## Description
Web-service to access to read-only comic book data written in Golang.

## Pre-Installation

Web-service is powered by the **Grand Comics Database (GDC)** data. To begin, you will need to run the following:

1. Please log into [comics.org](https://www.comics.org/download/) and download latest database.
2. Import latest database.
3. Export the following tables as a **CSV** file:
   * stddata_country
   * gcd_publisher
   * gcd_series
   * gcd_issue
4. Move csv files to your projects **data/gcd** folder.

## Installation

1. Begin by cloning the project into your default workspace.

  ```bash
  cd ~/go/github.com/bartmika
  git clone https://github.com/LuchaComics/ccdata-server.git
  ```

2. Rn the following in your **postgres** database console:

  ```sql
  drop database ccdata_db;
  create database ccdata_db;
  \c ccdata_db;
  CREATE USER golang WITH PASSWORD '123password';
  GRANT ALL PRIVILEGES ON DATABASE ccdata_db to golang;
  ALTER USER golang CREATEDB;
  ALTER ROLE golang SUPERUSER;
  ```

3. Once the above SQL code has been executed, you will need to manually copy and paste the SQL code found in [this file](https://github.com/LuchaComics/ccdata-server/blob/master/migrations/0001_initial_up.sql) into your ``Postgres console`` so our applications database gets created.

4. Setup our environment variables by running the following into your console:

  ```bash
  export CCDATA_DB_HOST=localhost
  export CCDATA_DB_PORT=5432
  export CCDATA_DB_USER=golang
  export CCDATA_DB_PASSWORD=123password
  export CCDATA_DB_NAME=ccdata_db
  ```

5. Import our dependencies

  ```bash
  go run main.go import_country -f="./data/gcd/stddata_country.csv"

  //TODO: ADD MORE
  ```

## License
This library is licensed under the **AGPLv3 license**. See [LICENSE](LICENSE) for more information. Copyrighted ©2021 [Lucha Comics](https://luchacomics.com/).
