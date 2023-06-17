# Level 4: Controllers Must Be Eaten

The controller uses Models to get stuff out of databases, and uses views to show.

In the previous level, we had something that looked like this:

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

But this is fishy because there is the line `tweet = Tweet.find(1)`. This is fishy because it really shouldn't be in the view, but rather the controller.

When a request first happens, the application will go through the controller, which is located in `app/controllers/tweets_controller.rb`. Then it will go to the view.

If we use conventions, then it is no coincidence that the word `tweet` appears all over the place.

Inside our controller, `app/controllers/tweets_controller.rb`, we have:

```ruby
class TweetsController < ApplicationController
  def show
    @tweet = Tweet.find(1)
    render action: 'status'
  end
end
```

And our html, `app/views/tweets/show.html.erb`, we have:

```html
<h1><%= @tweet.status %></h1>
<p>Posted by <%= @tweet.zombie.name %></p>
```

We add the `@` whenever we want to access the tweet so that it becomes an instance variable.

However, as the code for the controller is, it will only do stuff for the tweet with id 1. So to fix this, we need to do:

```ruby
class TweetsController < ApplicationController
  def show
    @tweet = Tweet.find(params[:id])
    render action: 'status'
  end
end
```

The `params` thing comes with Rails, and it is a hash of all of the necessary IDs.

Along with parameters... parameters are made when urls are used. It's kind of confusing, but all that needs to known is that in enables the going between of instances of tweets, zombies, etc.

Rails 4 requires strong parameters. These are just checks to make sure things don't go awry.

We can use json and xml as well:

```ruby
class TweetsController < ApplicationController
  def show
    @tweet = Tweet.find(params[:id])
    respond_to do |format|
      format.html # show.html.erb
      format.json { render json: @tweet }
      format.xml { render xml: @tweet }
    end
  end
end
```

Then the json section would render out to be:

```json
{"tweet":{"id":1, "status":"Where can I get a good bite to eat?", "zombie_id":1}}
```

And the xml section would render out to be:

```xml
<?xml version="1.0" encoding="UTF-8"?> ...
```

There are some repetitions that will be seen a lot in controllers:

```ruby
class TweetsController < ApplicationController
  def index #list all tweets
  def show #show a single tweet
  def new #show a new tweet form
  def edit #show an edit tweet form
  def create #create a new tweet
  def update #update a tweet
  def destroy #delete a tweet
end
```

Most of these actions would need views associated with them.

When it comes to editing and deleting, we'll want some sort of authentication.

```ruby
class TweetsController < ApplicationController
  def edit
    @tweet = Tweet.find(params[:id])
    if session[:zombie_id] != @tweet.zombie_id
      flash[:notice] = "Sorry, you can't edit this tweet"
      redirect_to(tweets_path)
    end
  end
end
```

The session is the current use of the application, and the flash is a way of sending messages to the user. We redirect so that they can't edit the tweet.

A shorter syntax that combines two of the lines:

```ruby
redirect_to(tweets_path, notice: "Sorry, you can't edit this tweet")
```

In the layout:

```html
<!DOCTYPE html>
<html>
  <head><title>RailsForZombies</title></head>
  <body>
    <header>...</header>
    <% if flash[:notice] %>
      <div id="notice"><%= flash[:notice] %></div>
    <% end %>
    <% yield %>
  </body>
</html>

```

We can stop repetition in the controller by:

```ruby
class TweetsController < ApplicationController
  before_action :get_tweet, only: [:edit, :update, :destroy]
  before_action :check_auth, only => [:edit, :update, :destroy]

  def get_tweet
    @tweet = Tweet.find(params[:id])
  end

  def check_auth
    if session[:zombie_id] != @tweet.zombie_id
      flash[:notice] = "Sorry, you can't edit this tweet"
      redirect_to(tweets_path)
    end
  end

  def edit
    ...
  end

  def update
    ...
  end

  def destroy
    ...
  end
end
```
