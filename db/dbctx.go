package db

const dbName = "personsdb"
const collectionName = "user"
const dpphones = "dpphone"
const counting = "counting"
const port = 8000

var UserCollection, _ = GetMongoDbCollection(dbName, collectionName)
var DpPhoneCollection, _ = GetMongoDbCollection(dbName, dpphones)
var CountingCollection, _ = GetMongoDbCollection(dbName, counting)
