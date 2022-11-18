/**
 * JetBrains Space Automation
 * This Kotlin-script file lets you automate build activities
 * For more info, see https://www.jetbrains.com/help/space/automation.html
 */

job("Build & Test") {
    container(displayName = "GoLang", image = "golang:1.19.2") {
        shellScript {
            interpreter = "/bin/bash"
            content = """
                export GO111MODULE=on
                export GOPROXY=https://goproxy.cn
                go mod tidy
                go mod download
                go build -o nyabot-gocqhttp
            """
        }
    }
    startOn {
        gitPush { enabled = true }
    }
}