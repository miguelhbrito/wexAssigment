# wexassigment

## 🧰 Configuration

To install golang just follow the steps from website:
- https://golang.org/doc/install

To install docker and docker-compose just follow the steps from website:
- https://docs.docker.com/engine/install/
- https://docs.docker.com/compose/install/

## 🛠 How to use

Start API wex:
``` powershell
make config-up
```

```
To shutdown:
``` powershell
make config-down
```

## 🚀 Endpoints

##### `/wex` POST to create a new transaction
##### `/wex/{id}` GET a transaction by id
##### `/wex` GET to list all transactions