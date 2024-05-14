# What is this RSS feed aggregator?
It is a web server that allows clients to:

* Add RSS feeds to be collected
* Follow and unfollow RSS feeds that other users have added
* Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. User can use this project to keep up with their favorite blogs, news sites, podcasts, and more!

# How to deploy this aggregator in your local machine?
## install [docker](https://docs.docker.com/get-docker/), if you havn't already

## install [golang](https://go.dev/doc/install)

## Copy this project

```bash
git clone github.com/nabin3/RSS-feed-agg
cd RSS-feed-agg
```

## Setup database
```bash
docker pull postgres:15-alpine
make postgresinit
make createdb
```
## Installing goose to migrate database
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
## Migrate database
```bash
make migrateupdb
```

## Installing dependencies 
```bash
go mod download
```

## Creating environement
create .env file in root directory inside RSS-feed-agg.

    PORT="8080"
    CONN="postgres://root:postgres@localhost:5434/blogator?sslmode=disable"

## Run
```bash
go build 
./RSS_feed_agg
```

And the server starts to listen on mentioned port

# Different endpoints of this server:

## /v1/readiness
This endpoint is checking readiness of the server

    * Method: GET
    * Request Content-Type: empty
    * Response Content-Type: application/json

## /v1/users 
This endpoint is for creating a new user

    * Method: GET
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    
    Example Request:
        {
            "name": "Mr.X"
        }
        
    Example Response:
    {
        "id": "66a959f0-efd5-4572-8e80-67a336e6b995",
        "created_at": "2024-05-12T03:40:38.646767Z",
        "updated_at": "2024-05-12T03:40:38.646767Z",
        "name": "Mr.X",
        "api_key": "dc71925beb7fb83060528594c96b671ca8486ea4be036e49cebfe08c3e7fd498"
    } 

## /v1/users
This end point will provide details about an user, whose api_key will be provided

    * Method: GET
    * Request Content-Type: empty
    * Authorization: Bearer api_key 

    Example Authorization: Bearer
        ApiKey  dc71925beb7fb83060528594c96b671ca8486ea4be036e49cebfe08c3e7fd498
        (Note: Give "ApiKey" word before the key string in every request which requires authentication, it's mandetory)
        
    Example Response: 
        {
            "id": "66a959f0-efd5-4572-8e80-67a336e6b995",
            "created_at": "2024-05-12T03:40:38.646767Z",
            "updated_at": "2024-05-12T03:40:38.646767Z",
            "name": "Mr.X",
            "api_key": "dc71925beb7fb83060528594c96b671ca8486ea4be036e49cebfe08c3e7fd498"        
        }

## /v1/feeds
This endpoint is for adding a feed, whoose api_key will be provided and he will autmatically follow this given feed

    * Method: POST
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    * Authorization: Bearer api_key 

    Example Authorization: Bearer
        ApiKey  dc71925beb7fb83060528594c96b671ca8486ea4be036e49cebfe08c3e7fd498

    Example Request:
            {
                "name": "boot.dev blogs",
                "url": "https://blog.boot.dev/index.xml"
            }

    Example Response:
        {
        "feed": {
            "id": "c8a6a42c-a571-441d-b1de-122f0fd6a4c6",
            "created_at": "2024-05-12T03:57:07.608693Z",
            "updated_at": "2024-05-12T03:57:07.608693Z",
            "name": "boot.dev blogs",
            "url": "https://blog.boot.dev/index.xml",
            "user_id": "66a959f0-efd5-4572-8e80-67a336e6b995"
        },
        "feed_follow": {
            "id": "c179f1f2-d4fb-443d-ab44-69058b703d94",
            "created_at": "2024-05-12T03:57:07.608693Z",
            "updated_at": "2024-05-12T03:57:07.608693Z",
            "user_id": "66a959f0-efd5-4572-8e80-67a336e6b995",
            "feed_id": "c8a6a42c-a571-441d-b1de-122f0fd6a4c6"
        }
        }

## /v1/all-feeds
This endpoint will fetch all feeds added by all users

    * Method: GET
    * Request Content-Type: empty
    * Response Content-Type: application/json
    * Authorization: empty

    Example Response:
        [
            {
                "id": "c8a6a42c-a571-441d-b1de-122f0fd6a4c6",
                "created_at": "2024-05-12T03:57:07.608693Z",
                "updated_at": "2024-05-12T04:00:59.420246Z",
                "name": "boot.dev blogs",
                "url": "https://blog.boot.dev/index.xml",
                "user_id": "66a959f0-efd5-4572-8e80-67a336e6b995"
            }
        ]


## /v1/feed_follows
This endpoint is for creating feed_follow, whose api_key will be provided that user will start to follow the feed, which id will be provided in request body

    * Method: POST
    * Request Content-Type: application/json
    * Response Content-Type: application/json
    * Authjorization: Bearer api_key

    Example Authorization: Bearer
        ApiKey  dc71925beb7fb83060528594c96b671ca8486ea4be036e49cebfe08c3e7fd498

    Example Request:
        {
            "feed_id": "60e01f39-26dc-4e4c-8276-a679e8263c49"
        }

    Example Response:
        {
            "id": "e5e17051-14c9-4dbd-96db-9ccacef3f4e2",
            "created_at": "2024-05-12T10:08:58.512671Z",
            "updated_at": "2024-05-12T10:08:58.512671Z",
            "user_id": "36723b49-9ae5-4937-8531-335aadd563fb",
            "feed_id": "60e01f39-26dc-4e4c-8276-a679e8263c49"
        }    

## /v1/feed_follows/{feedFollowID}
   This endpoint is for deleting a feed_follow, by providing feed_follow id.

    * Method: DELETE
    * Request Content-Type: empty
    * Response Content-Type: application/json
    * Authorization: empty
    * Query parameter: Require feed_follow id

## /v1/feed_follows
This endpoint is for retrieveing all feed_follows of a user whose api_key is provided

    * Method: GET
    * Request Content-Type: empty
    * Response Content-Type: application/json
    * Authorization: Bearer api_key

    Example Authorization: Bearer
        ApiKey  6f8c1d6fdeb5ab429e2bb537bb8f652018135afdc00edeee803889a6d0d7a3ea

    Example Response:
    [
        {
            "id": "e5e17051-14c9-4dbd-96db-9ccacef3f4e2",
            "created_at": "2024-05-12T10:08:58.512671Z",
            "updated_at": "2024-05-12T10:08:58.512671Z",
            "user_id": "36723b49-9ae5-4937-8531-335aadd563fb",
            "feed_id": "60e01f39-26dc-4e4c-8276-a679e8263c49"
        }
    ] 

## "GET /v1/posts/{limit}
This endpoint will serve posts retrieved from a feed followed by a user whoose api_key is provided

    * Method: GET
    * Request Content-Type: empty
    * Response Content-Type: application/json
    * Authorization: Bearer api_key
    * query parameter:
        If passed limit query parameter then endpoint will serve that number of posts and if no query parameter is given then endpoint will serve only 5 posts


    Example Authorization: Bearer
        ApiKey  6f8c1d6fdeb5ab429e2bb537bb8f652018135afdc00edeee803889a6d0d7a3ea

    Example Response:
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
