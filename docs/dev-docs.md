### open shell container

```
  docker exec -it <container_id> bash
```

### open mongo service
```
  mongosh
```

### auth:
```  
  mongo -u "admin" -p "admin" --authenticationDatabase "admin"
```

### connect db:
``` 
  use default
```

### find table:
```
  db.table.find()
```