run_app:
	docker build -t test_app .
	--docker stop test
	--docker rm test
	docker run -p 8080:8080 --name test -d test_app

