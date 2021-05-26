Example code base used for my presentation.

1. Install [go swag](https://github.com/swaggo/swag)

2. Generate doc before running:
`make generate-doc`
or
`swag init --generalInfo app/main.go --output /tmp/docs`

3. Run the code, http server will run on port `:8080`

4. Open http://localhost:8080/swagger
