# Level 3: The Views Ain't Always Pretty

The view is the user interface for applications.

In our application, we have an `app` folder. In that, we'll have a `views` folder. In the `view` folder, we'll put a `zombies` folder and a `tweets` folder.

Within each folder in `views`, we'll have an `index.html.erb` and a `show.html.erb`. `erb` stands for embedded ruby.

Inside `app/views/tweets/show.html.erb`:

```html
<!DOCTYPE html>
<html>
	<head><title>RailsForZombies</title></head>
	<body>
		<header>...</header>
		<% tweet = Tweet.find(1) %>
		<h1><%= tweet.status %></h1>
		<p>Posted by <%= tweet.zombie.name %></p>
	</body>
</html>
```

The `<%...%>` is a way of saying this is going to Ruby code, so execute it. When we follow it with an `=` sign tells Rails that whatever the code returns should be printed into the html.

Now, if we wanted to use this basic header information again, it isn't very D.R.Y., or don't repeat yourself. Let's move it to `app/views/layouts/application.html.erb`.

```html
<!DOCTYPE html>
<html>
	<head><title>RailsForZombies</title></head>
	<body>
		<header>...</header>
		<%= yield %>
	</body>
</html>
```

The yield command means this where the contents of the `show.html.erb` would go. This file would end up looking like this:

```html
<% tweet = Tweet.find(1) %>
<h1><%= tweet.status %></h1>
<p>Posted by <%= tweet.zombie.name %></p>
```

What if we wanted to create a link that would show the zombie's profile page?

```html
<%= link_to tweet.zombie.name, zombie_path(tweet.zombie) %>

<!-- OR -->

<%= link_to tweet.zombie.name, tweet.zombie %>
```

The recipe for that is:

```html
<%= link_to text_to_show, model_instance %>
```

As with tweets, shorter is better. Now we move the link up into the code:

```html
<% tweet = Tweet.find(1) %>
<h1><%= tweet.status %></h1>
<p>Posted by <%= link_to tweet.zombie.name, tweet.zombie %></p>
```

`link_to` is a special Rails command, meaning it has its own API and helper pieces to go with it. A confirmation message is one of these:

```html
<% tweet = Tweet.find(1) %>
<h1><%= tweet.status %></h1>
<p>Posted by <%= link_to tweet.zombie.name
					     tweet.zombie,
					     confirm: "Are you sure?" %></p>
```

Remember how the template for the entire application is in `app/views/layouts/application.html.erb`? But what about the `app/views/tweets/index.html.erb`? This is where we'll list all of the tweets.

```html
<h1>Listing tweets</h1>
<table>
	<tr>
		<th>Status</th>
		<th>Zombie</th>
	</tr>
<% Tweet.all.each do |tweet| %>
	<tr>
		<td><%= tweet.status %></td>
		<td><%= tweet.zombie.name %></td>
	</tr>
<% end %>
</table>
```

The loop is saying to go through all of the tweets, and then print the status and the name. Now what if we wanted everything in the table to be links:

```html
<% Tweet.all.each do |tweet| %>
	<tr>
		<td><%= link_to tweet.status, tweet%></td>
		<td><%= link_to tweet.zombie.name, tweet.zombie %></td>
	</tr>
<% end %>
```

What if we pull up the page and there aren't any tweets yet? We want it to actually tell us that there aren't any tweets. So we need an `if`:

```html
<% tweets = Tweet.all %>
<% tweets.each do |tweet| %>
	<tr>
		<td><%= link_to tweet.status, tweet%></td>
		<td><%= link_to tweet.zombie.name, tweet.zombie %></td>
	</tr>
<% end %>
<% if tweets.size == 0 %>
	<em>No Tweets found</em>
<% end %>
```

What if a zombie wants to edit or delete a tweet?

```html
<% tweets.each do |tweet| %>
	<tr>
		<td><%= link_to tweet.status, tweet %></td>
		<td><%= link_to tweet.zombie.name, tweet.zombie %></td>
		<td><%= link_to "Edit", edit_tweet_path(tweet) %></td>
		<td><%= link_to "Destroy", tweet, method: :delete %></td>
	</tr>
<% end %>
```

There are several different URL generator methods that Rails supports:

| Action          | Code                     | The URL        |
|:----------------|:-------------------------|:---------------|
| List all tweets | `tweets_path`            | /tweets        |
| New tweet form  | `new_tweet_path`         | /tweets/new    |
| Show a tweet    | `tweet`                  | /tweets/1      |
| Edit a tweet    | `edit_tweet_path(tweet)` | /tweets/1/edit |
| Delete a tweet  | `tweet, method: :delete` | /tweets/1      |

The code column coes in the `link_to` as the second argument.
