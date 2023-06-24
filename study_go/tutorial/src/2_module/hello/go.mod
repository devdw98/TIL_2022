module tutorial/2_module/hello

go 1.20

replace tutorial/2_module/greetings => ../greetings

require tutorial/2_module/greetings v0.0.0-00010101000000-000000000000
