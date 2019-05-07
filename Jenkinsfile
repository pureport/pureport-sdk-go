#!/usr/bin/env groovy

@Library('jenkins-build-utils')_

def utils = new com.pureport.Utils()

pipeline {
    agent {
      docker {
        image 'golang:1.12'
      }
    }
    options {
        disableConcurrentBuilds()
    }
    environment {
        GOPATH="/go"
        GOCACHE="/tmp/go/.cache"
        PUREPORT_ENDPOINT="https://dev1-api.pureportdev.com"
        PUREPORT_API_KEY="mKBkM3l1ScUHW"
        PUREPORT_API_SECRET="JMzOfGAbLRcrNziGO"
    }
    stages {
        stage('Build') {
            steps {
                sh "make"
            }
        }
    }
    post {
        success {
            slackSend(color: '#30A452', message: "SUCCESS: <${env.BUILD_URL}|${env.JOB_NAME}#${env.BUILD_NUMBER}>")
        }
        unstable {
            slackSend(color: '#DD9F3D', message: "UNSTABLE: <${env.BUILD_URL}|${env.JOB_NAME}#${env.BUILD_NUMBER}>")

            script {
                utils.sendUnstableEmail()
            }
        }
        failure {
            slackSend(color: '#D41519', message: "FAILED: <${env.BUILD_URL}|${env.JOB_NAME}#${env.BUILD_NUMBER}>")
            script {
                utils.sendFailureEmail()
            }
        }
    }
}
