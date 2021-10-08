package CompositeMode

import "fmt"

type folder struct {
	components []component
	name       string
}

func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)

	for _, component := range f.components {
		component.search(keyword)
	}
}

func (f *folder) add(c component) {
	f.components = append(f.components, c)
}
