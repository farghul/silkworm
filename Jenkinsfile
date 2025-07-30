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
        stage("Pull Config Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/bitbucket/desso-automation-conf") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git switch main
                        git pull
                        '''
                    }
                }
            }
        }
        stage("Pull Program Changes") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/automation/github/silkworm") {
                        sh '''#!/bin/bash
                        source ~/.bashrc
                        git fetch --all
                        git checkout main
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
                        sh "/data/apps/go/bin/go build -o /data/automation/bin/silkworm"
                    }
                }
            }
        }
        stage("Run Silkworm") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            dir("/data/automation/bitbucket/desso-automation-conf/scripts/plugin") {
                                sh "./silkworm.sh"
                            }
                        }
                    }
                }
            }
        }
    }
}