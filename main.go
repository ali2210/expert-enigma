package main

import (
	"log"
	"reflect"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/ali2210/expert-enigma/alerts"
	"github.com/ali2210/expert-enigma/bridge"
	"github.com/ali2210/expert-enigma/quanta"
	"github.com/ali2210/expert-enigma/store"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/itsubaki/q/pkg/quantum/qubit"
)

type API_Manager struct {
	Package_Name             string
	Package_version          string
	Package_Language         string
	Package_Manager_Link_URL string
}

type Login struct {
	Email    string
	Password string
}

type Verify_Code struct {
	i0 string
	i1 string
	i2 string
}

var Cloud_credentials *firestore.Client = nil
var bob []qubit.State
var alice []qubit.State

func main() {

	// handlebar templates
	app := fiber.New(fiber.Config{

		// handlebar views access
		Views: html.New("./views", ".hbs"),
	})

	app.Static("/", "./public")

	// home page [get] renderer
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// download [get] page renderer
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Render("download", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// projects [get] information renderer
	app.Get("/projects", func(c *fiber.Ctx) error {
		return c.Render("project", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// sdk [get] route renderer
	app.Get("/sdk", func(c *fiber.Ctx) error {
		return c.Render("sdk", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// api [get] route renderer
	app.Get("/api", func(c *fiber.Ctx) error {

		// api reference
		api_manager := make([]API_Manager, 3)

		api_manager[0] = API_Manager{
			Package_Name:             "hashed_block",
			Package_version:          "0.1.1",
			Package_Language:         "rust",
			Package_Manager_Link_URL: "https://crates.io/crates/hashed_block",
		}

		api_manager[1] = API_Manager{
			Package_Manager_Link_URL: "https://crates.io/crates/payPeers",
			Package_Name:             "payPeers",
			Package_version:          "0.1.1",
			Package_Language:         "rust",
		}

		api_manager[2] = API_Manager{
			Package_Manager_Link_URL: "https://crates.io/crates/probables",
			Package_Language:         "rust",
			Package_Name:             "probables",
			Package_version:          "0.1.2",
		}

		return c.Render("api", fiber.Map{
			"Title":  "WisdomEnigma",
			"Crate":  api_manager[0],
			"Crate1": api_manager[1],
			"Crate2": api_manager[2],
		})
	})

	// sign [get] route renderer
	app.Get("/sign", func(c *fiber.Ctx) error {

		return c.Render("sign", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// sign [post] route renderer
	app.Post("/sign", func(c *fiber.Ctx) error {

		portal := Login{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		Cloud_credentials = bridge.Firestore_Object()

		if len(portal.Password) < 7 {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Password must be at least 7 characters long",
			})
		}

		loginProfile := bridge.GetUserLogin(portal.Email, portal.Password, Cloud_credentials)

		if !reflect.DeepEqual(loginProfile["email"], portal.Email) {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No account on our server... Please signup on wizdwarfs!",
			})
		}

		if reflect.DeepEqual(loginProfile["email"], portal.Email) && !reflect.DeepEqual(loginProfile["password"], portal.Password) {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Try Different combination",
			})
		}

		bob, alice = quanta.QuantaCode()

		//channels.Channel_Start(alice[0].BinaryString, "login")

		alert_message := map[string]interface{}{
			"Title": "User Login",
			"Body":  "Welcome " + reflect.ValueOf(loginProfile["name"]).String() + "!." + " Your reedme code is " + strings.Join(bob[0].BinaryString, ""),
		}

		if ok := alerts.Watchpoint(alert_message, "login"); !ok {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Internal Error report",
			})
		}

		return c.Render("dashboard", fiber.Map{
			"Title": "WisdomEnigma",
		})

	})

	app.Post("/dashboard", func(c *fiber.Ctx) error {

		verify := Verify_Code{
			i0: c.FormValue("qubit0"),
			i1: c.FormValue("qubit1"),
			i2: c.FormValue("qubit2"),
		}

		if !reflect.DeepEqual(verify.i0, "") && !reflect.DeepEqual(verify.i1, "") && !reflect.DeepEqual(verify.i2, "") {

			if ok := quanta.QuantaValid(bob[0].BinaryString, alice[0].BinaryString); !ok {
				return c.Status(404).JSON(&fiber.Map{
					"success": false,
					"error":   "Internal Error report",
				})
			}

		}

		return c.Render("userboard", fiber.Map{
			"Title": "WisdomEnigma",
			"Badge": alerts.TAlerts(),
		})
	})

	app.Post("/userboard", func(c *fiber.Ctx) error {

		alerts.SetAlerts(0)
		return c.Render("userboard", fiber.Map{
			"Title": "WisdomEnigma",
			"Badge": alerts.TAlerts(),
		})
	})

	app.Post("/search", func(c *fiber.Ctx) error {
		searchRoute := c.FormValue("search")
		if len(searchRoute) <= 0 {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No route specified",
			})
		}

		routepath := make([]store.RenderRoutes, 6)
		routepath[0] = store.SetRouteParams("index", "/")
		routepath[1] = store.SetRouteParams("download", "/download")
		routepath[2] = store.SetRouteParams("project", "/projects")
		routepath[3] = store.SetRouteParams("sdk", "/sdk")
		routepath[4] = store.SetRouteParams("api", "/api")
		routepath[5] = store.SetRouteParams("sign", "/sign")

		store.SIZE_OF = 6

		store.SetRoute(routepath)
		render := store.MapRoute(searchRoute)

		if reflect.DeepEqual(render, store.RenderRoutes{}) {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "DNS lookup Failed",
			})
		}

		if err := render.Validate(); err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Route is invalid",
			})
		}

		return c.Render(render.Route_Handler_Name, fiber.Map{
			"Title": "WisdomEnigma",
		})

	})

	log.Fatal(app.Listen(":3000"))
}
