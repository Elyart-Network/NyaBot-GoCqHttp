/**
 * JetBrains Space Automation
 * This Kotlin-script file lets you automate build activities
 * For more info, see https://www.jetbrains.com/help/space/automation.html
 */

job("Build & Test") {
    container(displayName = "GoLang", image = "golang:1.19.2") {
        env.set("GO111MODULE", "on")
        env.set("GOPROXY", "https://goproxy.cn")
        args("go", "mod download")
        args("go", "build -o nyabot-gocqhttp")
    }
    startOn {
        gitPush { enabled = true }
    }
}