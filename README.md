# redlights

## Installing with docker

```bash
git clone https://gitlab.com/shizzuru/redlights.git
cd redlights
docker build -t redlights .
docker run -d -p 5600:5600 -p 8080:8080 --name=redlights redlights
```
## Running with Go
```bash
git clone https://gitlab.com/shizzuru/redlights.git
cd redlights
go run main.go
```