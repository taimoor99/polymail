
# Polymail  

A simple CRUD app for sending draft mails.

All the required configuration is in config.env for setting os envoirment veriables. (source ./config.env) will set all required config.

To run handlers test change directory to polymail/app/controler then run go test -v.

App contain 5 routes of draft emails management. 

{draftmailid} is mongodb document _id.

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
