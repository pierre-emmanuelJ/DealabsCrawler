# DealabsCrawler

DealabsCrawler is a crawler to have lasted comment by mail in this forum topic:

* https://www.dealabs.com/discussions/le-topic-des-erreurs-de-prix-1056379?page=9999

When a new comment appear you recieve an email with formmated comment in HTML.


### How to run it:

* First edit docker-compose.yml

Put the good dealabs URL and mailinglist path file
```
environment:
      DEALABS_URL: https://www.dealabs.com/discussions/le-topic-des-erreurs-de-prix-1056379?page=9999
      DEALABS_HOSTNAME: smtp.gmail.com
      DEALABS_HOSTNAME_PORT: 587
      DEALABS_MAILINGLIST_PATH: ./mailinglist.txt
```

* Create a `mailinglist.txt` file in the repo

this is a `mailinglist.txt` file example:
```
example1@gmail.com
example2@gmail.com
example3@gmail.com

```

Edit dockerfile at L-13 to add you sender email credentials
```
CMD ["./dealabscrawler", "--sender-mail", "test@gmail.com", "--sender-mail-password", "password"]
```

* Then just run with this command:
```
$ docker-compose up
```

