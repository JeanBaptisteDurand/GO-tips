# GO-tips
```
a variable declare inside a loop, will be a new instance at each iteration
Il y a des gestionnaires de projet
External function call with a capital letter : math.Pi

go mod init

books to read :

Concurrency in Go: Tools and Techniques for Developers
100 Go Mistakes and How to Avoid Them: Common Go Mistakes


conference de rob pike (google)


Good site to learn syntax :
1. a tour of go
2. go by exemple
3. effective go
4. lets go + lets go further (web dev)
```


## Print
```
fmt.Print: Affiche les arguments sur la sortie standard.
fmt.Println: Comme fmt.Print, mais ajoute un saut de ligne à la fin.
fmt.Printf: Permet de formater les chaînes de caractères comme avec printf en C.
fmt.Sprint: Renvoie une chaîne de caractères au lieu de l'afficher.
fmt.Sprintln: Comme fmt.Sprint, mais ajoute un saut de ligne.
fmt.Sprintf: Comme fmt.Printf, mais renvoie une chaîne de caractères au lieu de l'afficher.
```

## Slice
```
Une tranche a trois composants : un pointeur vers le tableau sous-jacent, une longueur et une capacité.
```

## Ressources
https://www.reddit.com/r/golang/comments/lvwu45/go_and_freelance_work/

## Methode

```
Methods and pointer indirection
Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
For the statement v.Scale(5), even though v is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver.

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
while methods with value receivers take either a value or a pointer as the receiver when they are called:

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
In this case, the method call p.Abs() is interpreted as (*p).Abs().
```
