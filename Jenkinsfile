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
    triggers {
        cron "H 9 * * 3"
    }
    stages {
        stage("Git Pull") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/silkworm") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git pull
                        '''
                    }
                }
            }
        }
        stage("Build Silkworm") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/silkworm") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        go build -o /data/automation/bin/silkworm .
                        '''
                    }
                }
            }
        }
        stage("Run") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            sh "/data/automation/scripts/run_silkworm.sh"
                        }
                    }
                }
            }
        }
    }
}