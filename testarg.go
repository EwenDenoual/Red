package main

type charcter struct {
	name string
	lvl int 
	pv int
}

func main() {
	pedro := charcter{name: "pedro", lvl: 1, pv: 20}

	println(pedro)
}