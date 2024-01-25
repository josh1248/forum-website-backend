# Forum Website Backend

Welcome! This is the backend instance of a planned web forum website. The lightweight `go-gin` framework is used to simplify HTTP server processes. It uses a SQLite3 database for serverless management as well.

### Note
NOTE: This repository is effectively a fork of my previous work for CVWO's 2024 Assignment at https://github.com/josh1248/forum-website-backend, which is itself a fork of CVWO's provided Go skeleton project at https://github.com/CVWO/sample-go-app. Self-forking is not allowed, and forking through an alternate account did not seem optimal. Hence, I copied my local files over to this repository.

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


