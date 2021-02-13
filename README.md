# ccdata-server

Web-service to access to read-only comic book data written in Golang.

## Installation

Please read the [documentation](https://github.com/LuchaComics/ccdata-server/wiki/Developer-Machine-Installation).

## Usage
The steps are as follows:

1. Clone the project.

    ```bash
    $ git clone https://github.com/LuchaComics/ccdata-server.git
    ```

2. Setup the database and the tables

3. Set the environmental variables

4. Start the server with the following code

    ```bash
    $ go run main.go serve
    ```

## Documentation

There are more **commands** available, you can find out about them [**here**](https://github.com/LuchaComics/ccdata-server/wiki/Command-Line-Interface).

The **HTTP application programming** interface can be read
[**here**](https://github.com/LuchaComics/ccdata-server/wiki/API-Docs).

## License
This library is licensed under the **AGPLv3 license**. See [LICENSE](LICENSE) for more information. Copyrighted ©2021 [Lucha Comics](https://luchacomics.com/).


TODO: Once you serve, you can run the following:

http get 127.0.0.1:5000/api/v1/version
http get 127.0.0.1:5000/api/v1/version "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/countries page_token==0 page_size==250 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/publishers page_token==0 page_size==250 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/publisher/1 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/series page_token==0 page_size==250 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/series/10000 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/issues page_token==0 page_size==250 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
http get 127.0.0.1:5000/api/v1/issue/1 "Authorization:Bearer $CCDATA_APP_ACCESS_TOKEN"
