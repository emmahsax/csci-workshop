# Level 5: Routing into Darkness

Routes are the fourth component to the Application Stack.

In order for everything to essentially be evaluated correctly, we need to declare the routes inside the router. This is `config/routes.rb`.

```ruby
ZombieTwitter::Application.routes.draw do
	resources :tweets
end
```

This one line gives us everything that we need to use inside of the `tweets` controller.

When it comes to custom routes, we need to add stuff:

```ruby
ZombieTwitter::Application.routes.draw do
	resources :tweets
	get '/new_tweet' => 'tweets#new'
	get '/all' => 'tweets#index'
end
```

The new line specifies the url that would be used to associate with each of the controller/view pairs.

To add a custom url generator, we need to specify:

```ruby
get '/all' => 'tweets#index', as: 'all_tweets'
```

Then in the view, we can specify:

```ruby
<%= link_to "All Tweets", all_tweets_path %>
```

Now what if we want the url to change?

```ruby
get '/all' => redirect('/tweets')
```

This will redirect from an old url to a new url.

Now, what if we want to redirect the root route. We can do this by simply saying:

```ruby
root to: "tweets#index"
```

And make it a link:

```ruby
<%= link_to "All Tweets", root_path %>
```

Now let's say we want to return all the tweets made in a certain zipcode. We can do this in `/app/controllers/tweets_controller.rb` in the `index` method, since `index.html.erb` is what lists tweets.

```ruby
def index
	if params[:zipcode]
		@tweets = Tweet.where(zipcode: params[:zipcode])
	else
		@tweets = Tweet.all
	end
	respond_to do |format|
		format.html #index.html.erb
		format.xml {render xml: @tweets}
	end
end
```

Now, inside of the routing file:

```ruby
get '/local_tweets/:zipcode' => 'tweets#index'
```

This way, the application knows to store the zipcodes in a parameter with key `zipcode`.

We can name the route too:

```ruby
get '/local_tweets/:zipcode' => 'tweets#index', as: 'local_tweets'
```

Then, the link would be:

```ruby
<%= link_to "tweets in 32828", local_tweets_path(32828) %>
```

If someone goes to a url with a `/zombie_name`, then we have to add:

```ruby
get ':name' => 'tweets#index', as: 'zombie_tweets'
```

And the link is:

```ruby
<%= link_to "Gregg", zombie_tweets_path('greggpollack') %>
```

Then we have to hop over to the `tweets_controller.rb` to add a check to see if a name was sent in:

```ruby
def index
	if params[:name]
		@zombie = Zombie.where(:name params[:name]).first
		@tweets = @zombie.tweets
	else
		@tweets = Tweet.all
	end
end
```
