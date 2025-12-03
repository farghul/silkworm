pipeline {
    agent { label "cactuar && deploy" }
    options {
        buildDiscarder logRotator(
            artifactDaysToKeepStr: "",
            artifactNumToKeepStr: "10",
            daysToKeepStr: "",
            numToKeepStr: "10"
        )
    }
    triggers {
        cron "H 9 * * 3"
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
                dir("/data/automation/checkouts/silkworm"){
                    script {
                        sh "./silkworm.sh"
                    }
                }
            }
        }
    }
}