# Open Source Issue Träcking System - opists
Track the bugs of your software projects.
Plan new features simply and easily via the user interface*. REST APIs are also available for developers.

\* The user interface is work in progress.

## Planned features:
UI for create projects and manage issues same as REST API.

Reinforce opists with the [postgresql database](https://www.postgresql.org/).

User management.
* Add/Read/Update/Delete users via REST API and UI.
* Assign roles to users/Remove roles from users/List roles of users.
* Add/Read/Update/Delete members of projects.

The projects can be planned using roadmaps for all components and versions.

Boards for agile development (focused on Scrum).
* Backlog Board
* Sprint Backlog Board
* Kanban Board (todo, in progress, done)
* Roadmap Board
* Tester Board

Testers can plan and execute a test cycle.
DevOps can evaluate metrics via [prometheus](https://prometheus.io/).

# Security

Add your admin and users to the configuration via example `config.yml`:

````yaml
server:
  host: "0.0.0.0"
  port: 8080
opists:
  security:
    enabled: true
    enabled_user_logging: true # logs the current caller user
    authz_model_file_path: "authz_model.conf"  # casbins configuration mapping
    authz_policy_file_path: "authz_policy.csv" # defined permissions, users, roles
    admin_username: "admin" # role:superadmin
    admin_password: "password"
    users:
      - username: "bob" #role:user
        password: "password"
      - username: "alice" #role:admin
        password: "password"
````

# User Interface
see [ui/docs/UserInterface.md](ui/docs/UserInterface.md)

# Build
see [apidoc/Build.md](apidoc/Build.md)

## Testing API with cucumber godog
Switch to the test folder and execute:
```
cd test/api
go test
```
Output:
```
2020/11/13 13:43:29 GET  GetAllProject 0s
...2020/11/13 13:43:29 GET  GetProject 0s
... 6


2 scenarios (2 passed)
6 steps (6 passed)
32.9977ms
testing: warning: no tests to run
PASS
ok  	github.com/DevBoxFanBoy/opists/tests	0.322s [no tests to run]

Process finished with exit code 0
```