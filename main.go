package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

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
	Git_Package_Download     uint64
	Docker_Image_Download    uint64
	Package_Import           uint64
}

type LiveProject struct {
	Name       string
	Release_At string
	Source_url string
	Traffic    uint64
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
var cookieExpiration time.Time = time.Now().Add(time.Hour * 12)

const cookieName string = "wisdomenigma"

func main() {

	// handlebar templates
	app := fiber.New(fiber.Config{

		// handlebar views access
		Views: html.New("./views", ".hbs"),

		// app_name
		AppName: "Wisdomenigma.com",

		// private features

		// specialize header for the app
		ServerHeader: "enigma",

		// server methods
		GETOnly: false,

		// route should be exist "/home" & "/home/" both are different
		CaseSensitive: false,
	})

	// static assets attributes
	app.Static("/", "./public", fiber.Static{

		// optimize the assets
		Compress:      true,
		CacheDuration: (time.Second) * 24,
		MaxAge:        24 * 3600,
		Index:         "index.hbs",
	})

	// web hooks connection established
	test_hook := map[string]interface{}{"test": "event"}

	data, err := json.Marshal(test_hook)
	if err != nil {
		fmt.Println("Error creating json object", err.Error())
		return
	}

	err = json.Unmarshal(data, &test_hook)
	if err != nil {
		fmt.Println("Error json encoding", err.Error())
		return
	}

	req, err := http.NewRequest("POST", "https://eolw1b0hoycy0wi.m.pipedream.net", bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error creating request", err.Error())
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error creating response", err.Error())
		return
	}
	defer resp.Body.Close()

	// home page [get] renderer
	app.Get("/", func(c *fiber.Ctx) error {

		// unicode encoding
		c.AcceptsCharsets("utf-8")

		// gzip, deflate, br & compress reduce content space through encoding algorithms.
		// Content space , after compression will accepted by server. There may
		// chance both machines are negotiation on compression algorithm
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		// cookies expire
		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
		}

		// create new cookie
		c.Cookie(&fiber.Cookie{
			Name:     cookieName,
			Value:    strconv.Itoa(rand.Intn(1000)),
			Expires:  time.Now().Add(time.Hour * 12),
			HTTPOnly: true,
			SameSite: "lax",
		})

		// render the template
		return c.Render("index", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// download [get] page renderer
	app.Get("/download", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:     cookieName,
				Value:    strconv.Itoa(rand.Intn(1000)),
				Expires:  time.Now().Add(time.Hour * 12),
				HTTPOnly: true,
				SameSite: "lax",
			})
		}

		return c.Render("download", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// projects [get] information renderer
	app.Get("/projects", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:     cookieName,
				Value:    strconv.Itoa(rand.Intn(1000)),
				Expires:  time.Now().Add(time.Hour * 12),
				HTTPOnly: true,
				SameSite: "lax",
			})
		}

		return c.Render("project", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// sdk [get] route renderer
	app.Get("/sdk", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:     cookieName,
				Value:    strconv.Itoa(rand.Intn(1000)),
				Expires:  time.Now().Add(time.Hour * 12),
				HTTPOnly: true,
				SameSite: "lax",
			})
		}

		return c.Render("sdk", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// api [get] route renderer
	app.Get("/api", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:     cookieName,
				Value:    strconv.Itoa(rand.Intn(1000)),
				Expires:  time.Now().Add(time.Hour * 12),
				HTTPOnly: true,
				SameSite: "lax",
			})
		}

		// modules already build
		api_manager := make([]API_Manager, 3)

		api_manager[0] = API_Manager{
			Package_Name:             "Hashed_block",
			Package_version:          "0.1.1",
			Package_Language:         "Rust",
			Package_Manager_Link_URL: "https://crates.io/crates/hashed_block",
			Git_Package_Download:     0,
			Docker_Image_Download:    0,
			Package_Import:           +55,
		}

		api_manager[1] = API_Manager{
			Package_Manager_Link_URL: "https://crates.io/crates/payPeers",
			Package_Name:             "PayPeers",
			Package_version:          "0.1.1",
			Package_Language:         "Rust",
			Git_Package_Download:     0,
			Docker_Image_Download:    0,
			Package_Import:           +150,
		}

		api_manager[2] = API_Manager{
			Package_Manager_Link_URL: "https://crates.io/crates/probables",
			Package_Language:         "Rust",
			Package_Name:             "Probables",
			Package_version:          "0.1.2",
			Git_Package_Download:     0,
			Docker_Image_Download:    6,
			Package_Import:           +80,
		}

		return c.Render("api", fiber.Map{
			"Title":     "WisdomEnigma",
			"Crate":     api_manager[0].Package_Name,
			"Crate1":    api_manager[1].Package_Name,
			"Crate2":    api_manager[2].Package_Name,
			"Language0": api_manager[0].Package_Language,
			"Language1": api_manager[1].Package_Language,
			"Language2": api_manager[2].Package_Language,
			"Link":      api_manager[0].Package_Manager_Link_URL,
			"Link_1":    api_manager[1].Package_Manager_Link_URL,
			"Link_2":    api_manager[2].Package_Manager_Link_URL,
			"Version":   api_manager[0].Package_version,
			"Version_1": api_manager[1].Package_version,
			"Version_2": api_manager[2].Package_version,
			"Total":     api_manager[0].Git_Package_Download + api_manager[0].Docker_Image_Download + api_manager[0].Package_Import,
			"Total_1":   api_manager[1].Git_Package_Download + api_manager[1].Docker_Image_Download + api_manager[1].Package_Import,
			"Total_2":   api_manager[2].Git_Package_Download + api_manager[2].Docker_Image_Download + api_manager[2].Package_Import,
		})
	})

	// sign [get] route renderer
	app.Get("/sign", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:    cookieName,
				Value:   strconv.Itoa(rand.Intn(1000)),
				Expires: time.Now().Add(time.Hour * 12),
				// HTTPOnly: true,
				SameSite: "lax",
			})
		}

		return c.Render("sign", fiber.Map{
			"Title": "WisdomEnigma",
		})
	})

	// sign [post] route renderer
	app.Post("/sign", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:    cookieName,
				Value:   strconv.Itoa(rand.Intn(1000)),
				Expires: time.Now().Add(time.Hour * 12),
				// HTTPOnly: true,
				SameSite: "lax",
			})
		}

		// get user info
		portal := Login{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		// database reference
		Cloud_credentials = bridge.Firestore_Object()

		// there may be probable that user provided password is less than 7 digits
		if len(portal.Password) < 7 {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Password must be at least 7 characters long",
			})
		}

		// open user portal & validate user credentials
		loginProfile := bridge.GetUserLogin(portal.Email, portal.Password, Cloud_credentials)

		// either user provided credentials not exist then throw exception
		if !reflect.DeepEqual(loginProfile["email"], portal.Email) {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No account on our server... Please signup on wizdwarfs!",
			})
		}

		// or user provided credentials are not different combination
		if reflect.DeepEqual(loginProfile["email"], portal.Email) && !reflect.DeepEqual(loginProfile["password"], portal.Password) {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Try Different combination",
			})
		}

		// Build secure code
		bob, alice = quanta.QuantaCode()

		//channels.Channel_Start(alice[0].BinaryString, "login")

		// create alert message
		alert_message := map[string]interface{}{
			"Title": "User Login",
			"Body":  "Welcome " + reflect.ValueOf(loginProfile["name"]).String() + "!." + " Your reedme code is " + strings.Join(bob[0].BinaryString, ""),
		}

		// engage notification
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

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:    cookieName,
				Value:   strconv.Itoa(rand.Intn(1000)),
				Expires: time.Now().Add(time.Hour * 12),
				// HTTPOnly: true,
				SameSite: "lax",
			})
		}

		// verify code
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

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		// delete alerts
		alerts.SetAlerts(0)

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:    cookieName,
				Value:   strconv.Itoa(rand.Intn(1000)),
				Expires: time.Now().Add(time.Hour * 12),
				// HTTPOnly: true,
				SameSite: "lax",
			})
		}

		return c.Render("userboard", fiber.Map{
			"Title": "WisdomEnigma",
			"Badge": alerts.TAlerts(),
		})
	})

	app.Post("/search", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:    cookieName,
				Value:   strconv.Itoa(rand.Intn(1000)),
				Expires: time.Now().Add(time.Hour * 12),
				// HTTPOnly: true,
				SameSite: "lax",
			})
		}

		// search any route incase route is not provided then throw exception
		searchRoute := c.FormValue("search")
		if len(searchRoute) <= 0 {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "No route specified",
			})
		}

		// create specified route
		routepath := make([]store.RenderRoutes, 6)
		routepath[0] = store.SetRouteParams("index", "/")
		routepath[1] = store.SetRouteParams("download", "/download")
		routepath[2] = store.SetRouteParams("project", "/projects")
		routepath[3] = store.SetRouteParams("sdk", "/sdk")
		routepath[4] = store.SetRouteParams("api", "/api")
		routepath[5] = store.SetRouteParams("sign", "/sign")

		store.SIZE_OF = 6

		store.SetRoute(routepath)

		// search provided route with specified route
		render := store.MapRoute(searchRoute)

		// in case route not exist
		if reflect.DeepEqual(render, store.RenderRoutes{}) {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "DNS lookup Failed",
			})
		}

		// invalid character stream
		if err := render.Validate(); err != nil {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   "Route is invalid",
			})
		}

		// then render page
		return c.Render(render.Route_Handler_Name, fiber.Map{
			"Title": "WisdomEnigma",
		})

	})

	app.Group("userboard/project/traffic", func(c *fiber.Ctx) error {

		c.AcceptsCharsets("utf-8")
		c.AcceptsEncodings("gzip;q=2.0, deflate;q=0.5, *;q=0.8, br;q=1.0, compress;q=1.5,")

		if cookieExpiration == time.Now().Add(time.Hour*12) {
			c.ClearCookie(cookieName)
			c.Cookie(&fiber.Cookie{
				Name:     cookieName,
				Value:    strconv.Itoa(rand.Intn(1000)),
				Expires:  time.Now().Add(time.Hour * 12),
				HTTPOnly: true,
				SameSite: "lax",
			})
		}

		// build projects
		plive := make([]LiveProject, 1)
		plive[0] = LiveProject{
			Name:       "Wizdwarfs",
			Release_At: "Feb,2022",
			Source_url: "https://github.com/ali2210/WizDwarf/releases",
			Traffic:    1300 + 5 + 7,
		}
		return c.Render("ptraffic", fiber.Map{
			"Title":   "WisdomEnigma",
			"Name":    plive[0].Name,
			"Release": plive[0].Release_At,
			"Traffic": plive[0].Traffic,
			"Source":  plive[0].Source_url,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
