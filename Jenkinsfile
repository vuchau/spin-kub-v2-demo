node {
    
      def registry = "vuchauthanh/helloworld"
      def registryCredential = 'dockerhub'
      def dockerImage
      def BRANCH_NAME
      def imageTag
      def buildProp

    stage('Clone repository') {
          
          checkout scm
          /* Get current branch name
          Remove prefix and newline
          */
          BRANCH_NAME = sh(script: "git name-rev --name-only HEAD | sed -e 's|remotes/origin/||g' | tr -d '\n'", returnStdout: true)          
          env.BRANCH_NAME = BRANCH_NAME
          env.IMAGE_TAG = "$BRANCH_NAME.${env.BUILD_NUMBER}"
          buildProp = readProperties file: 'build.properties'
          echo """The version name ${buildProp['version']}"""
    }
    stage('Unittest') {
        // Build images for run unittest
        def dockerfile = 'Dockerfile.dev'
        def testImage = docker.build(registry, "-f ${dockerfile} .")
        testImage.inside {
          sh 'go test --cover'
      }
    }
    stage('Build image') {
      // Build images for deploy
      script {
        dockerImage = docker.build registry
      }
    }
    stage('Deploy Image') {
        script {
          docker.withRegistry('', registryCredential ) {
            /* Finally, we'll push the image with two tags:
            * First, the incremental build number from Jenkins
            * Second, the 'latest' tag.
            * Pushing multiple tags is cheap, as all the layers are reused. */
            dockerImage.push(${buildProp['version']})          
            dockerImage.push("latest")
        }
      }
      archiveArtifacts artifacts: 'build.properties', onlyIfSuccessful: true
    }
}
