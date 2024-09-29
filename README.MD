# GoFrame
refer to https://goframe.org
## gf-web
### gen ctrl
> Generate controller by defining API layer parameters
>
```shell
# Generate the corresponding ctrl and merge the ctrl in the same package into one
gf gen ctrl --merge=true
```
### gen dao
> Generate corresponding entity/do/dao by defining database tables
>
hack/config.yaml Define the generated path in the file (if there is a need for sub-modules)
```shell
gf gen dao 
```

### gen service
> Generate the corresponding service by defining the business logic in logic and inject it
>
hack/config.yaml Define the generated path in the file (if there is a need for sub-modules)
```shell
gf gen service 
```

## Support common CURD, PAGE, TX, ORM operations