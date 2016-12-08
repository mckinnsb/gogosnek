guard :build do

  listener = Listen.to ".", :ignore => /\.js/ do | *args |

    system "clear" or system "cls"

    test = `go test`
    fail = test.match "FAIL"

    puts "RUNNING GUARD AT : #{Time.now}"

    unless fail
      puts "**TESTS PASS**"
      result = `go test && go install && gopherjs build -o game.js` && "**GAMEBUILT**"
      puts result
    else
      puts "**TESTS FAIL**"
      puts test
    end

  end

  listener.start

  sleep

end

