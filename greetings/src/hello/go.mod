module example.com/hello

go 1.18

replace example.com/greetings => ../

require example.com/greeting v0.0.0-00010101000000-000000000000

replace example.com/greeting => ../greeting
