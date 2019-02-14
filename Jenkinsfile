node {
    
      def registry = "vuchauthanh/helloworld"
      def registryCredential = 'dockerhub'
      def dockerImage
      def golangVersion = '1.11.4'
      def BRANCH_NAME = 'dev'

    stage('Clone repository') {
          checkout scm
          BRANCH_NAME = sh(script: "git name-rev --name-only HEAD | sed -e 's|remotes/origin/||g'", returnStdout: true)
          
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
        dockerImage = docker.build registry + ":$BRANCH_NAME"
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
        sh "docker rmi $registry:$BRANCH_NAME"
    }   
}
