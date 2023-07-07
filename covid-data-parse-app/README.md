# parseApp

This application parses the covid data from the internet and writes it to the database in the docker container every 10 minutes.

- MariaDB Docker Container

```docker
merlins@jimmmmy ~$ docker run --name mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password1 -e MYSQL_DATABASE=db -d mariadb:latest
merlins@jimmmmy ~$ docker ps
```

- Create the database you will use in the next step by importing the .sql file into the mariadb container you have installed.

```bash
merlins@jimmmmy ~$ docker exec -it mariadb bash
merlins@jimmmmy ~$ mysql -u root -ppassword1
>CREATE DATABASE new_database;
merlins@jimmmmy ~$ docker exec -i mariadb mysql -uroot -ppassword1 --database=new_database < /home/merlins/parseApp/covid19-table.sql
```

- Script code that takes the data from address [https://www.worldometers.info/coronavirus/](https://www.worldometers.info/coronavirus/) and parses it and writes it to both .csv and .sql files.

###### Run a script

```bash
python main.py
```

###### Listing directory

```bash
merlins@jimmmmy ~/parseApp$ tree
.
├── automaticScripts
├── covid-19-data.csv
├── covid19-table.sql
├── deneme.sql
├── main.py
└── README.md

0 directories, 6 files
```

- We import the .sql file you have created into the database in the MariaDB container.

```bash
merlins@jimmmmy ~$ docker exec -i mariadb mysql -uroot -ppassword1 --database=challenge < /home/merlins/parseApp/deneme.sql
```

or

```bash
merlins@jimmmmy ~$ docker cp deneme.sql bd3171f3a711:/ && docker exec -it mariadb bash
merlins@jimmmmy ~$ mysql -uroot -ppassword challenge < deneme.sql
```

- We have created a shell script that repeats the operations you have done every 10 minutes.

```bash
#!/bin/bash
cmd="python main.py && docker exec -i mariadb mysql -uroot -ppassword1 --database=challenge < /home/merlins/Downloads/covid19-table.sql"
while true
do
    $cmd &
    sleep 600
done
```
