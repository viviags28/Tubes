pipeline {
    agent any

    environment {
        DOCKER_HUB_USER = 'viviags'
        DOCKER_HUB_CRED = 'docker-hub-credentials'
        K8S_NAMESPACE = 'default'
    }

    stages {

        stage('1. Checkout Repo') {
            steps {
                checkout scm
            }
        }

        stage('2. Unit Tests') {
            steps {
                script {

                    catchError(buildResult: 'UNSTABLE', stageResult: 'UNSTABLE') {

                        dir('order-service') {
                            bat 'go test -short ./...'
                        }

                        dir('user-service') {
                            bat 'go test -short ./...'
                        }

                        dir('report-service') {
                            bat 'go test -short ./...'
                        }
                    }
                }
            }
        }

        stage('3. Lint/Vet') {
            steps {
                script {

                    dir('order-service') {
                        bat 'go vet ./...'
                    }

                    dir('user-service') {
                        bat 'go vet ./...'
                    }

                    dir('report-service') {
                        bat 'go vet ./...'
                    }
                }
            }
        }

        stage('4. Build Images (lokal)') {
            steps {
                script {

                    dir('order-service') {
                        bat 'docker build -t %DOCKER_HUB_USER%/order-service:latest .'
                    }

                    dir('user-service') {
                        bat 'docker build -t %DOCKER_HUB_USER%/user-service:latest .'
                    }

                    dir('report-service') {
                        bat 'docker build -t %DOCKER_HUB_USER%/report-service:latest .'
                    }
                }
            }
        }

        stage('5. Functional Tests') {
            steps {
                script {

                    catchError(buildResult: 'UNSTABLE', stageResult: 'UNSTABLE') {

                        dir('order-service') {
                            bat 'docker compose up -d'
                            bat 'go test -tags=functional ./...'
                        }

                        dir('user-service') {
                            bat 'docker compose up -d'
                            bat 'go test -tags=functional ./...'
                        }

                        dir('report-service') {
                            bat 'docker compose up -d'
                            bat 'go test -tags=functional ./...'
                        }
                    }
                }
            }
        }

        stage('6. Push Images') {
            steps {
                script {

                    withDockerRegistry(credentialsId: "${DOCKER_HUB_CRED}", url: 'https://index.docker.io/v1/') {

                        bat 'docker push %DOCKER_HUB_USER%/order-service:latest'
                        bat 'docker push %DOCKER_HUB_USER%/user-service:latest'
                        bat 'docker push %DOCKER_HUB_USER%/report-service:latest'
                    }
                }
            }
        }

        stage('7. Deploy to Kubernetes') {
            steps {
                script {

                    dir('k8s') {

                        bat 'kubectl apply -f order-deployment.yaml'
                        bat 'kubectl apply -f order-service.yaml'

                        bat 'kubectl apply -f user-deployment.yaml'
                        bat 'kubectl apply -f user-service.yaml'

                        bat 'kubectl apply -f report-deployment.yaml'
                    }
                }
            }
        }

        stage('8. Verify') {
            steps {
                script {

                    bat 'kubectl get pods -n %K8S_NAMESPACE%'

                    bat 'kubectl rollout status deployment/report-service --timeout=60s'

                    bat 'kubectl rollout status deployment/order-service --timeout=60s'

                    bat 'kubectl rollout status deployment/user-service --timeout=60s'
                }
            }
        }
    }

    post {

        always {
            echo 'Pipeline selesai.'
        }

        failure {
            echo 'Pipeline gagal di salah satu stage.'
        }
    }
}