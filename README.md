# Scouting System for Megiddo Lions in FTC 2022

## Usage

Firstly we need to install all the project dependencies.

```shell
go get -u github.com/NYTimes/gziphandler github.com/mattn/go-sqlite3 golang.org/x/crypto
```

After this we need to enter our [toa api key](https://theorangealliance.org/apidocs) to file secrets/secret.go.

```go
package secrets

const TOA_API_KEY = "--- YOUR API KEY ---"
```

And pass the event key as an argument to NewServer in [main.go](main.go).\
Also you will need to create a directory named robots.
Then after all that we finally can start our server.

```shell
go run .
```

Please note that the default port is 80 So if you have a problem with this you are welcome to change it in [server.go](server/server.go).\
And just like that you are done! The server is running and you are welcome to use it.

## Project Structure

### Database

We use SQLite for a database in our project so if you are willing to change it to other SQL server you can change it in
the [database module](server/database). Or if you aren't using SQL just rewrite the module. It's just a drop in replacement.

### The orange alliance API

We use it to fill all the teams for an event. How ever you can do without it in that case you can drop all it's api
code in the project.
The code for the integration is in the [toa_api module](server/toa_api) and is a part of the initialization of the server.
To remove it just remove the relevant code and remove the arguments [NewServer method](server/server.go) requires.

### The Server

Each web page is implemented by a function in the server module. Some are simple and just serve a file from the [www](www) folder.
Like [main.css](main.css) or favicon.ico.
But most of them require some backend code so they include a handleFunction. All of them are mentioned in [configHTTP](server/server.go).
In this function you can see all the available urls and where they are implemented.
For more reading about the server use the [site documentation](www/README.md).

## Final Notes

The code was writen in about a month while a FTC season was going so we didn't spend a lot of time on debugging.
So if you spot a bug please resport it or even better make a pull request with the fix.

## Credits

All of the backend - Gal\
Almost all of the front end - Ori\
Some help - Mayan
