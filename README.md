# Revel performance
Revel performance verification repository.

# Docker image
Docker pull
```
$ docker pull tomohito/revel-performance
```

Set Environment variable.
```
DB_SOURCE={database source}
```

Docker run.
```
$ docker run -p 9000:9000 tomohito/revel-performance
```

# Build by myself
Git clone.
```
$ git clone git@github.com:tomoyane/revel-performance.git
```

Docker build.
```
$ docker build -t {username}/revel-performance:latest .
```

Set Environment variable.
```
DB_SOURCE={database source}
```

Docker run.
```
$ docker run -p 9000:9000 {username}/revel-performance
```

# Performance verification sql
```
CREATE TABLE items (
  id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(128) NOT NULL,
  category VARCHAR(128) NOT NULL,
  created_at DATETIME(6),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
