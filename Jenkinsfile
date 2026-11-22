pipeline {
    agent any

    environment {
        PATH = "/usr/local/go/bin:${env.PATH}"
    }

    stages {
        stage('Esegui test') {
            steps {
                sh 'go version'
                sh 'go test ./...'
            }
        }
    }
}