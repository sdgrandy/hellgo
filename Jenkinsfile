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
        echo "url is ${API_URL}"
        echo "port is ${API_PORT}"

        if( ['prod', 'qa', 'dev'].contains("${ENVIRONMENT}".toString()) ) {
            withCredentials([
                [
                    $class: 'UsernamePasswordMultiBinding',
                    credentialsId: "creds_id",
                    usernameVariable: 'USERNAME',
                    passwordVariable: 'PASSWORD'
                ],
            ]) {
                stage 'Publish - API'
                //sh "make docker-build"

            }
        }
    }
}
