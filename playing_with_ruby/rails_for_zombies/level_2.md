# Level 2: Models Taste Like Chicken

Models are how we communicate with our data store in Rails.

So in Level 1, we called something that looked like this:

```
t = Tweet.find(3)
```

This works because a model existed. This model is in `app/models/tweet.rb`.

The model might include:

```
class Tweet < ActiveRecord::Base

end
```

An instance of the class `Tweet`, would be found, and then that's what would be saved in our variable.

In the last level, there were a couple times where we made a new tweet, but it didn't have any values. Normally, we don't want that. So we can put a check in:

```
class Tweet < ActiveRecord::Base
	validates_presence_of :status
end
```

So now, when we would make a new tweet without values, `false` would be returned. If we don't know what went wrong, we can type:

```
t = Tweet.new
=> #<Tweet id: nil, status: nil, zombie: nil>

t.save
=> false

t.errors.messages
=> {status:["can't be blank"]}

t.errors[:status][0]
=> "can't be blank"
```

Other validates are:

```
validates_presence_of :status
validates_numericality_of :fingers
validates_uniqueness_of :toothmarks
validates_confirmation_of :password
validates_acceptance_of :zombification
validates_length_of :password, minimum: 3
validates_formate_of :email, with: /regex/i
validates_inclusion_of :age, in: 21, 99
validates_exclusion_of :age, in 0...21,
					   message: "Sorry, you must be over 21"
```
 
We can write this differently:

```
validates :status, 
          presence: true
          length {minimum: 3}
```

Now, what if we had a `zombies` table as well as a `tweets` table. Instead of `name` in `tweets`, can have a `zombie_id`, which can correspond to the IDs in the `zombies` table.

Now we have to somehow map the two models together to tell our application that there's a relationship.

Let's go into the `Zombie` model and tell the application that a zombie can have multiple tweets:

```
class Zombie < ActiveRecord::Base
	has_many :tweets
end
```

Notice that `:tweets` is plural because a zombie can have multiple tweets.

To tell the application that a single tweet belongs to a zombie, we go into the `Tweet` model:

```
class Tweet < ActiveRecord::Base
	belongs_to :zombie
end
```

Notice that `:zombie` is singular because a tweet can only belong to one zombie.

Let's use some of this:

```
ash = Zombie.find(1)
=> #<Zombie id: 1, name: "Ash", graveyard: "Glen Haven Memorial Cemetery">

t = Tweet.create(status: "Your eyelids taste like bacon.",
                 zombie: ash)
=> #<Tweet id: 5, status: "Your eyelids taste like bacon.", zombie_id: 1>

ash.tweets.count
=> 3

ash.tweets
=> [#<Tweet id: 1, status: "Where can I get a good bite to eat?", zombie_id: 1>,
    #<Tweet id: 4, status: "OMG, my fingers turned green? #FML", zombie_id: 1>,
    #<Tweet id: 5, status: "Your eyelids taste like bacon.", zombie_id: 1>]
```

Let's go the other way and say we want to find information about a zombie from a tweet:

```
t = Tweet.find(5)
=> #<Tweet id: 5, status: "Your eyelids taste like bacon.", zombie_id: 1>

t.zombie
=> #<Zombie id: 1, name: "Ash", graveyard: "Glen Haven Memorial Cemetery">

t.zombie.name
=> "Ash"
```
