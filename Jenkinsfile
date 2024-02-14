pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                echo "Building....."
                echo "Success build image"
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying...."
                echo "Push to local registry"
                echo "Preparing ENV"
		sh "cat ${ENV_PATH}/.env.kajian-auth >> ${ENV_PATH}/.env.${GIT_COMMIT}"
                sh "cat ${ENV_PATH}/.env.docker >> ${ENV_PATH}/.env.${GIT_COMMIT}"
		sh "mkdir -p ${ENV_PATH}/kajian_auth"
                sh "rm ${ENV_PATH}/kajian_auth/.env"
                sh "cp ${ENV_PATH}/.env.kajian-auth ${ENV_PATH}/kajian_auth/.env"
                sh "/usr/local/bin/docker-compose  --env-file ${ENV_PATH}/.env.${GIT_COMMIT} up --build -d"
                // sh "ssh -i ${JENKINS_HOME}/light-sail.pem ${LIGHTSAIL_USER}@${LIGHTSAIL_HOST} 'ln -s ${HTML_PATH}/portofolio_build/${BUILD_NUMBER} ${HTML_PATH}/portofolio'"
            }
        }
        stage('Migrations') {
            steps {
                build job: "Kajian/Kajian-auth-migration", wait: false
            }
        }
    }
    post {
        always {
            echo 'One way or another, I have finished'
	    echo "Deleting meta files"
	    sh "rm ${ENV_PATH}/.env.${GIT_COMMIT}"
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
