create-deploy:
	kubectl create -f deployment.yaml
	kubectl expose deployment my-go-app --type=NodePort --name=my-go-app-svc --target-port=3000