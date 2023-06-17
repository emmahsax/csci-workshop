# Level 1: Deep in the Crud

For this tutorial, we're going to create Twitter for Zombies. Let's start in the database with a table called `Tweets`. It looks like an Excel spreadsheet with 4 rows and 3 columns. The columns are called ID, status, and zombie.

Now, each row of `Tweets` will be a hash:

```ruby
b = {
  id: 3,
  status: "I just ate some delicious brains",
  zombie: "Jim"
}
```

All of the keys are symbols, meaning they start with a colon.

Once we've saved a hash object to a variable, like above, we can retrieve pieces:

```ruby
b[:zombie]
=> "Jim"
```

There's an alternate way to retrive pieces, too:

```ruby
b.zombie
=> "Jim"
```

It mainly comes down to personal preference which one to use, but this tutorial will mostly use the dot syntax.

So that covers the basics of how to read from a database, but CRUD actually is an acronym for Create, Read, Update, Delete. Let's walk through how to do each of those.

### C -> Create

To make a new tweet, we can use either `.new` and `.create`. For our example, to put a new tweet into `tweets`, we type:

```ruby
Tweet.create()
=> new blank tweet
```

Remember that we type `Tweet` when the hash name is `tweets`. To use the hash, we singularize and capitalize the name of the hash.

There's an alternate way of making a new tweet as well:

```ruby
t = Tweet.new()
=> new blank tweet

t.name = "Emma"
=> "Emma"

t.save
```

The `.create` method automatically saves, so in many ways, this is easier.

And if you know what information you already want, you can put that in directly in hash form:

```ruby
t = Tweet.new(
  status: "I am funny",
  name: "Emma"
)
=> new tweet

t.save
```

We can do this for `.create` as well.

```ruby
Tweet.create(
  status: "I am funny",
  name: "Emma"
)
=> new tweet
```

Note that we don't need to set an ID when we make a new zombie; Rails does that for us automatically.

### R -> Read

To find a tweet, one can use `.find`. For our example, we have a hash database with multiple objects. To find a tweet where the ID is 1, we type:

```ruby
Tweet.find(1)
=> #<Tweet id: 1, status: "I am funny", name: "Emma">
```

If we want to find multiple zombies:

```ruby
Tweet.find(1, 2, 3)
=> array of tweets
```

Let's say we want to find the last tweet in the hash:

```ruby
Tweet.last()
=> #<Tweet id: 1, status: "I am witty", name: "Lemmon">
```

Likewise, there is also a `.first`. A last few useful methods would be: `.count`, `.order`, `.limit`, `.all`, .`where`.

```ruby
Tweet.count
=> total number of tweets

Tweet.order(:name)
=> all tweets

Tweet.limit(2)
=> first 2 tweets

Tweet.where(name: "Emma")
=> all zombies with name Emma
```

We can combine them to do method chaining:

What if we wanted to find all of the zombies, but alphabetized by name:

```ruby
Tweet.all.order(:name)
=> all tweets alphabetized by name
```

### U -> Update

What if we wanted to update a piece of information about a zombie:

```ruby
t = Tweet.find(3)
=> #<Zombie id: 3, status: "I'm only a little crazy", name: "Anonymous">

t.status = "Changing my status"
=> "Changing my status"

t.save
```

We can update multiple attributes at once:

```ruby
t = Tweet.find(3)
=> #<Zombie id: 3, status: "I'm only a little crazy", name: "Anonymous">

t.attributes = {
  status: "Now I'm a lot crazy",
  name: "Elise"
}
=> {status: "Now I'm a lot crazy", name: "Elise"}

t.save
```

Or, we can just use the `.update` method which will automatically save.

```ruby
t = Tweet.find(3)
=> #<Zombie id: 3, status: "I'm only a little crazy", name: "Anonymous">

t.update(
  status: "Now I'm a lot crazy",
  name: "Elise"
)
=> {status: "Now I'm a lot crazy", name: "Elise"}
```

### D -> Delete

To delete a zombie, we use `.destroy` in a few different ways:

```ruby
Tweet.destroy(3)
=> #<Zombie id: 3, status: "I'm only a little crazy", name: "Anonymous">
```

```ruby
t = Tweet.find(3)
=> #<Zombie id: 3, status: "I'm only a little crazy", name: "Anonymous">

t.destroy
```

```ruby
Tweet.find(3).destroy
=> #<Zombie id: 3, status: "I'm only a little crazy", name: "Anonymous">
```

Or to destroy all of the tweets:

```ruby
Tweet.destroy_all
=> DESTROYS ALL TWEETS
```
