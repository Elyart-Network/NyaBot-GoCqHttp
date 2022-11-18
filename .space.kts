/**
 * JetBrains Space Automation
 * This Kotlin-script file lets you automate build activities
 * For more info, see https://www.jetbrains.com/help/space/automation.html
 */

job("Build & Test") {
    container(displayName = "GoLang", image = "golang:1.19.2") {
        shellScript { // Set Go Proxy
            interpreter = "/bin/bash"
            content = """
                export GO111MODULE=on
                export GOPROXY=https://goproxy.cn
            """
        }
        shellScript { // Download Go Modules
            interpreter = "/bin/bash"
            content = """
                go mod tidy
                go mod download
            """
        }
        shellScript { // Build Project
            interpreter = "/bin/bash"
            content = """
                go build -o nyabot-gocqhttp
            """
        }
    }
    startOn {
        gitPush { enabled = true }
    }
}