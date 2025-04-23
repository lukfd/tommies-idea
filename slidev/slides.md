---
theme: default
background: /cover_image.jpg
title: Tommies Idea Board
description: A place for Tommies to share ideas
author: Luca Comba
layout: cover
drawings:
  persist: false
transition: slide-left
mdc: true
hideInToc: true
addons:
  - slidev-addon-prime
---

<img src="https://redis.io/wp-content/uploads/2024/04/Logotype.svg" width="50%" />
<br>
Presented by Luca Comba

<div class="absolute left-30px bottom-30px text-neutral-300/50">
  Photo by <a href="https://unsplash.com/@profwicks?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Ben Wicks</a> on <a href="https://unsplash.com/photos/a-close-up-of-a-book-shelf-rKyV2xZDbDg?utm_content=creditCopyText&utm_medium=referral&utm_source=unsplash">Unsplash</a>
</div>

---
title: History
---

<div class="flex flex-row items-baseline">
    <p class="text-2xl pr-3">The History Of Redis</p>
    <!-- <img src="https://redis.io/wp-content/uploads/2024/04/footlogo.svg"/> -->
</div>

<v-switch>
    <template #-1>
        <RedisTimeline v-click :data="[
            {
                date: '2009',
                icon: 'pi pi-star', // Example icon
                title: 'Redis Launched',
                description: 'Created by Salvatore “Antirez” Sanfilippo initially to improve the scalability of his startup LLOOGG. Open-sourced later the same year.'
            },
            {
                date: '2010',
                icon: 'pi pi-briefcase', // Example icon
                title: 'VMware Sponsorship',
                description: 'VMware hired Salvatore Sanfilippo and Pieter Noordhuis to work full-time on Redis development.'
            },
            {
                date: '2011',
                icon: 'pi pi-building', // Example icon
                title: 'Garantia Data Founded',
                description: 'The company that would later become Redis Labs (and then Redis Inc.) was founded by Ofer Bengal and Yiftach Shoolman.'
            },
            {}
        ]" />
    </template>
    <template #0>
        <RedisTimeline v-click :data="[{
                date: '2013',
                icon: 'pi pi-briefcase', // Example icon
                title: 'Pivotal Sponsorship',
                description: 'VMware spun out Pivotal, which took over the sponsorship of Redis development, with Antirez moving to Pivotal.'
            },
            {
                date: '2015',
                icon: 'pi pi-user-plus', // Example icon
                title: 'Antirez Joins Redis Labs',
                description: 'Salvatore Sanfilippo left Pivotal and joined Redis Labs (formerly Garantia Data) as the leader of open source development. Redis Labs became the official sponsor.'
            },
            {
                date: '2016',
                icon: 'pi pi-puzzle', // Example icon
                title: 'Redis Modules Introduced',
                description: 'Redis 4.0 was released, introducing the Redis Modules API, enabling extensions to Redis functionality.'
            },
            {}
        ]" />
    </template>
    <template #1>
        <RedisTimeline :data="[
            {
                date: '2018',
                icon: 'pi pi-sitemap', // Example icon
                title: 'Redis Streams Introduced',
                description: 'Redis 5.0 was released, adding the powerful Redis Streams data type for managing streaming data.'
            },
            {
                date: '2018',
                icon: 'pi pi-file-edit', // Example icon
                title: 'Module License Change',
                description: 'Redis Labs changed the license for certain Redis Modules (not core Redis) from AGPL to a source-available license (initially including Commons Clause, later RSAL).'
            },
            {
                date: '2020',
                icon: 'pi pi-user-minus', // Example icon
                title: 'Antirez Steps Down',
                description: 'Salvatore Sanfilippo stepped down as the lead maintainer of the open-source Redis project.'
            },
            {}
        ]" />
    </template>
    <template #2>
        <RedisTimeline :data="[
            {
                date: '2021',
                icon: 'pi pi-building', // Example icon
                title: 'Redis Labs Rebrands to Redis Inc.',
                description: 'The company officially changed its name from Redis Labs to Redis Inc.'
            },
            {
                date: '2024',
                icon: 'pi pi-file-edit', // Example icon
                title: 'Core License Change & Forks',
                description: 'Redis Inc. changed the license for Redis core from the 3-clause BSD license to dual source-available licenses (RSALv2/SSPLv1). This led to the creation of community forks like Valkey (Linux Foundation) and Redict.'
            }
        ]" />
    </template>
