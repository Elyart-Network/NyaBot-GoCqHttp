/**
 * JetBrains Space Automation
 * This Kotlin-script file lets you automate build activities
 * For more info, see https://www.jetbrains.com/help/space/automation.html
 */

job("Build & Test") {
    container(displayName = "GoLang", image = "golang:1.19.2") {
        args("export", "GO111MODULE=on")
        args("export", "GOPROXY=https://goproxy.cn,direct")
        args("go", "get")
        args("go", "build")
    }
    startOn {
        gitPush { enabled = true }
    }
}