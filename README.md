# http-static-server
Trivial golang webserver to serve a single static file on a given port

# Usage

```
./http-static-server [options]
-p %port% | --port %port%     The port to serve on (defaults to 8080)
-f %path% | --file %path%     The static file to serve (defaults to index.html)
```

# Example

```
./http-static-server --port 80 --file example.html
```
