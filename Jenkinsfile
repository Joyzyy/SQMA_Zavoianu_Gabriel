pipeline {
  agent any

  tools {
    go '1.25.1'
  }

  parameters {
    choice(
      name: 'SUITE',
      choices: ['all', 'addition', 'subtraction'],
      description: 'Which test suite to run'
    )
    string(
      name: 'GO_TEST_FLAGS',
      defaultValue: '-v -count=1',
      description: 'Extra flags passed to `go test`'
    )
  }

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Test') {
      steps {
        dir('go_tests') {
          sh 'go version'

          script {
            def runPattern = ''
            switch (params.SUITE) {
              case 'addition':
                runPattern = '^TestAddition$'
                break
              case 'subtraction':
                runPattern = '^TestSubtraction$'
                break
              default:
                runPattern = '.'
            }

            sh "go test ${params.GO_TEST_FLAGS} -run '${runPattern}'"
          }
        }
      }
    }
  }
}
