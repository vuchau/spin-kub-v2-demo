node {
    
      def registry = "vuchauthanh/helloworld"
      def registryCredential = 'dockerhub'
      def dockerImage
      def BRANCH_NAME

    stage('Clone repository') {
          checkout scm
          // Get current branch name
          env.BRANCH_NAME = sh(script: "git name-rev --name-only HEAD | sed -e 's|remotes/origin/||g' | xargs", returnStdout: true)
          sh "env"
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
        dockerImage = docker.build registry
      }
    }
    stage('Deploy Image') {
        script {
          docker.withRegistry('', registryCredential ) {
          dockerImage.push()
          dockerImage.push(env.BRANCH_NAME)          
        }
      }
    }
}
