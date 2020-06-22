node {
    def projectName = env.JOB_NAME.split("/")[0]
    def url = "default"
    def port = "default"
    if(env.BRANCH_NAME=="master"){
        url = env.API_URL_MASTER
        port = env.API_PORT_MASTER
    }
    else if(env.BRANCH_NAME=="qa"){
        url = env.API_URL_QA
        port = env.API_PORT_QA
    }
    else if(env.BRANCH_NAME=="dev"){
        url = env.API_URL_DEV
        port = env.API_PORT_DEV
    }
    withEnv([
        "PROJECT_NAME=${projectName}",
        "WORKSPACE=${pwd()}",
        "ENVIRONMENT=${env.BRANCH_NAME}",
        "API_URL=${url}",
        "API_PORT=${port}"
    ]) {
        stage 'Checkout'
        checkout scm
        sh 'printenv'
        echo "environment is ${ENVIRONMENT}"
        echo "url is ${API_URL}"
        echo "port is ${API_PORT}"

        stage 'Build'
        

        if( ['master', 'qa', 'dev'].contains("${ENVIRONMENT}".toString()) ) {
            withCredentials([
                [
                    $class: 'UsernamePasswordMultiBinding',
                    credentialsId: "creds-id-${ENVIRONMENT}",
                    usernameVariable: 'API_USER',
                    passwordVariable: 'API_PASSWORD'
                ],
            ]) {
                stage 'RUN'
                //print 'username.collect { it }=' + API_USERNAME.collect { it }
                //print 'password.collect { it }=' + API_PASSWORD.collect { it }
                // sh "API_USERNAME.collect { it } >> "
                //echo "username: ${API_USERNAME}"
                //print ${API_USERNAME}.collect { it }
                // API_USERNAME = transform(API_USER)
                // API_PASSWORD = transform(API_PASSWORD)
                // sh "rm vars.env"
                // sh "make vars"
                // sh "echo HELLGO_API_USER=${API_USER} >> vars.env"
                // sh "echo HELLGO_API_PASSWORD=${API_PASSWORD} >> vars.env"
                // sh "cat vars.env"
                sh "make print-vars"
                sh "make env-vars"
                sh "make docker-build"
                sh "make docker-up"
            }
        }
    }
}

def transform(s) {
   s = s.collect { it }
   return s
}
