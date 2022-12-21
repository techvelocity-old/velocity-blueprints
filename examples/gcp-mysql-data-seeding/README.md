# Create a snapshot from your existing database

## dump and zip data
```
mysqldump -u root -padmin example_db  --result-file=/tmp/example_db.sql
gzip /tmp/example_db.sql
```

## Upload snapshot to Velocity
```
veloctl snapshot put --target mysql-seeding-example -f example_db_a.sql.gz --default
```

## Create env
```
veloctl env create -f mysql-single-job.yaml
```

# Seed a Mysql database with example data

## Upload snapshot to Velocity
```
veloctl snapshot put --target mysql-seeding-example -f example_db_a.sql.gz --default
```

## Create env
```
veloctl env create -f mysql-single-job.yaml
```