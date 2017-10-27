## Usage

```sh
$ docker-compose up
$ docker-compose exec mysql mysql -u root -proot example < setup.sql
$ docker-compose exec mysql mysql -u root -p -e "set global max_connections = 5000;"
$ docker-compose exec mysql mysql -u root -p -e "truncate example.test;"

# Use docker
$ docker-compose exec mysql mysql -u root -p
# Use local client
$ mysql -u root -p --protocol=tcp --port=13306
```
