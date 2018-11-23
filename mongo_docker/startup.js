// development database
db.createUser(
  {
    user: "turing",
    pwd: "turing",
    roles: [
      { role: "readWrite", db: "turing" },
    ]
  }
);

db.createCollection("users");
db.createCollection("playgrounds");

db.users.createIndex({ "name": 1 }, { unique: true })

// testing database
db = db.getSiblingDB("testing")

db.createUser(
  {
    user: "turing",
    pwd: "turing",
    roles: [
      { role: "readWrite", db: "testing" },
    ]
  }
);

db.createCollection("users");
db.createCollection("playgrounds");

db.users.createIndex({ "name": 1 }, { unique: true })