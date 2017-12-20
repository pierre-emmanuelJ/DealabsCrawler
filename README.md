# DealabsCrawler

DealabsCrawler is a crawler to have lasted comment by mail in this forum topic:

* https://www.dealabs.com/discussions/le-topic-des-erreurs-de-prix-1056379?page=9999

When a new comment appear you just recieve a mail.


### How to run it:

* First edit docker-compose.yml
```
environment:
      DEALABS_URL: https://www.dealabs.com/discussions/le-topic-des-erreurs-de-prix-1056379?page=9999
      DEALABS_HOSTNAME: smtp.gmail.com
      DEALABS_HOSTNAME_PORT: 587
      DEALABS_MAIL_SENDER: exemple@gmail.com
      DEALABS_MAIL_SENDER_PASSWORD: PasswordExemple
```
Put your sender mail and password and the good dealabs URL
* !! TODO implement mailingList.txt use by program !!
* Currently you have to add your mail recievers in the /mail/sendMail.go file

* Then just run with this command:
```
docker-compose up
```

