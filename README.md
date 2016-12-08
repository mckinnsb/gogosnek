# gogosnek

A simple snake game in go/gopherjs, done for fun and no profit.

To run, you will need:

1) Go
2) Ebiten & required libs ( see Ebiten wiki, you will need lib-xorgdev and lib-x11dev for starters )
2) GopherJS
4) (Optional) Ruby > 2.2.5

I recommend you build with Guard, using ruby.

To install that, simply

    bundle install

And to run:
    bundle exec guard

Guard will watch for changes to all your files and run the tests. If the tests pass, it will build the JS and executable,
which will be available both as "gogosnek" and on index.html ( as part of game.js ).
