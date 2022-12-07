run:
	@echo "building..."
	go build -o build/ cmd/*	&& ./build/server
	
test:
	go test -v ./... -coverpkg ./... 

html_coverage:
	go tool cover -html=coverage/cover.out -o coverage/coverage.html

.PHONY: coverage
coverage:
	@mkdir -p coverage
	@go test -coverpkg=`go list ./... | grep -v "mocks"| tr '\n' ','` -coverprofile=coverage/cover.out ./... > /dev/null && go tool cover -func coverage/cover.out
# go tool cover -func coverage/cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}'

docker-build:
	docker build -t famPay/youtube.

k8s-apply:
	kubectl apply -f infra/k8s

k8s-delete:
	kubectl delete -f infra/k8s


