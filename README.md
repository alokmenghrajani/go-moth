# go-moth

A secure MOTH (message of the hour) web application. Enjoy!

## Build and run

```bash
head -c 32 /dev/urandom > utils/aes.key
docker build -t go-moth .
docker run -p 127.0.0.1:8000:8000 go-moth
```

And then connect to [http://localhost:8000/](http://localhost:8000/).
