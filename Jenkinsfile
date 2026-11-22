pipeline {
    agent any

    stages {
        stage('Scarica codice') {
            steps {
                checkout scm
            }
        }

        stage('Esegui test') {
            steps {
                sh 'go version'
                sh 'go test ./...'
            }
        }
    }
}