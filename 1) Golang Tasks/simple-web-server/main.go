package main

import "github.com/gofiber/fiber"

func main() {
	app := fiber.New()

	app.Static("/", "index.html")

	app.Listen(":3000")
	// // Define the handler function
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Serve the static HTML file
	// 	http.ServeFile(w, r, "index.html")
	// })

	// // Start the server
	// fmt.Println("Starting server at :8080")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Error starting server:", err)
	// }

}
