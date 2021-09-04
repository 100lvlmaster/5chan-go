# 5chan

_The 4chan clone with a one up_

https://fivechan-go.herokuapp.com

![5chanlogo.png](https://5chan.vercel.app/favicon.png)

This repository contains the backend code for the [5chan](https://github.com/100lvlmaster/5chan) project.

It uses [Gofiber](https://github.com/gofiber/fiber).

Installing:

> go install

Debug:

> air

Build:

> go build

Copy the env.example and rename it as .env
and set the env vars as required.

## Fetching data:

Posts:

> https://fivechan-go.herokuapp.com/api/v1/posts

Posts by Id:

> https://fivechan-go.herokuapp.com/api/v1/posts/[id]

Replies:

> https://fivechan-go.herokuapp.com/api/v1/replies

Replies for a post:

> https://fivechan-go.herokuapp.com/api/v1/replies/[postID]
