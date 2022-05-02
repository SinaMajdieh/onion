package egg

import "fmt"

type pair struct {
	X int
	Y int
}

func Pair(x, y int) pair {
	return pair{x, y}
}

type layer struct {
	Name     string
	Size     pair
	Starting pair
	Addr     string
}

func (l layer) String() string {
	return fmt.Sprintf("name=\"%s\" size=(%d , %d) starting at (%d,%d) Address=\"%s\"", l.Name, l.Size.X, l.Size.Y, l.Starting.X, l.Starting.Y, l.Addr)
}

func New_layer(name string, size, starting pair, addr string) layer {
	return layer{
		Name:     name,
		Size:     size,
		Starting: starting,
		Addr:     addr,
	}
}

var layers []layer

func Add_layer(l *layer) {
	layers = append(layers, *l)
}

func Print_layers() {
	for _, v := range layers {
		fmt.Println(v)
	}
}
