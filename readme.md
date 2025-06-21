## QP Flutter CLI

<p>This is just a CLI that clones my flutter template for me to use when trying to build any flutter app. I was struggling with always trying to scaffold a flutter project setting up all I need in the project which mostly eats up my productive time</p>

<p>So I decided to build this CLI using Go Lang to always clone the repo that I have already setup and made public. It clones it and run the flutter pub get command to make sure the project is just ready for me to start building apps or anything with flutter instead of trying to setup my folder structure.</p>

<h4>To run this if you want to use this code for your flutter template:</h4>

## Simple steps
1. make build -> "this will build the binary"
2. make move -> "this will require sudo privileges"
3. qp_flutter create test_app -p com.example.app -> "this will create a flutter app with my template"

## Build Issues?
1. make build-clean -> "this will clean the build binary"

<p>You can change qp_flutter to what you want, the second step is for moving the build into your local bin folder, this is for Mac OS as of the moment</p>

<p>The last step there is for creating the flutter app with my template that uses Bloc and Cubit as State Management. You will be given a Counter App page that uses Cubit and not SetState</p>