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
        sh "make docker-build"
        echo "url is ${API_URL}"
        echo "port is ${API_PORT}"

        if( ['prod', 'qa', 'dev'].contains("${ENVIRONMENT}".toString()) ) {
            withCredentials([
                [
                    $class: 'UsernamePasswordMultiBinding',
                    credentialsId: "creds_id",
                    usernameVariable: 'API_USERNAME',
                    passwordVariable: 'API_PASSWORD'
                ],
            ]) {
                stage 'RUN'
                print 'username.collect { it }=' + API_USERNAME.collect { it }
                print 'password.collect { it }=' + API_PASSWORD.collect { it }
                sh "username: ${API_USERNAME}"
                //sh "make vars"
                sh "make docker-up"
            }
        }
    }
}
