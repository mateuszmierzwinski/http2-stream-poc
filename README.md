# HTTP2 Streams vs WebSocket vs. Workers pull PoC #

This is Proof of Concept using 3 techniques to get app notification from Golang application to fronted app.

To run this exmaple clone this repository and run prepared binaries for your platform.

After running this app navigate your browser to https://localhost:8443/

## Docker support ##
Pull repository and run inside:
```bash
docker build -t http2push .
docker run -p 8443:8443 -it http2push
```