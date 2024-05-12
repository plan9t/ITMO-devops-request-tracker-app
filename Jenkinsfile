pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo "Downloading dependencies and building"
                sh "go mod download"
                sh "go build -o request-tracker-app"
                echo "Successful building request-tracker-app"
            }
        }

        stage('Deployment') {
            steps {
                echo "Check who am i before ssh connection."
                sh "whoami"

                sh "ssh -tt temon01@51.250.86.139 pkill -f request-tracker-app || true"
                echo "Old version app stopped."

                echo "Connecting to devops-server by SSH and execute whoami and pwd commands"
                sh "ssh -tt temon01@51.250.86.139 'whoami; pwd'"
                echo "Successful connection to devops-server"

                echo "Check who am in jenkins server"
                sh "whoami"

                echo "Check pwd in jenkins"
                sh "pwd"

                sh "scp /var/lib/jenkins/workspace/request-tracker-app/request-tracker-app temon01@51.250.86.139:/home/temon01/nats-builded"
                echo "Build successful copied"


                sh "ssh -o BatchMode=yes temon01@51.250.86.139 'whoami; pwd; cd /home/temon01/nats-builded; nohup ./request-tracker-app > request-tracker-app.log 2>&1 & exit;'"
                echo "EOS"
            }
        }
    }
}