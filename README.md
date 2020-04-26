# gin-graphql-postgres

## clone the repository 

```
git clone git@github.com:AnjaneyuluBatta505/gin-graphql-postgres.git
```

##  Run the GIN server

```
go serve server.go
```

##  Make changes to schema by editing the file `schema.graphqls`

After making changes run the below command to auto generate the `models` and `resolvers`. We need to implement the resolvers.

```
gqlgen generate
```

## Now, run the server and goto http://localhost:8080 to play with graphql.

References: https://gorm.io/, https://graphql.org/learn/
