clear-deploy:
	kubectl delete svc my-go-app-svc
	kubectl delete deploy my-go-app

create-deploy:
	kubectl create -f deployment.yaml
	kubectl expose deployment my-go-app --type=NodePort --name=my-go-app-svc --target-port=3000

run-all:
	make clear-deploy
	make create-deploy