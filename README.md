# What is this RSS feed aggregator?
It is a web server that allows clients to:

* Add RSS feeds to be collected
* Follow and unfollow RSS feeds that other users have added
* Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. User can use this project to keep up with their favorite blogs, news sites, podcasts, and more!

# Different endpoints of this aggregator server:

* "POST /v1/users" =>
    * body fromat:
        {
            "name": "xyz"
        }
    
    * response fromat:
        {
            "id": "039750f0-c577-4134-8c45-32b327f332e2",
            "created_at": "2024-04-17T15:17:19.527067Z",
            "updated_at": "2024-04-17T15:17:19.527067Z",
            "name": "xyz",
            "api_key": "f17867af30cc295c80808b6fd47a9554e8a60b6643502756de701ba8e2b4d220"
        }

* "GET /v1/users"
    * body format:
        there is no need to provide anything in request body
    
    * header format:
        in Auth at Bearer just give like this ApiKey  f17867af30cc295c80808b6fd47a9554e8a60b6643502756de701ba8e2b4d220
    
    * response format: 
        {
            "id": "039750f0-c577-4134-8c45-32b327f332e2",
            "created_at": "2024-04-17T15:17:19.527067Z",
            "updated_at": "2024-04-17T15:17:19.527067Z",
            "name": "xyz",
            "api_key": "f17867af30cc295c80808b6fd47a9554e8a60b6643502756de701ba8e2b4d220"
        }

