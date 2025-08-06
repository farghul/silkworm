pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "28",
            artifactNumToKeepStr: "5",
            daysToKeepStr: "56",
            numToKeepStr: "10"
        )
    }
    stages {
        stage("Empty_Folder") {
            steps {
                dir('/data/automation/checkouts'){
                    script {
                        deleteDir()
                    }
                }
            }
        }
        stage('Checkout_Silkworm'){
            steps{
                dir('/data/automation/checkouts/silkworm'){
                    git url: 'https://github.com/farghul/silkworm.git' , branch: 'main'
                }
            }
        }
        stage('Build_Silkworm') {
            steps {
                dir('/data/automation/checkouts/silkworm'){
                    script {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/silkworm"
                    }
                }
            }
        }
        stage("Checkout_DAC") {
            steps{
                dir('/data/automation/checkouts/dac'){
                    git credentialsId: 'DES-Project', url: 'https://bitbucket.org/bc-gov/desso-automation-conf.git', branch: 'main'
                }
            }
        }
        stage('Run_Silkworm') {
            steps {
                dir('/data/automation/checkouts/dac'){
                    script {
                        sh './scripts/plugin/silkworm.sh'
                    }
                }
            }
        }
    }
}