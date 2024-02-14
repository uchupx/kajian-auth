pipeline {
    agent any
    stages {
        stage('Preparing') {
            steps {
                echo "Preparing ENV"
                echo "Job name : ${JOB_BASE_NAME}"
                echo "Build number : ${BUILD_DISPLAY_NAME}"
                echo "Git commit : ${GIT_COMMIT}"
                echo "Git branch : ${GIT_BRANCH}"
		        sh "cp ${ENV_PATH}/.env.${JOB_BASE_NAME}  ${TEMP_PATH}/.env.${GIT_COMMIT}"
                sh "cat ${ENV_PATH}/.env.docker | tee -a ${TEMP_PATH}/.env.${GIT_COMMIT}"
		        sh "mkdir -p ${ENV_PATH}/${JOB_BASE_NAME}"
                sh "rm ${ENV_PATH}/${JOB_BASE_NAME}/.env || true"
                sh "cp ${ENV_PATH}/.env.${JOB_BASE_NAME} ${ENV_PATH}/${JOB_BASE_NAME}/.env"
            }
        }
        stage('Deploy') {
            steps {
                echo "Deploying...."
                // echo "Push to local registry"
                sh "/usr/local/bin/docker-compose  --env-file ${TEMP_PATH}/.env.${GIT_COMMIT} up --build -d"
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
            sh "rm ${TEMP_PATH}/.env.${GIT_COMMIT}"
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
