pipeline {
    agent any
    environment {
        registry = "vuchauthanh/helloworld"
        registryCredential = 'dockerhub'
        dockerImage = ''
        golangVersion = '1.11.4'
    }
    stages {
        stage('Clone repository') {
        steps {
            checkout scm
        }
    }
    stage('Unittest') {
        steps {
            sh 'docker build -t vuchauthanh/helloworld -f Dockerfile.dev .'
            sh 'docker run vuchauthanh/helloworld go test --cover'
        }
    }
    
    stage('Building image') {
      steps{
        script {
          dockerImage = docker.build registry + ":$BUILD_NUMBER"
        }
      }
    }
    stage('Deploy Image') {
      steps{
         script {
            docker.withRegistry( '', registryCredential ) {
            dockerImage.push()
          }
        }
      }
    }
    stage('Remove Unused docker image') {
      steps{
        sh "docker rmi $registry:$BUILD_NUMBER"
      }
    }   
    }
}
