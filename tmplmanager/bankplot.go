package tmplmanager

import (
	"fmt"
	"math"
	"net/http"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func BankPlot(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Bank plotter")

	// Create a new plot
	p := plot.New()
	// if err != nil {
	// 	panic(err)
	// }

	p.Title.Text = "Sine and Cosine"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	plotutil.AddLinePoints(p,
		"Sine", generatePoints(func(x float64) float64 { return math.Sin(x) }),
		"Cosine", generatePoints(func(x float64) float64 { return math.Cos(x) }),
	)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// })
	// buf := new(bytes.Buffer)

	// w.Header().Set("Content-Type", "image/png")

	buf := ""
	if err := p.Save(4, 4, buf); err != nil {

		fmt.Println("byte save"+err.Error(), http.StatusInternalServerError)
		http.Error(w, "byte save"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("plotter"))
}

func generatePoints(f func(float64) float64) plotter.XYs {
	pts := make(plotter.XYs, 100)
	for i := range pts {
		pts[i].X = float64(i) / 10
		pts[i].Y = f(pts[i].X)
	}
	return pts
}

// func Plot() {
// 	// Set up the route for the plot image
// 	http.HandleFunc("/plot", plotHandler)

// 	// Start the server
// 	fmt.Println("Server started at http://localhost:8080/plot")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
