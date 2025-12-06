```bash
$ docker pull postgres:17-alpine
```

```bash
$ docker run --name simpleBank -e POSTGRES_USER=deng -e POSTGRES_PASSWORD=deng -p 5432:5432 -d postgres:17-alpine
```

```bash
$ docker exec -it simpleBank psql -U deng
```

```bash
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-arm64.tar.gz | tar xvz
```

```bash
$ chmod +x migrate
```

```bash
$ sudo mv migrate /usr/local/bin/
```

```bash
$ migrate create -ext sql -dir db/migration -seq init_schema
```

```bash
$ docker start simpleBank
```

```bash
$ docker stop simpleBank
```

```bash
$ docker exec -it simpleBank /bin/sh
$ createdb --username=deng --owner=deng simple_bank
```

### **常用 PostgreSQL 命令**

1. **查看当前数据库中的所有表**：

   ```sql
   \dt
   ```

2. **查看当前数据库中的所有角色**：

   ```sql
   \du
   ```

3. **查看数据库的连接信息**：

   ```sql
   \conninfo
   ```

4. **创建新表**： 示例 SQL 查询，用于创建一个简单的 `accounts` 表：

   ```sql
   CREATE TABLE accounts (
       id SERIAL PRIMARY KEY,
       owner VARCHAR(100),
       balance DECIMAL(10, 2)
   );
   ```

5. **插入数据**： 插入一条记录到 `accounts` 表：

   ```sql
   INSERT INTO accounts (owner, balance) VALUES ('John Doe', 1000.00);
   ```

6. **查看表中的数据**： 查询 `accounts` 表中的所有数据：

   ```sql
   SELECT * FROM accounts;
   ```

7. **退出 psql**： 输入以下命令退出 PostgreSQL：

   ```sql
   \q
   ```





```bash
$ psql -U deng -d simple_bank

$ dropdb -U deng simple_bank

$ docker exec -it simpleBank createdb --username=deng --owner=deng simple_bank
```





在wsl上创建的postgres，在wsl上查看ip地址，连接那个ip

```bash
$ SERVER_ADDRESS=0.0.0.0:8081 make server
#覆盖配置文件中的变量


```bash
mockgen -package mockdb -destination db/mock/store.go simple-bank/db/sqlc Store

```bash
 docker build -t simplebank:latest .

 ```bash
 docker network create bank-network
 docker network connect bank-network simpleBank
 docker network inspect bank-network
 docker container inspect simpleBank
 docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE=postgresql://deng:deng@simpleBank:5432/simple_bank?sslmode=disable simplebank:latest
