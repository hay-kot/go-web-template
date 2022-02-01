<h1 align="center"> Go Web Template</h1>
<p align="center" style="width: 100%">
  <a href="https://github.com/hay-kot/go-web-template/actions/workflows/go.yaml">
    <img src="https://github.com/hay-kot/go-web-template/actions/workflows/go.yaml/badge.svg?branch=master"/>
  </a>
  <a href="https://codecov.io/gh/hay-kot/go-web-template">
    <img src="https://codecov.io/gh/hay-kot/go-web-template/branch/master/graph/badge.svg?token=8EN4BQLIUS"/>
  </a>
</p>
    
This Go Web Template is a simple starter template for a Go web application. It includes a simple web server API, as well as a start CLI to manage the web server/database. It should be noted that while while use of the standard library is a high priority, this template does make use of multiple external packages. It does however abide by the standard http handler pattern.

- [Template Features](#template-features)
  - [General](#general)
  - [Mailer](#mailer)
  - [Admin / Superuser Management](#admin--superuser-management)
  - [User Services](#user-services)
    - [Admin](#admin)
    - [Self Service](#self-service)
  - [Logging](#logging)
  - [App Router](#app-router)
  - [Web Server](#web-server)
  - [Database](#database)
  - [Application Configuration](#application-configuration)
- [Management CLI](#management-cli)
  - [Docker Setup](#docker-setup)
- [Makefile](#makefile)
- [How To Use: Application API](#how-to-use-application-api)
  - [Package Structure](#package-structure)
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

## Template Features

### General

- [ ] Test Coverage (WIP)
- [x] Basic CI/CD Workflow

### Mailer

- [ ] Mailer classes for easy email sending
- [x] Starter email templates
  - [x] Activate Account
  - [ ] Password Reset
- [ ] Bulk Messages

### Admin / Superuser Management

### User Services

- [ ] User password hashing

#### Admin

- [ ] CRUD Operations for Users

#### Self Service

- [ ] User sign-up
- [ ] Require Activation by Email
- [ ] Stateful Token Auth
- [ ] Login/Logout
- [ ] Password Reset by Email

### Logging

- [x] Logging
- [x] File Logging + STDOUT
- [x] Request Logging (sugar in development structured in prod)
- [x] Dependency Free
- [x] Basic Structured Logging

### App Router

- [x] Built on Chi Router
- [x] Basic Middleware Stack
  - [x] Logging/Structured Logging
  - [x] RealIP
  - [x] RequestID
  - [x] Strip Trailing Slash
  - [x] Panic Recovery
  - [x] Timeout
  - [x] User Auth
  - [ ] Admin Auth
- [x] Auto log registered routes for easy debugging

### Web Server

- [x] Router agnostic
- [x] Graceful shutdown
  - [x] Finish HTTP requests with timeout
  - [x] Finish background tasks (no timeout)
- [x] Background Tasks
- [ ] Limited Worker Pool
- [x] Response Helpers
  - [x] Error response builder
  - [x] Utility responses
  - [x] Wrapper class for uniform responses

### Database

- [x] [Ent for Database](https://entgo.io/)

### Application Configuration

- [x] Yaml/CLI/ENV Configuration

<details>
<summary> CLI Args </summary>

```
Usage: api [options] [arguments]

OPTIONS
  --mode/$API_MODE                                    <string>            (default: development)
  --web-port/$API_WEB_PORT                            <string>            (default: 3000)
  --web-host/$API_WEB_HOST                            <string>            (default: 127.0.0.1)
  --database-driver/$API_DATABASE_DRIVER              <string>            (default: sqlite3)
  --database-sqlite-url/$API_DATABASE_SQLITE_URL      <string>            (default: file:ent?mode=memory&cache=shared&_fk=1)
  --database-postgres-url/$API_DATABASE_POSTGRES_URL  <string>
  --log-level/$API_LOG_LEVEL                          <string>            (default: debug)
  --log-file/$API_LOG_FILE                            <string>
  --mailer-host/$API_MAILER_HOST                      <string>
  --mailer-port/$API_MAILER_PORT                      <int>
  --mailer-username/$API_MAILER_USERNAME              <string>
  --mailer-password/$API_MAILER_PASSWORD              <string>
  --mailer-from/$API_MAILER_FROM                      <string>
  --seed-enabled/$API_SEED_ENABLED                    <bool>              (default: false)
  --seed-users/$API_SEED_USERS                        <value>,[value...]
  --help/-h
  display this help message
```

</details>

<details>
<summary> YAML Config </summary>

```yaml
# config.yml
---
mode: development
web:
  port: 3915
  host: 127.0.0.1
database:
  driver: sqlite3
  sqlite-url: ./ent.db?_fk=1
logger:
  level: debug
  file: api.log
mailer:
  host: smtp.example.com
  port: 465
  username:
  password:
  from: example@email.com
```

</details>

## Management CLI

- [ ] CLI Interface (Partial)

### Docker Setup

- [x] Build and Run API
- [x] Build and Setup CLI in path

## Makefile

- **Build and Run API:** `make api`
- **Build Production Image** `make prod`
- **Build CLI** `make cli`
- **Test** `make test`
- **Coverage** `make coverage`

## How To Use: Application API

### Package Structure

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

See the [Application Configuration](#application-configuration) section for more information.

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
