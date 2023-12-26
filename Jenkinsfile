pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo "Building....."
                sh "docker-compose up --env-file ${ENV_PATH}/.env.kajian-auth build"
                echo "Success build image"
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying...."
                echo "Push to local registry"
                sh "docker-compose up --env-file ${ENV_PATH}/.env.kajian-auth push kajian-auth"
                // sh "ssh -i ${JENKINS_HOME}/light-sail.pem ${LIGHTSAIL_USER}@${LIGHTSAIL_HOST} 'ln -s ${HTML_PATH}/portofolio_build/${BUILD_NUMBER} ${HTML_PATH}/portofolio'"
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
}                                                                                                                                                                                                                                         