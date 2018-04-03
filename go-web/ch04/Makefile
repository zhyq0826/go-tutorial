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