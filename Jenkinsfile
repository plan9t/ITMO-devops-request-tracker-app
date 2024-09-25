pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo "Downloading dependencies and building"
                echo "Changing directory to backend"
                dir('backend')
                sh "go mod download"
                sh "go get github.com/nats-io/stan.go"
                sh "go build -o request-tracker-app"
                echo "Successful building request-tracker-app"
            }
        }

        stage('Test') {
            steps {
                echo "Test stage "
                dir('backend')
                sh "go test -v"
            }
        }

        stage('Deployment') {
            steps {
                echo "TEST SSH Connecting  "
                sh "ssh  jenkins@193.176.158.224"

                echo "Check who am i before ssh connection .111"
                sh "whoami"

                sh "ssh -tt jenkins@193.176.158.224 pkill -f request-tracker-app || true"
                echo "Old version app stopped!"

                echo "Connecting to devops-server by SSH and execute whoami and pwd commands111"
                sh "ssh -tt jenkins@193.176.158.224 whoami; pwd"
                echo "Successful connection to devops-server"

                echo "Check who am in main server"
                sh "whoami"

                echo "Check pwd in main server.11"
                sh "pwd"

                sh "scp /var/lib/jenkins/workspace/RequestTracker-DevOps/request-tracker-app jenkins@193.176.158.224:/home/jenkins/go-builded"
                echo "Build successful copied!!"


                sh "ssh -o BatchMode=yes jenkins@193.176.158.224 'whoami; pwd; cd /home/jenkins/request-tracker-app; nohup ./go-builded > request-tracker-app.log 2>&1 & exit;'"
                echo "End script!!!!!"
            }
        }
    }
}