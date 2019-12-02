[![Codacy Badge](https://api.codacy.com/project/badge/Grade/08fcf86b05874e4189f35b5f54d6b164)](https://www.codacy.com/manual/elheck/XKCD_subscriber?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=elheck/XKCD_subscriber&amp;utm_campaign=Badge_Grade)

# XKCD_subscriber

A very first try at writing a very basic XKCD email subscription service.

Info is retrieved from info.json, which is structured in the following way.

    "username": "username",
    "pw": "password",
    "smtp_server": "smtpserver",
    "smtp_port": "port",
    "addresslist":[  
        "subscriber1@fakemail.com",
        "subscriber2@fakemail.com"
    ]
