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
        stage("Clear") {
            steps {
                dir("/data/automation/checkouts"){
                    script {
                        deleteDir()
                    }
                }
            }
        }
        stage("Checkouts"){
            steps{
                dir("/data/automation/checkouts/silkworm"){
                    git url: "https://github.com/farghul/silkworm.git", branch: "main"
                }
                dir("/data/automation/checkouts/dac"){
                    git credentialsId: "DES-Project", url: "https://bitbucket.org/bc-gov/desso-automation-conf.git", branch: "main"
                }
            }
        }
        stage("Build") {
            steps {
                dir("/data/automation/checkouts/silkworm"){
                    script {
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/silkworm"
                    }
                }
            }
        }
        stage("Run") {
            steps {
                dir("/data/automation/checkouts/dac/scripts/plugin"){
                    script {
                        sh "./silkworm.sh"
                    }
                }
            }
        }
    }
}