module github.com/mikhailbolshakov/appA

go 1.16

replace github.com/mikhailbolshakov/appB-proto => ../appB/proto

require github.com/mikhailbolshakov/appB-proto v0.0.0-00010101000000-000000000000
