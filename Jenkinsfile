node {
    def projectName = env.JOB_NAME.split("/")[0]
    def url = env.API_URL
    def port = env.API_PORT
    def env = "dev"
    def envFile = new File("vars.env")

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
                //print 'username.collect { it }=' + API_USERNAME.collect { it }
                //print 'password.collect { it }=' + API_PASSWORD.collect { it }
                // sh "API_USERNAME.collect { it } >> "
                //echo "username: ${API_USERNAME}"
                sh "make vars"
                print ${API_USERNAME}.collect { it }
                envFile.append(${API_USERNAME})
                sh "cat vars.env"
                sh "make docker-build"
                sh "make docker-up"
            }
        }
    }
}
