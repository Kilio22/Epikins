// Exemple:
// db.users.insertOne({"email": "foo.bar@epitech.eu", "roles": ["credentials", "log", "module", "projects", "users"], "jenkinsLogin": "foobar"});
// db.credentials.insertOne({"login": "foobar", "apiKey": "baz"});

db.users.insertOne({"email": "firstname.lastname@epitech.eu", "roles": ["credentials", "log", "module", "projects", "users"], "jenkinsLogin": ""}); 
db.credentials.insertOne({"login": "", "apiKey": ""});
