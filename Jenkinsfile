node {
    
      def registry = "vuchauthanh/helloworld"
      def registryCredential = 'dockerhub'
      def dockerImage
      def golangVersion = '1.11.4'
      def branch = 'dev'
    stage('Clone repository') {
          //checkout scm
          checkoutt[$class: 'LocalBranch', localBranch: "**"] scm
          sh 'printenv'
          //branch = sh(script: 'rev=$(git name-rev --name-only HEAD)', returnStdout: true)
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
        dockerImage = docker.build registry + ":$branch"
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
        sh "docker rmi $registry:$branch"
    }   
}
