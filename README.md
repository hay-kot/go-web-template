# Go Web Template

[![Go Build/Test](https://github.com/hay-kot/go-web-template/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/hay-kot/go-web-template/actions/workflows/go.yaml)

This Go Web Template is a simple starter template for a Go web application. It includes a simple web server API, as well as a start CLI to manage the web server/database. It should be noted that while while use of the standard library is a high priority, this template does make use of multiple external packages. It does however abide by the standard http handler pattern.

- [Go Web Template](#go-web-template)
  - [Template Includes](#template-includes)
    - [Web API](#web-api)
  - [Management CLI](#management-cli)
    - [Docker Setup](#docker-setup)
  - [Makefile](#makefile)
  - [How To Use: Application API](#how-to-use-application-api)
    - [Package Structre](#package-structre)
      - [app](#app)
      - [internal](#internal)
      - [pkgs](#pkgs)
      - [ent](#ent)
    - [Configuring The API](#configuring-the-api)
  - [How To Use: Application CLI](#how-to-use-application-cli)
    - [Manage Users](#manage-users)
      - [List Users](#list-users)
      - [Create User](#create-user)
      - [Delete User](#delete-user)

## Template Includes

- [ ] Test Coverage
- [ ] Basic CI/CD Workflow
- [ ] Swappable SQLite/Postgres backends

### Web API

- [x] Chi Router
  - [x] Auto log registered routes for easy debugging
- [x] OAuth2 JWT Tokens
  - [x] Email/Password Login
  - [x] Token Refresh
- [x] Ent for Database
  - [x] Basic users with hashed password storage
- [x] Yaml/Args Config

## Management CLI

- [x] CLI Interface

### Docker Setup

- [ ] Build and Run API
- [ ] Build and Setup CLI in path

## Makefile

- **Build and Run API:** `make api`
- **Build Production Image** `make prod`
- **Build CLI** `make cli`

## How To Use: Application API

### Package Structre

#### app

The App folder contains the main modules packages/applications that utilize the other packages. These are the applications that are compiled and shipped with the docker-image.

#### internal

Internal packages are used to provide the core functionality of the application that need to be shared across Applications _but_ are still tightly coupled to other packages or applications. These can often be bridges from the pkgs folder to the app folder to provide a common interface.

#### pkgs

The packages directory contains packages that are considered drop-in and are not tightly coupled to the application. These packages should provide a simple and easily describable feature. For example. The `hasher` package provides a Password Hashing function and checker and can easily be used in this application or any other. 

A good rule to follow is, if you can copy the code from one package to a completely. different project with no-modifications, it belongs here.

#### ent

As an exception to the above, this project adhears to the convention set by `Ent` we use a `ent` folder to contain the database schema. If you'd like to replace the Ent package with an alternative, you can review the repository layer in the `internal` folder.

[Checkout the Entgo.io Getting Started Page](https://entgo.io/docs/getting-started)

### Configuring The API

```yaml
# config.yml
web:
  port: 3001
  host: 127.0.0.1
database:
  driver: sqlite3
  sqlite-url: ./ent.db?_fk=1
#  postgres-url:
logger:
  level: debug
  file: api.log
```

## How To Use: Application CLI

### Manage Users

#### List Users

```bash
go run ./app/cli/*.go users list
```

#### Create User

**Development**

```bash
go run ./app/cli/*.go users add --name=hay-kot --password=password --email=hay-kot@pm.me --is-super
```

**Docker**

```bash
manage users add --name=hay-kot --password=password --email=hay-kot@pm.me
```

#### Delete User

**Development**

```bash
go run ./app/cli/*.go users delete --id=2
```

**Docker**

```bash
manage users delete --id=2
```
