# What is this RSS feed aggregator?
It is a web server that allows clients to:

* Add RSS feeds to be collected
* Follow and unfollow RSS feeds that other users have added
* Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. User can use this project to keep up with their favorite blogs, news sites, podcasts, and more!

# How to deploy this aggregator in your local machine?
## install [postgresql](postgresql.org)

## Copy this project

```bash
git clone github.com/nabin3/RSS-feed-agg
cd RSS-feed-agg
```

## Installing dependencies 
```bash
go get github.com/google/uuid@v1.6.0
go get github.com/joho/godotenv@v1.5.1
go get github.com/lib/pq@v1.10.9
```

## Installing goose to migrate database
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

# Migrating database
```bash
cd sql/schema/
goose postgres postgres://username:password@localhost:database_port_number/database_name up
```

## Creating environement
create .env file in root directory inside RSS-feed-agg. Now inside that file mention
    * PORT number on which you want to set the server for listening 
    * CONN, database connection string(protocol://database_username:password@localhost:database_port_number/database_name?ssl_mode=disable)


## Run
```bash
go build 
./RSS_feed_agg
```

And the server starts to listen on mentioned port

# Different endpoints of this aggregator server:

## "POST /v1/users" 
    This endpoint is for creating a new user

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

## "GET /v1/users"
    This end point will provide details about an user, whose api_key will be provided

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

## "POST /v1/feeds"
    This endpoint is for adding a feed, whoose api_key will be provided he will autmatically follow this given feed

    * body format:
        {
            "name": "sunday_suspense",
            "url": "https://radio_mirrchi.fm"
        }

    * header format:
        in Auth at Bearer just give like this ApiKey  f17867af30cc295c80808b6fd47a9554e8a60b6643502756de701ba8e2b4d220

    * response format:
        {
            "feed": {
                "id": "3bcd8c59-4762-4256-af4a-8d8030dc06e5",
                "created_at": "2024-04-25T14:17:16.036641Z",
                "updated_at": "2024-04-25T14:17:16.036641Z",
                "name": "sunday_suspense",
                "url": "https://radio_mirrchi.fm",
                "user_id": "41e1ad3b-a5c1-4c41-80ff-28303d97ea0c"
            },
            "feed_follow": {
                "id": "52c4e485-02e8-4dea-b28e-81b847ca9b69",
                "created_at": "2024-04-25T14:17:16.036641Z",
                "updated_at": "2024-04-25T14:17:16.036641Z",
                "user_id": "41e1ad3b-a5c1-4c41-80ff-28303d97ea0c",
                "feed_id": "3bcd8c59-4762-4256-af4a-8d8030dc06e5"
            }
        }

## "GET /v1/all-feeds"
    This endpoint will fetch all feeds added by all users

    nothing needed in body and header, GET method to this end-point will retrieve all feeds

## "POST /v1/feed_follows"
    This endpoint is for creating feed_follow, whose api_key will be provided that user will start to follow the feed, whick id will be provided in request body

    * header format: 
        ApiKey 35hk34k43k59879ffdsf.........
        need api_key

    * body format:
        {
            "feed_id": "7316015e-8a79-4b67-b961-452b4781c27e"
        }

    * response format:
        {
            "id": "6b2cca69-f4c9-4379-94f4-0c664de2ca24",
            "created_at": "2024-04-25T16:01:59.405443Z",
            "updated_at": "2024-04-25T16:01:59.405444Z",
            "user_id": "48ea7dbc-56f4-4ba0-997c-a6f00fbe91e0",
            "feed_id": "7316015e-8a79-4b67-b961-452b4781c27e"
        }

## "DELETE /v1/feed_follows/{feedFollowID}"
   This endpoint is for deleting a feed_follow, by providing feed_follow id.
   Nothing needed in header, body. Just need a query parameter holding feedFollowID

## "GET /v1/feed_follows"
    This endpoint is for retrieveing all feed_follows of a user whose api_key is provided

    * header: 
        ApiKey gjsagj664871684...............
        Need api_key only and this endpoint will serve feed_follows of the user whoose api_key was is included in header

    * response format:
        [
            {
                "id": "6b2cca69-f4c9-4379-94f4-0c664de2ca24",
                "created_at": "2024-04-25T16:01:59.405443Z",
                "updated_at": "2024-04-25T16:01:59.405444Z",
                "user_id": "48ea7dbc-56f4-4ba0-997c-a6f00fbe91e0",
                "feed_id": "7316015e-8a79-4b67-b961-452b4781c27e"
            }
        ]

## "GET /v1/posts/{limit}
    This endpoint will serve posts retrieved from a feed followed by a user whoose api_key is provided

    * header:
        ApiKey gjsagj664871684...............
        Need api_key only and api key will serve all posts(latest posts first) in feeds followed by the user whose api_key is given     
    
    * query parameter:
        If passed limit query parameter then endpoint will serve that number of posts and if no query parameter is given then endpoint will serve only 5 posts

    * response format:
        [
            {
                "id": "a524805b-8378-43c3-9028-e27fc85d7b5b",
                "title": "The Boot.dev Beat. May 2024",
                "url": "https://blog.boot.dev/news/bootdev-beat-2024-05/",
                "description": "A new Pub/Sub Architecture course, lootable chests, and ThePrimeagen&rsquo;s Git course is only a couple weeks away.",
                "published_at": "2024-05-01T00:00:00Z"
            },
            {
                "id": "36e5daf3-68ef-45ae-b95c-a006f4877ee3",
                "title": "The Boot.dev Beat. April 2024",
                "url": "https://blog.boot.dev/news/bootdev-beat-2024-04/",
                "description": "Pythogoras returned in our second community-wide boss battle. He was vanquished, and there was much rejoicing.",
                "published_at": "2024-04-03T00:00:00Z"
            },
        ]