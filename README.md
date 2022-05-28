# go-clean-architecture

```
go-clean-architecture
├── domain      Enterprise business rule
├── infra       Frameworks and drivers
├── interface   Interface adapters
├── usecase     Application business rule
└── schema      API schemas
```

## develop

### setup

Install tools.

```
$ make setup
```

Initialize DB.

```
$ make init-db
```

## clean up

Clean up DB.

```
$ make clean-db
```

### generate code and mock

```
$ make gen
```

### run test

```
$ make test
```

### run server

```
$ make start
```
