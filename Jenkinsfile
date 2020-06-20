node {
    def projectName = env.JOB_NAME.split("/")[0]
    def url = env.API_URL
    def port = env.API_PORT
    def env = "dev"

    withEnv([
        "PROJECT_NAME=${projectName}",
        "WORKSPACE=${pwd()}",
        "ENVIRONMENT=${env}",
        "API_URL=${url}",
        "API_PORT=${port}"
    ]) {
        stage 'Checkout'
        checkout scm

        stage 'Build'
        //sh "make docker-build"
        //echo "url is ${API_URL}"
        //echo "port is ${API_PORT}"

        if( ['prod', 'qa', 'dev'].contains("${ENVIRONMENT}".toString()) ) {
            withCredentials([
                [
                    $class: 'UsernamePasswordMultiBinding',
                    credentialsId: "creds_id",
                    usernameVariable: 'API_USERNAME',
                    passwordVariable: 'API_PASSWORD'
                ],
            ]) {
                stage 'Publish - API'
                //sh "make print-vars"
                sh "make vars"
                sh "eval '$(./env.sh)' "
                sh "make run"
                sh "rm env.sh"
            }
        }
    }
}
