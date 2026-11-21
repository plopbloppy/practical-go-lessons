package main

import "thisIsATest2/room"

func main() {
   room.PrintDetails(112, 3, 2)
}

/*
// holds all initialization tasks necessary for the program to run correctly
func init() {
    fmt.Println("launch initialization")
}

// executed after all initialization tasks have been performed
// in this fcn, you usually call other pkgs and implement program logic
func main() {
    fmt.Println("launch the program !")
}

// Output:
// launch initialization <- init fcn is launched, then the main fcn
// launch the program!

6 The main package
Notes:
- it is not an obligation to name the file that holds main pkg main.go
(in fact it is much more helpful to use a more relevant name instead)
- main pkgs typically live in a cmd/ folder at the root of the repo

7 The go.mod file
Ex of module:
module kaili-tang.com/myProg

go 1.27.4

- Defn: a module is a collection of Go pkgs stored in a file tree with a go.mod file at its root
- when adding a pkg main to a module, it is made into a program that can be compiled and executed
- the module path for the module is the unique location of the module

Ex of module paths:
- Hashicorp Vault
module github.com/hashicorp/vault (link to repo, access can be public/private)

- Kubernetes
module k8s.io/kubernetes (only accessible to authorized devs)

Note: the program will not be automatically shared or exposed by Go if a remote path is chosen
*/