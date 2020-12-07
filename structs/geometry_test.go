package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 3.0}
	got := Perimeter(rectangle)

	want := 10.0

	if got != want {
		t.Errorf("Got %.2f, Want %.2f\n", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g, Want %g\n", shape, got, want)
		}
	}

	// t.Run("check area of rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{2.0, 3.0}
	// 	want := 6.0

	// 	checkArea(t, &rectangle, want)
	// })
	// t.Run("check area of circle", func(t *testing.T) {
	// 	circle := Circle{2.0}
	// 	want := 12.566370614359172

	// 	checkArea(t, &circle, want)
	// })

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", &Rectangle{2.0, 3.0}, 6.0},
		{"Circle", &Circle{2.0}, 12.566370614359172},
		{"Triangle", &Triangle{2.0, 3.0}, 3.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.hasArea)
		})
	}
}
