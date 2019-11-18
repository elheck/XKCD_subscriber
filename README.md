# XKCD_subscriber

A very first try at writing a very basic XKCD email subscription service.

Info is retrieved from info.json, which is structured in the following way.

{
    "user":{
        "username": "sender_email_address",
        "pw": "my_password",
        "smtp_server": "the server",
        "smtp_port": "port"
    },
    "addresslist":[
        {
            "email": "a subscriber"
        }
    ]
}