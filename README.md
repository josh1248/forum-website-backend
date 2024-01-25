# CVWO 2024 Assignment Backend

Name: Josh Thoo Jen Sen

This is the progress of my Golang backend as of 25 January 2024.

I will be doing my own changes to enhance this beyond the deadline, as I think it is worth learning. Please consider hopping by here to see my product, which I will improve incrementally:
https://github.com/josh1248/forum-website-backend

This Golang app can handle:
- Users: View of users, Creation of users, login of users (with generation of JWT tokens in cookies for successful logins)
- Posts: View of posts

I was unfortunately unable to complete:
- Users: I managed to generate a JWT in a cookie upon successful authentication. However, I have yet to implement post creation by reading off this JWT cookie. JWT verification check implementation is not yet complete.
- Posts: route to create posts.
- Comments in posts.

## Getting Started

### Installing Go

Download and install Go by following the instructions [here](https://go.dev/doc/install).

### Running the backend
1. [Fork](https://docs.github.com/en/get-started/quickstart/fork-a-repo#forking-a-repository) this repo.
2. [Clone](https://docs.github.com/en/get-started/quickstart/fork-a-repo#cloning-your-forked-repository) **your** forked repo.
3. Open your terminal and navigate to the directory containing your cloned project.
4. Duplicate `.exampleenv` in the root repository.
5. Rename this file to `.env`, in which the server's secret key will be hidden in. It should not be shared with anyone, and will be hidden by this repository's .gitignore.
6. Run `go run cmd/server/main.go`.
7. To ensure that your API is running, check `localhost:8080/api/posts`. You should see a list of 4 posts.
8. Return to the [frontend setup](https://github.com/josh1248/cvwo-assignment-24-frontend).

### Implementation Details

This backend uses Golang. It uses SQLite3 for its database, which is connected via a 3rd party SQLite driver at `github.com/glebarez/go-sqlite`, augmented with `sqlx` for more developer-friendly SQL queries.

This backend uses `go-gin` as its lightweight framework for HTTP server functionality.

This is the main file structure. It is roughly based on the MVC framework. It is generated with the `file-tree-generator` extension at VSCode.

```
📦cvwo-assignment-24-backend
 ┣ 📂cmd
 ┃ ┣ 📂server
 ┃ ┃ ┗ 📜main.go //entry point.
 ┣ 📂internal
 ┃ ┣ 📂api remnants from skeleton golang fork mean. not yet used.
 ┃ ┃ ┗ 📜api.go 
 ┃ ┣ 📂auth
 ┃ ┃ ┣ 📜bcrypt.go //conversion of plaintext passwords to hashed + salted passwords for db storage.
 ┃ ┃ ┗ 📜jwt.go //generates JWT upon successful login
 ┃ ┣ 📂controllers
 ┃ ┃ ┣ 📜posts.go
 ┃ ┃ ┗ 📜users.go
 ┃ ┣ 📂db
 ┃ ┃ ┗ 📜forumdb //SQLite3 database data
 ┃ ┣ 📂entities //Golang types
 ┃ ┃ ┣ 📜comment.go
 ┃ ┃ ┣ 📜post.go
 ┃ ┃ ┗ 📜user.go
 ┃ ┣ 📂models //M part of MVC. Interacts with the db.
 ┃ ┃ ┣ 📜connect.go
 ┃ ┃ ┣ 📜posts.go
 ┃ ┃ ┣ 📜reset.go //clears junk data.
 ┃ ┃ ┗ 📜users.go
 ┃ ┣ 📂router
 ┃ ┃ ┗ 📜router.go
 ┃ ┗ 📂routes
 ┃ ┃ ┗ 📜routes.go
 ┣ 📂repotest
 ┃ ┗ 📜hello.go //used to verify remote import statements initially. not in use.
 ┣ 📜.gitignore
 ┣ 📜ERD_snapshot_050124.png
 ┣ 📜MasterKey.env //JWT signing secrets
 ┣ 📜MasterKey.exampleenv //for user setup
 ┣ 📜README.md
 ┣ 📜devlog.md
 ┣ 📜go.mod
 ┗ 📜go.sum
```