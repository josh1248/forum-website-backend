Planned backend:
- go-gin for routing
- TODO: write into struct with jmoiron/sqlx?
- ~Postgres as the DB, with 3rd part driver at github.com/lib/pq~
  
**Edit:** I have decided to change the DB implementation from Postgres to SQLite3. A few reasons:
  - A previous consideration in using Postgres before SQLite was because SQLite in Go used to rely on a more complicated SQLite implementation with gcc at https://github.com/mattn/go-sqlite3, requiring complicated initialization. However I discovered a pure Go-based SQLite driver at https://github.com/glebarez/go-sqlite that has embedded SQLite and requires no environment or gcc shenanigans. Hence, setup is much easier. A small downside is the high number of dependencies that the latter driver uses (~20 total dependencies) compared to the former (2 dependencies only).
  - SQLite stores its databases within files, making for very easy portability. Researching online, I discovered that Postgres-based Go APIs are not very friendly to use. Some require Postgres to be installed before running, others require a containerized Postgres within Docker to work. SQLite offers a more straightforward path to implementing a portable API since it is serverless.
  - Lack of access controls in SQLite does not pose an issue for my small web project, since what I am doing is not sensitive.

  Self note:
  - PascalCase for public access functions is convention, camelCase for private functionality. https://golang.org/doc/effective_go.html#mixed-caps
  - cmd/server/main.go as the main starting point of the backend.
  - database folder is in charge of talking with the database server.
  - models contain the structs we expect to use, and form the M part of the MVC that handles data and interacts with the DB.
  - should probably keep a secret file for duplication in database, so that users can type their own password. TODO: Allow user setup to copy some preset configuration file, then fill in their own postgres details.

current design: MVC.
- routers direct URLs to appropriate controllers
- controllers communicate with models
- models communicate with SQLite for data.
- models return Golang data. controllers then convert Golang data to JSON data with appropriate headers using go-gin functions.

# How to initialize (self dev log)
Set up working import statements within my repository with the following:
- Commit this template repo into github. I used a dummy package called `repotest` with some trivial public functions to check for linkage later.
- follow https://go.dev/doc/tutorial/create-module, and run `go mod init <link>` in the root repo, where `<link>` was the github link to my root repo without the `http://` at the front. While you technically do not need to have your module name be an actual URL, it is a standard procedure so that others can get your packet with `go get -u ...` remotely without having to download them.
- get your URL link to your test package, like `repotest`, by checking with `go list ./...`
- In some other go file, use this list as your package statement to test this out.
- Run `go mod tidy`. This command will check all your Go files for import statements and download them as needed. It will then adjust the `go.mod` and `go.sum` files accordingly.
- Run your `main.go` file, in this case within `cmd/server/main.go`, to ensure that your packages have been set up nicely

## API testing
We can test our API's CRUD capabilities using the common HTTP requests `GET`, `POST`, `PATCH`, and `DELETE`.

For this project, the DB lies in `localhost:5432` and the API lies in `localhost:8080`. This is combined with the specified routes in `internal/router` (or `internal/routes`) to form the URL we can use.

Note: `localhost`, for our purposes, refers to the self-referential `127.0.0.1`.

### GET
I can test the routes `/users`, which should return a list of all users, by direct access to `localhost:8080/users`. Alternatively, use the following curl command:
```Bash
curl -X GET http://localhost:8080/users
```
curl defaults to `GET`, so `-X GET` is not necessary.

A much easier method involves using Postman. I am using the Postman extension within VSCode, and I will be using Postman from now on. I will revisit `curl` when I am more comfortable.

## Other Dev Logs

- Renamed data structure folder and package from `models` to `entities`. Instead, the Go package that interacts with the DB files will be called `models`.

SQLite controls: 
- access a SQLite db with `sqlite3 <filename>`.
- find the database layout with `.schema`.
- exit the sqlite3 shell with `;` to exit multiline query mode, then type `.exit`.

A key problem which I have faced is cleanly processing `SELECT` SQL queries. The built-in version is cumbersome and forces you to declare every variable in your table, which makes marshalling into JSON a nightmare. I considered using an ORM, like GORM, but I thought it was overkill for a project like this. I would like to learn the intricacies of running a web server and API before abstracting it away later. Additionally, I quite enjoy having control over the SQL queries I make. Hence, I will be using the lightweight extension package `sqlx` for my work.

Also learnt what CORS is after I couldnt load my API stuff, which was on `localhost:8080/`, onto my frontend at `localhost:3000/`: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#requests_with_credentials.
It is essentially a way to guard against attacks like cross site scripting (XSS) from other websites while permitting communcations.

Chose to implement a UNIQUE tag to the username field so that form re-submissions do not create spam repeatusernames.

It is now time to enable deleting usernames. I think I should implement authentication right away, so here goes. Following the tutorial at `https://www.youtube.com/watch?v=d4Y2DkKbxM0`. Although it uses mySQL, GORM, and Go Fiber, I know deep enough about my setup that it shouldnt be an issue.

First, we need a way to store passwords in our DB. Plaintext is definitely a no-go. Reading up, i shall be using the `bcrypt` library at `"golang.org/x/crypto/bcrypt"`. `bcrypt` not only helps in hashing passwords, but also in salting it (suffixing with random characters) to ensure unique hashes for the same password. It then stores the hash and cost into the hash itself for verification algorithms! It is named after the Blowfish cipher crypt that allows for safe password storage. (The industry standard now is to use 2FA, but this is an issue for another time.)

Encountered import cycle issues. My models were taking the hashing algorithm from my auth folder, but my auth folder was taking database data from my models. Acyclic import cycles are enforced to ensure fast compile times in Go. The solution is to make a separate package that handles hashing only under a utils folder.

sqlx's StructScan helps to convert data from a SQL query into a struct, and sqlx's NamedInsert does the reverse! how awesome.