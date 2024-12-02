# gator
A blog aggregator written in Go, with PostgreSQL for the DB

## Get Started

To get started, make sure you: 
1. Install Go 1.23.2 from the [official website](https://go.dev/doc/install). 
2. Install PostgreSQL from [their website](https://www.postgresql.org/download/). *I won't go into detail about how to setup PostgreSQL, so make sure you look up a detailed guide for that.*

Once completed, you can install Gator using the Go toolchain:
```bash
go install github.com/mehkij/gator@latest
```

Once installed, create a config file in your home directory `~/.gatorconfig.json` with your PostgreSQL DB connection string in the content:

```json
{
    "db_url": "postgres://example"
}
```

Just register a user, and you're ready to start using Gator!
```bash
gator register <name>
```

## Commands

There are several commands you can use right out of the box. We've covered how to register a user, but you can also log in with a username provided it is registered in the database:
```bash
gator login <name>
```

You can see the list of available users with the users command:
```bash
gator users
```
```bash
* john
* jane
* pizzalover13 (current)
```

### Feeds

You can add a feed to the database, which will also automatically follow the feed:
```bash
gator addfeed <name> <url>

gator addfeed ExampleFeed https://www.examplefeed.com
```
Maybe a different user added a feed. You can list all feeds registered to the database like so:

```bash
gator feeds

* Name: ExampleFeed, URL: https://www.examplefeed.com, Created by: pizzalover13
* Name: HackerNews, URL: https://news.ycombinator.com/, Created by: jane
```

Looks like there's a feed that you didn't know about that someone else added! You don't seem to be following that feed... let's fix that:

```bash
gator follow https://news.ycombinator.com/
```

You also don't really want to follow ExampleFeed anymore. Let's unfollow it:

```bash
gator unfollow https://www.examplefeed.com
```

### Aggregation

The `agg` command pulls all posts from every feed you follow. It *requires* that you set an interval at which requests to the website are made:

```bash
gator agg <1s|1m|1h>

gator agg 3m
```

Finally, the `browse` command will print a list of posts that you have already fetched using the `agg` command. Note that the `browse` command takes an optional paramater "limit", which sets the amount of posts that it will print. If no number is specified, it will print a maximum of 2 posts:

```bash
gator browse <limit>

gator browse 50
```

### Reset
The `reset` command will delete all users and the feeds they follow from the database. This includes any posts collected using the `agg` command. This is a *non-reversible* change.

```bash
gator reset
```
