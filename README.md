# ccdata-server

**PROJECT IS STILL IN DEVELOPMENT AND NOT READY FOR USAGE**

## Description
Web-service to access to read-only comic book data written in Golang.

## Installation

1. Begin by cloning the project into your default workspace.

  ```bash
  cd ~/go/github.com/bartmika
  git clone https://github.com/LuchaComics/ccdata-server.git
  ```

2. Run the following in your **postgres** database console:

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
  export CCDATA_APP_SIGNING_KEY=lalalalalalalalalalala
  ```

5. This web service is powered by the **Grand Comics Database (GDC)** data. To begin, you will need to log into [comics.org](https://www.comics.org/download/) and download latest database.

6. Once file was downloaded, you will need import the latest database into either ``MariaDB`` or ``MySQL``.

7. Export the following tables as a **CSV** file from your database.
   * stddata_country
   * gcd_publisher
   * gcd_series
   * gcd_issue

8. Move the csv files to your projects **data/gcd** folder. When you finish you should have the following files:

  ```text
  github.com/luchacomics/ccdata-server/data/gcd/stddata_country
  github.com/luchacomics/ccdata-server/data/gcd/gcd_publisher
  github.com/luchacomics/ccdata-server/data/gcd/gcd_series
  github.com/luchacomics/ccdata-server/data/gcd/gcd_issue
  ```

9. Import our dependencies

  ```bash
  go run main.go import_country -f="./data/gcd/stddata_country.csv"
  go run main.go import_publisher -f="./data/gcd/gcd_publisher.csv"

  //TODO: ADD MORE
  ```

## Usage

**TODO: PLEASE FINISH**


```
go run main.go add_user --email="a@a.com" --password="123password" --fname="Bart" --lname="Mika" --state=1

go run main.go genkey --email="a@a.com"

go run main.go lookupkey --apikey=xxx

go run main.go state --email="a@a.com" --state=2

go run main.go serve
```

## License
This library is licensed under the **AGPLv3 license**. See [LICENSE](LICENSE) for more information. Copyrighted ©2021 [Lucha Comics](https://luchacomics.com/).


TODO: Once you serve, you can run the following:

http get 127.0.0.1:5000/api/v1/version
http get 127.0.0.1:5000/api/v1/version "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/countries page_token==0 page_size==250 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/publishers page_token==0 page_size==250 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/publisher/1 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
