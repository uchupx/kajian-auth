pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo "Building....."
                sh "/usr/local/bin/docker-compose  --env-file ${ENV_PATH}/.env.kajian-auth build"
                echo "Success build image"
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying...."
                echo "Push to local registry"
                sh "/usr/local/bin/docker-compose  --env-file ${ENV_PATH}/.env.kajian-auth push kajian-auth"
                echo "Creating symlink for env"
                sh "mkdir -p ${ENV_PATH}/kajian_auth"
                sh "rm ${ENV_PATH}/kajian_auth/.env"
                sh "ln -s ${ENV_PATH}/.env.kajian-auth ${ENV_PATH}/kajian_auth/.env"
                sh "/usr/local/bin/docker-compose  --env-file ${ENV_PATH}/.env.kajian-auth up --build -d"
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