pipeline {
    agent any

    environment {
        DOCKER_HUB_USER = 'viviags'
    }

    stages {

        stage('1. Checkout Repo') {
            steps {
                checkout scm
            }
        }

        stage('2. Unit Tests') {
            steps {
                dir('report-service') {
                    bat 'go test -short ./...'
                }
            }
        }

        stage('3. Lint/Vet') {
            steps {
                dir('report-service') {
                    bat 'go vet ./...'
                }
            }
        }

        stage('4. Build Images (lokal)') {
            steps {
                dir('report-service') {
                    bat 'docker build -t %DOCKER_HUB_USER%/report-service:latest .'
                }
            }
        }

        stage('5. Functional Tests') {
            steps {
                dir('report-service') {
                    bat 'set DB_HOST=localhost && go test -v -tags=functional ./...'
                }
            }
        }
    }

    post {
        always {
            echo 'Pipeline selesai.'
        }

        failure {
            echo 'Pipeline gagal.'
        }
    }
}