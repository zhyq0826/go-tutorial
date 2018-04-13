w01:
	go run web_01.go
body_curl:
	curl -id "first_name=san&last_name=hello" 127.0.0.1:8000/body -v
form_curl:
	curl -id "first_name=san&last_name=hello" 127.0.0.1:8000/form?hello=world
multiform:
	curl -F "myfile=@main.md" -iF "first_name=san&last_name=hello" 127.0.0.1:8000/multiform?hello=world
valueCurl:
	curl -id "first_name=san&first_name=hello" 127.0.0.1:8000/value?hello=world
valueCurl2:
	curl -F "myfile=@main.md" 127.0.0.1:8000/value?hello=world
file_curl:
	curl -F "myfile=@main.md" 127.0.0.1:8000/file
file_curl2:
	curl -F "myfile=@main.md" 127.0.0.1:8000/file2
w02:
	go run web_02.go
write:
	curl 127.0.0.1:8000/write -v
writeHeader:
	curl 127.0.0.1:8000/writeHeader -v
header:
	curl 127.0.0.1:8000/header -v
json:
	curl 127.0.0.1:8000/json -v
w03:
	go run web_03.go
set_cookie: 
	curl 127.0.0.1:8000/set_cookie -v
get_cookie: 
	curl 127.0.0.1:8000/get_cookie -v