## How to start

Run `docker-compose up -d --build` and you will be able to spin up both the APIs and MongoDB. 
The user that will be used to run all the queries is:

```
username: turing
password: turing
db: turing
```

very simple. The Authentication mechanism used is `SCRAM-SHA-1` and it's the default choice.