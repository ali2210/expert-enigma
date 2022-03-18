path "sys/mount/*"{
    capabilities = ["create", "delete", "read", "update", "list"]
}

path "sys/mounts"{
    capabilities = ["read", "list"]
}

path "pki*"{
    capabilities = ["create", "delete", "read", "update", "list", "sudo"]
}