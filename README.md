# github.com/lopolopen/t-fiber-kafka-gorm

## A template with fiber + kafka + gorm.

## 1. Install gonew.
```sh
go install golang.org/x/tools/cmd/gonew@latest
```

## 2. Use this template to create your project.
```sh
mkdir project_name && gonew github.com/lopolopen/t-fiber-kafka-gorm@v0.1 github.com/your_name/project_name ./project_name
```

## 3. Use make.
* Use make to build swagger documents.
```sh
make swag
```

* Use make to run your project.
```sh
make dev
```

* Use make to wire dependency injection
```sh
make sire
```

* Use make to generate all.
```sh
make gen
```

## 4. Run and access http://127.0.0.1:8081/swagger/index.htm
