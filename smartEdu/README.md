# Smartedu Projesi

Bu projemizde Smartedu eğitim portali [template](https://html.design/download/smartedu-education-template/) kullanılmıştır. Bu projemizde üzerinde konuşulacak bazı konu başlıkları:

- Flash mesajlar
- Kullanıcı kontrolü ve doğrulama
- Kullanıcı girişi
- Özel Middleware
- Şifre gizleme
- Kategori oluşturma
- "Search" alanı
- Mail gönderimi
- Session kavramı
- Admin alanı
- Postman kullanımı
- Modeller arası ilişkiler
- Slugify

## Moduls

- ORM->mongoose
- teplateEngine->ejs
- Docker

```Bash
# start container
docker run --name mongodb -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=root -d -p 27017:27017 mongo
# entry the container
docker exec -it mongodb mongo 
# list databases 
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
test    0.000GB
# switched to db
> use test
switched to db test
# list collections
> show collections
categories
courses
users
# show collections
> db.users.find()
{ "_id" : ObjectId("61cc53de40a4c16246c95e7c"), "name" : "test-2", "email" : "test-2@test.com", "password" : "$2b$10$JBSWDvm/DHs7nnCce1lvBOAaiQR883lvMY4ZrJgvIgrVs.puRYY6u", "__v" : 0 }

```
