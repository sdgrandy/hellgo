node {
    def projectName = env.JOB_NAME.split("/")[0]
    def url = env.URL
    def port = env.PORT
    def env = "dev"

    withEnv([
        "PROJECT_NAME=${projectName}",
        "WORKSPACE=${pwd()}",
        "ENVIRONMENT=${dev}",
        "URL=${url}",
        "PORT=${port}"
    ]) {
        stage 'Checkout'
        checkout scm

        stage 'Build'
        //sh "make docker-build"
        echo "url is ${URL}"

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
