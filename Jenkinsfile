node {
    
      def registry = "vuchauthanh/helloworld"
      def registryCredential = 'dockerhub'
      def dockerImage
      def golangVersion = '1.11.4'
    stage('Clone repository') {
          checkout scm
    }
    stage('Unittest') {
        def dockerfile = 'Dockerfile.dev'
        def testImage = docker.build(registry, "-f ${dockerfile} .")
        testImage.inside {
          sh 'go test --cover'
      }
    }
    stage('Building image') {
      script {
        dockerImage = docker.build registry + ":$BUILD_NUMBER"
      }
    }
    stage('Deploy Image') {
        script {
          docker.withRegistry('', registryCredential ) {
          dockerImage.push()
        }
      }
    }
    stage('Remove Unused docker image') {      
        sh "docker rmi $registry:$BUILD_NUMBER"
    }   
}