</v-switch>

---

<div class="flex flex-row items-baseline pb-12">
    <!-- <p class="text-2xl pr-3">How Does Redis Work?</p> -->
    <img src="https://redis.io/wp-content/uploads/2024/04/footlogo.svg" class="w-50"/>
</div>

<div class="grid grid-cols-2 gap-4">
    <Card class="flex flex-col justify-between p-12 h-full" v-click>
        <template #title>
            <p class="text-green-400">Redis stores data in <b>memory</b></p>
        </template>
        <template #content>
            <p class="m-0">
                which makes it much faster than most DBMSs (which persist data to disk).
            </p>
        </template>
    </Card>
    <Card class="flex flex-col justify-between p-12 h-full" v-click>
        <template #title>
            <p class="text-lime-400">Redis supports <b>writing data to disk</b></p>
        </template>
        <template #content>
            <ul class="list-disc">
                <li><b>snapshotting</b> by default, every max 2 seconds</li>
                <li><b>append-only file</b> (AOF) logging.</li>
            </ul>
        </template>
    </Card>
</div>

---

<b class="text-2xl">Use Cases</b>

<br class="p-8" />

<div class="grid grid-cols-2">
    <div>
        <b class="text-2xl text-green-400">Pros</b>
        <ul class="pt-8" v-click>
            <li><b class="text-green">Caching</b> store frequently accessed data in memory for fast access.</li>
            <li><b class="text-green">Session Management</b> store user session information, like preferences and login status.</li>
            <li><b class="text-green">Real-time Analytics</b> efficient calculations of real-time metrics.</li>
            <li><b class="text-green">Leaderboards</b> sorted sets for ranking of items.</li>
            <li><b class="text-green">Messaging</b> pub/sub messaging to implement message queues and real-time client-server communication.</li>
            <li><b class="text-green">Product Catalogs</b> it can be a great choice for product catalogs to store product information.</li>
        </ul>
    </div>
    <div>
        <b class="text-2xl text-red">Cons</b>
        <div class="pt-8" v-click>
            <p class="text-2xl">When to <b>not</b> use Redis</p>
            <p><b class="text-red">complex queries</b> and <b class="text-red">relationships</b></p>
        </div>
    </div>
</div>

---

<div class="flex flex-row items-baseline">
    <p class="text-2xl pr-3">Commands</p>
</div>

<div class="p-2" />
<p>Redis is a key-value store. <b class="text-lime-400">Commands</b> are rather <span class="text-green">direct instructions</span> than queries.</p>

<div class="pl-8 pt-2 grid grid-cols-2 gap-4">

<div v-click>
String
```bash
> SET key value

> GET key
value

> DEL key
```
</div>

<div v-click>
List
```bash
> LPUSH mylist a

> LRANGE mylist 0 -1
a
```
</div>

<div v-click>
Set
```bash
> SADD myset a

> SMEMBERS myset
a

> SISMEMBER myset notamember
0
```
</div>

<div v-click>
Hash
```bash
> HMSET myuser name Salvatore surname Sanfilippo country Italy

> HGET myuser surname
Sanfilippo
```
</div>

</div>

<div class="absolute left-15px bottom-15px text-neutral-300/50 text-sm">
  <a href="https://redis.io/docs/latest/commands/">Redis commands documentation</a>
</div>

---

<div class="pb-8">
    <p class="text-2xl font-bold text-purple">The Tommies Idea Board</p>
    <i class="text-md text-teal">A place for Tommies to share idea</i>
</div>

<div class="gap-4">
    <p class="text-lg pb-4" v-click>Redis Data Model</p>
    <div class="grid grid-cols-2 gap-4 pb-4">
        <Card v-click>
            <template #title>
                <p class="text-green-400">Key Value</p>
            </template>
            <template #content>
                <p class="m-0">
                    Just for testing purposes, I have set up a Redis database with a key-value pair.
                    <br />
