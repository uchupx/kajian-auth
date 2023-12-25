pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
            }
        }
        stage('Deploy') {
            steps {
                echo "Creating symlink"
                sh "ssh -i ${JENKINS_HOME}/light-sail.pem ${LIGHTSAIL_USER}@${LIGHTSAIL_HOST} 'ln -s ${HTML_PATH}/portofolio_build/${BUILD_NUMBER} ${HTML_PATH}/portofolio'"
            }
        }
    }
    post {
        always {
            echo 'One way or another, I have finished'
            deleteDir() /* clean up our workspace */
        }
        success {
            echo 'I succeeeded! bisa'
        }
        unstable {
            echo 'I am unstable :/'
        }
        failure {
            echo 'I failed :('
        }
        changed {
            echo 'Things were different before...'
        }
    }
}%                                                                                                                                                                                                                                               