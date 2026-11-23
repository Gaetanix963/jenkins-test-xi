pipeline {
    agent any

    environment {
        PATH = "/usr/local/go/bin:${env.PATH}"
    }

    stages {
        stage('Analisi statica') {
            steps {
                sh 'go vet ./...'
            }
        }

        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }

        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }
    }
}