```go
r.Client.Set(r.Ctx, "foo", "bar", 0)
r.Client.Set(r.Ctx, "key", "value", 0)
```
Which is equivalent to the Redis command:
```bash
> SET foo bar
> SET key value
```
                </p>
            </template>
        </Card>
        <Card v-click>
            <template #title>
                <p class="text-green-400">Idea Hash</p>
            </template>
            <template #content>
                <p class="m-0">
                    The idea hash is a hash that contains all the ideas that are shared on the Tommies Idea Board.
                    <br />
```go
// Create a unique key for the idea
r.Client.Set(r.Ctx, "idea_uid", firstIdea.ID, 0)

// Add the idea to the Redis database
r.Client.HSet(r.Ctx, key, 
"timestamp", timestamp, 
"title", idea.Title, 
"description", idea.Description, 
"writer", idea.Writer, 
"tags", tags)
```
                </p>
            </template>
        </Card>
    </div>
</div>

---

Which is equivalent to the Redis command:

<div v-click>
```bash
> HSET 
 idea_uid 1 
 timestamp 1234567890
 title "My Idea" 
 description "This is my idea"
 writer "Luca Comba" 
 tags "tag1, tag2"
```
</div>

---

<div class="flex">
    <p class="text-2xl">Back End</p>
</div>

<div class="flex items-center">
    <p>Created a server with GoLang</p>
    <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png" alt="Go Logo" class="h-6 ml-2" />
</div>

Redis Client

```go
rdb := redis.NewClient(&redis.Options{
    Addr:     addr,
    Password: pass,
    DB:       db,
})

return &RedisClient{
    Client: rdb,
    Ctx:    context.Background(),
}
```

---

Server

```go
func main() {
	redis := model.NewRedisClient()
	log.Println("Initializing Redis")
	redis.Init()

	// Initialize the default ServeMux
	mux := http.NewServeMux()

	// Define routes
	mux.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("GET /store/{key}", func(w http.ResponseWriter, r *http.Request) {
		routes.Keyval(w, r, redis)
	})
	mux.HandleFunc("GET /idea/{id}", func(w http.ResponseWriter, r *http.Request) {
		routes.GetIdea(w, r, redis)
	})
	
    ...

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

---

Built with Docker
```yaml
version: '3.8'

services:
  go-server:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "8080:8080"
    depends_on:
      - redis
    env_file:
      - .env

  redis:
    image: redis:7.0
    container_name: redis
    ports:
      - "6379:6379"
    command: ["redis-server", "--requirepass", ""]
```

<div class="absolute left-15px bottom-15px text-neutral-300/50 text-sm">
  <a href="https://github.com/lukfd/tommies-idea">Source Code</a>
</div>

---

<div class="flex">
    <p class="text-2xl">Testing the API</p>
</div>

<br>

Fetching the stored value from Redis

```bash
$ curl http://localhost:8080/store/foo
Value from Redis: bar% 
```

<br>

Fetching the idea from Redis

```bash
$ curl http://localhost:8080/idea/1
{"id":1,"timestamp":"1745381047","title":"First Idea","description":"This is my first idea.","writer":"Luca","tags":["idea","redis"]}%
```

<br>

```bash
$ curl http://localhost:8080/ideas
[{"id":1,"timestamp":"1745381047","title":"First Idea","description":"This is my first idea.","writer":"Luca","tags":["idea","redis"]},{"id":2,"timestamp":"1745292170","title":"Second Idea","description":"Another silly description","writer":"Luca","tags":["second","redis"]}]%
```


---

<div class="flex">
    <p class="text-2xl">Front End</p>
</div>

<div class="flex text-center">
    <img src="/frontend.png" class="h-80" />
</div>

<br>

Available at <a href="https://www.tommies.peeperone.com" class="underline text-blue">https://www.tommies.peeperone.com</a>

---

# References

<br>

- [https://www.gomomento.com/blog/rip-redis-how-garantia-data-pulled-off-the-biggest-heist-in-open-source-history/](https://www.gomomento.com/blog/rip-redis-how-garantia-data-pulled-off-the-biggest-heist-in-open-source-history/)
- [https://codersee.com/redis-database-explained/](https://codersee.com/redis-database-explained/)
- [https://redis.io/docs/latest/develop/use/patterns/twitter-clone](https://redis.io/docs/latest/develop/use/patterns/twitter-clone)
- [https://github.com/lukfd/tommies-idea](https://github.com/lukfd/tommies-idea)
