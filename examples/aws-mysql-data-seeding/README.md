# Create mysql db with data. 

## spin up mysql in docker:

```
docker run -e MYSQL_ROOT_PASSWORD=admin bitnami/mysql:5.7
```

## exec into container

## log in to mysql

```
mysql -u root -p
```

## create db
```
CREATE DATABASE example_db;
```

```
USE example_db;
```
## create table
```
CREATE TABLE users(first_name TINYTEXT, last_name TINYTEXT, user_id INT);
```

## insert sample data
```
INSERT users(first_name, last_name, user_id)
VALUES ("bob", "smith", 1), ("tom", "smith", 2), ("rob", "smith", 3);
```

## exit mysql session

```
exit
```

## dump and zip sample data
```
mysqldump -u root -padmin example_db  --result-file=/tmp/example_db.sql
gzip /tmp/example_db.sql
```
## copy file from Docker continer to host
```
docker cp <containerId>:/tmp/example_db.sql.gz /host/path/example_db.sql.gz
```

## Upload snapshot to Velocity
```
veloctl snapshot put --target mysql-seeding-example-b -f example_db_a.sql.gz --default
```

## Create env
```
veloctl env create -f mysql-single-job.yaml
```




### NOTE: if job succeeds, but data isn't seeded, you may need to replace values as follows:

```
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
```
with 
```
ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
```
