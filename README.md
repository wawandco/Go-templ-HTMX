## Setup local development
### Install templ
With Go 1.20 or greater installed, run:
```
go install github.com/a-h/templ/cmd/templ@latest
```
## Run
```
templ generate
go run .
```
or

```
make
```
This will take the Makefile in the root of the project and will run

```
- templ generate
- go run main.go
```


For the first time you will have to wait for the chromium binary to download.

Then, go to http://localhost:8080
