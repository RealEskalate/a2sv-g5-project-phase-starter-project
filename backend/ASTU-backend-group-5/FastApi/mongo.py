from pymongo import MongoClient

# Replace the URI with your MongoDB deployment's connection string
client = MongoClient("mongodb+srv://tamratk02:ec9H2S9gV1iR5lmd@cluster0.kj75i.mongodb.net/blogDB")

# Replace "mydatabase" with your database name
db = client["blogDB"]

# Optionally, access a collection
collection = db["users"]

results = collection.find({})
for result in results:
    print(result)

