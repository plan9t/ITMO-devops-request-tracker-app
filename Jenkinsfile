pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo "Downloading dependencies and building"
                sh "go mod download"
                sh "go get github.com/nats-io/stan.go"
                sh "go build -o request-tracker-app"
                echo "Successful building request-tracker-app"
            }
        }

        stage('Deployment') {
            steps {
                echo "Check who am i before ssh connection ."
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

                sh "scp /var/lib/jenkins/workspace/request-tracker-app/request-tracker-app root@193.176.158.224:/home/temon01/nats-builded"
                echo "Build successful copied!!"


                sh "ssh -o BatchMode=yes root@193.176.158.224 'whoami; pwd; cd /home/temon01/nats-builded; nohup ./request-tracker-app > request-tracker-app.log 2>&1 & exit;'"
                echo "End script!!!"
            }
        }
    }
}