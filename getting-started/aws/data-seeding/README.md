# AWS + PostgreSQL data seeding sample

Prerequisite: you need to be registered to velocity and have an AWS managed K8s cluster.

## Single job data seeding example

### Step 1: upload the snapshot file

```bash
veloctl snapshot put -f <seed-file-path> -t data-seeding-job --default
```

### Step 2: create an environment with data seeding job

```bash
veloctl env create -f https://raw.githubusercontent.com/techvelocity/velocity-blueprints/main/getting-started/aws/data-seeding/postgres-single-job.yaml --env-version v2
```

## Multiple job data seeding example

### Step 1: upload snapshot files

```bash
veloctl snapshot put -f <migrate-file-path> -t migrate-job --default
veloctl snapshot put -f <data-file-path> -t data-job --default
```

### Step 2: create an environment with two seeding jobs

```bash
veloctl env create -f https://raw.githubusercontent.com/techvelocity/velocity-blueprints/main/getting-started/aws/data-seeding/postgres-multi-job.yaml --env-version v2
```
