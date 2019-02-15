node {
    
      def registry = "vuchauthanh/helloworld"
      def registryCredential = 'dockerhub'
      def dockerImage
      def golangVersion = '1.11.4'
      def BRANCH_NAME = 'dev'

    stage('Clone repository') {
          checkout scm
          // Get current branch name
          //BRANCH_NAME = sh(script: "git name-rev --name-only HEAD | sed -e 's|remotes/origin/|g'", returnStdout: true)
          env.BRANCH_NAME = sh(script: "git name-rev --name-only HEAD | sed -e 's|remotes/origin/||g' | xargs", returnStdout: true)
          sh "echo haha branch ${env.BRANCH_NAME}"
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
        dockerImage = docker.build registry + ":${env.BRANCH_NAME}"
      }
    }
    stage('Deploy Image') {
        script {
          docker.withRegistry('', registryCredential ) {
          dockerImage.push()
        }
      }
    }
}
