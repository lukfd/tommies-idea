---
theme: default
background: /cover_image.jpg
title: Redis
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

<div class="flex flex-row gap-12">
    <Card class="flex p-12" v-click>
        <template #title>Redis stores data in <b class="text-green-400">memory</b></template>
        <template #content>
            <p class="m-0">
                which makes it much faster than most DBMSs (which persist data to disk).
            </p>
        </template>
    </Card>
    <Card class="flex p-12" v-click>
        <template #title>Redis supports <b class="text-lime-400">writing data to disk</b></template>
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
    <p class="text-lg pb-4">Redis Data Model</p>
    <div class="grid grid-cols-2 gap-4 pb-4">
        <Card v-click>
            <template #title><p class="text-green">User Hash</p></template>
            <template #content>
                <p>
                    uid, username and password
```bash
INCR next_user_id => 1000
HMSET user:1000 username antirez password p1pp0
```
                </p>
            </template>
        </Card>
        <Card v-click>
            <template #title><p class="text-orange">Users Hash</p></template>
            <template #content>
                <p>
                    uid and username
                    <span class="text-xs font-italic text-gray">remember that we are only able to access data in a direct way, without secondary indexes.</span>
```bash
HSET users antirez 1000
```
                </p>
            </template>
        </Card>
    </div>
    <Card v-click>
        <template #title><p class="text-red">Posts Hash</p></template>
        <template #content>
            <p>
                The idea to share
```bash
HMSET post:10343 user_id $owner_id time $time body "This is a great idea!"
```
            </p>
        </template>
    </Card>
</div>

---

<div class="flex">
    <p class="text-2xl">Back End</p>
</div>

---

<div class="flex">
    <p class="text-2xl">Front End</p>
</div>

---

# References

- [https://www.gomomento.com/blog/rip-redis-how-garantia-data-pulled-off-the-biggest-heist-in-open-source-history/](https://www.gomomento.com/blog/rip-redis-how-garantia-data-pulled-off-the-biggest-heist-in-open-source-history/)
- [https://codersee.com/redis-database-explained/](https://codersee.com/redis-database-explained/)
- [https://redis.io/about/](https://redis.io/about/)
- [https://redis.io/docs/latest/get-started/](https://redis.io/docs/latest/get-started/)
