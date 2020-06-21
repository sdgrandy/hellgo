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
        echo "url is ${API_URL}"
        echo "port is ${API_PORT}"

        if( ['prod', 'qa', 'dev'].contains("${ENVIRONMENT}".toString()) ) {
            withCredentials([
                [
                    $class: 'UsernamePasswordMultiBinding',
                    credentialsId: 'creds_id',
                    usernameVariable: 'API_USERNAME',
                    passwordVariable: 'API_PASSWORD'
                ],
            ]) {
                stage 'RUN'
                //print 'username.collect { it }=' + API_USERNAME.collect { it }
                //print 'password.collect { it }=' + API_PASSWORD.collect { it }
                // sh "API_USERNAME.collect { it } >> "
                //echo "username: ${API_USERNAME}"
                //print ${API_USERNAME}.collect { it }
                API_USERNAME = transform(API_USERNAME)
                API_PASSWORD = transform(API_PASSWORD)
                sh "rm vars.env"
                sh "make vars"
                sh "echo HELLGO_API_USER=${API_USERNAME} >> vars.env"
                sh "echo HELLGO_API_PASSWORD=${API_PASSWORD} >> vars.env"
                sh "cat vars.env"
                sh "make docker-build"
                sh "make docker-up"
            }
        }
    }
}

def transform(s) {
   s = s.collect { it }
   s.replaceAll("[", "");
   s.replaceAll(']','')
   s.replaceAll(',','')
   return s
}
