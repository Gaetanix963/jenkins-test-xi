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
                sh 'go build -o app-jenkins .'
            }
        }
    }
    post {
        success {
            archiveArtifacts artifacts: 'app-jenkins', fingerprint: true
            echo "Pipeline OK! Binario 'app-jenkins' salvato tra gli artefatti."
        }
        failure {
            echo "Pipeline FALLITA! Controlla il Console Output per vedere quale stage ha dato errore."
        }
        always {
            echo "Esecuzione terminata per il commit: ${env.GIT_COMMIT ?: 'n/d'}"
        }
    }
}