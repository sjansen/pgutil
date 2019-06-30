target "demo" "msg1" {
    string = "!abbcF"
    max_concurrency = 3
}

target "demo" "msg2" {
    string = ".rvy n fv rxnp ruG"
    max_concurrency = 3
}

target "demo" "msg3" {
    string = ".revczni n fv qyebj ruG"
    max_concurrency = 3
}

task "demo/echo" "msg1/decrypted" {
    after = ["msg1/reverse", "msg1/rotate"]
    target = "msg1"
}

task "demo/rev" "msg1/reverse" {
    target = "msg1"
}

task "demo/rot13" "msg1/rotate" {
    target = "msg1"
}

task "demo/echo" "msg2/decrypted" {
    after = ["msg2/reverse", "msg2/rotate"]
    target = "msg2"
}

task "demo/rev" "msg2/reverse" {
    target = "msg2"
}

task "demo/rot13" "msg2/rotate" {
    target = "msg2"
}

task "demo/echo" "msg3/decrypted" {
    after = ["msg3/reverse", "msg3/rotate"]
    target = "msg3"
}

task "demo/rev" "msg3/reverse" {
    target = "msg3"
}

task "demo/rot13" "msg3/rotate" {
    target = "msg3"
}
