pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'go build -o main .'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v'
            }
        }
        stage('Docker Build') {
            steps {
                sh 'docker build -t my-go-api .'
            }
        }
        stage('Docker Push') {
            steps {
                withCredentials([string(credentialsId: 'docker_hub_credentials', variable: 'CREDENTIALS')]) {
                    sh "echo ${CREDENTIALS}"
                    sh 'docker login -u ${DOCKER_HUB_USERNAME} -p ${DOCKER_HUB_PASSWORD}'
                    sh 'docker push my-go-api'
                }
            }
        }
        stage('Deploy to K8S') {
            steps {
                sh 'kubectl apply -f k8s-deployment.yaml'
            }
        }
    }
}
