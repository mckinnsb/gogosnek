# gogosnek

A simple snake game in go/gopherjs, done for fun and no profit.

To run, you will need:

* Go
* Ebiten & required libs ( see Ebiten wiki, you will need lib-xorgdev and lib-x11dev for starters )
* GopherJS
* (Optional) Ruby > 2.2.5

I recommend you build with Guard, using ruby.

To install that, simply type ( assuming ruby is installed ):

    bundle install

And to run guard, simply type:

    bundle exec guard

Guard will watch for changes to all your files and run the tests. If the tests pass, it will build the JS and executable,
which will be available both as "gogosnek" and on index.html ( as part of game.js ).
