# wecho
Simple Webserver Echo Server

## Usage

```
wecho & curl localhost:8666/
```
### Paths
* /404, Echo request, with a 404 status code
* /500, Echo request, with a 500 status code
* /400,  Echo request, with a 400 status code
### Parameters
* -p/--port, port number to run, defaults to *8666*

## Compile
* For current plataform
  ``` make build ```
* For release
  ``` make compile ```

## Dependencies
 * None *
