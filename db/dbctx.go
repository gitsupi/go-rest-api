package db

const dbName = "personsdb"
const collectionName = "user"
const dpphones = "dpphone"
const port = 8000

var UserCollection, _ = GetMongoDbCollection(dbName, collectionName)
var DpPhoneCollection, _ = GetMongoDbCollection(dbName, dpphones)
