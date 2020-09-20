# Polymail test 

All the required configuration is in config.env for setting os envoirment veriables. (source ./config.env) will set all required config.

To run handlers test change directory to polymail/app/controler then run go test -v.

For functionality testing app is deployed at heroku with mongodb at https://polymail.herokuapp.com.

To perform send email sendgrid email service used. And currently sender_email for sendgrid must be taimoorshaukat6@gmail.com as i configured this as sender at send grid dashbaord.

App contain 5 routes of draft emails management. 

### /createmaildraft (POST)
body 

    {
      "payload": {
        "sender_email": "taimoorshaukat6@gmail.com",
	      "recipient_email": "taimoor.emallates@gmail.com",
        "message": "test update",
        "subject": "test"
      }
    }
      
###  /getmaildraft/{draftmailid} (GET)

###  /updatemaildraft/{draftmailid} (PUT)

body 

    {
      "payload": {
        "sender_email": "taimoorshaukat6@gmail.com",
	      "recipient_email": "taimoor.emallates@gmail.com",
        "message": "test update",
        "subject": "test"
      }
    }

###  /deletemaildraft/{draftmailid} (DELETE)

###  /senddraftemail/{draftmailid}  (PUT)
