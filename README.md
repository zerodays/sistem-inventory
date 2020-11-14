# Sistem Inventory
Inventory microservice for sistem.

## Running
1. Start postgresql database
```
docker run -p 5433:5432 -e POSTGRES_PASSWORD=postgres -v sistem_inventory:/var/lib/postgresql/data -d postgres
```
2. Build the executable
```
make build
```
3. Start the server
```
./inventory serve
```

## Developing
Instead of building the executable by hand every time you change the code, you can
run the server with hot reload using [air](https://github.com/cosmtrek/air).
After installing [air](https://github.com/cosmtrek/air) run `air` to do that.

When changing default configuration or database migrations in `configs` directory, you should embed the data
in binary executable using the tool [go-bindata](https://github.com/go-bindata/go-bindata). To do that run
```
make bindata
```

Example config change (using environment variables):
```bash
# export SISTEM_INVENTORY_{section}_{fieldname}={value}
export SISTEM_INVENTORY_DATABASE_URL="127.0.0.1"
# run server again
./inventory serve
```