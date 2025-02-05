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
        stage("Sync") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/scripts/automation/github/silkworm") {
                        sh "git pull"
                    }
                }
            }
        }
        stage("Build") {
            steps {
                lock("satis-rebuild-resource") {
                    dir("/data/scripts/automation/github/silkworm") {
                        sh "/data/apps/go/bin/go build -o /data/scripts/automation/programs/silkworm ."
                    }
                }
            }
        }
        stage("Run") {
            steps {
                lock("satis-rebuild-resource") {
                    timeout(time: 5, unit: "MINUTES") {
                        retry(2) {
                            sh "/data/scripts/automation/scripts/run_silkworm.sh"
                        }
                    }
                }
            }
        }
    }
}