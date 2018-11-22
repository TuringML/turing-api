db.createUser(
  {
    user: "turing",
    pwd: "turing",
    roles: [
      { role: "readWrite", db: "turing" }
    ]
  }
);

db.createCollection("users");
db.createCollection("playgrounds");

db.users.createIndex({ "name": 1 }, { unique: true })