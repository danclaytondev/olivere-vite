/*
Package vite implements a Vite backend integration for Go as described
in https://vitejs.dev/guide/backend-integration.html.

The integration is done by a HTTP handler that implements http.Handler.
The handler has two modes: Development and production.The handler is
configured by passing a vite.Config struct to the vite.NewHandler function.
The Config struct has four fields:

  - FS: The file system to serve files from. In production, this is the Vite
    output directory, which usually is the "dist" directory. In development,
    this is usually the root directory of the Vite app.
  - PublicFS: The file system to serve public files from. This is usually the
    "public" directory. It is optional and can be nil. If it is nil, we will
    check if the "public" directory exists in the Vite app, and serve files
    from there. If it does not exist, we will not serve any public files. It
    is only used in development mode.
  - IsDev: A boolean that is true if the server is running in development
    mode, false otherwise.
  - ViteURL: The URL of the Vite server, used to load the Vite client in
    development mode (if in Dev mode, we automatically use
    "http://localhost:5173" by default). It is unused in production mode.

Notice that you need to run the Vite server in the background in development
mode, so open up a 2nd console and run something like "npm run dev".

Example:

	// Serve in development mode (assuming your frontend code is in ./frontend,
	// relative to your binary)
	v, err := vite.NewHandler(vite.Config{
		FS:       os.DirFS("./frontend"),
		IsDev:    true,
		PublicFS: os.DirFS("./frontend/public"), // optional: we use the "public" directory under "FS" by default
		ViteURL:  "http://localhost:5173", // optional: we use "http://localhost:5173" by default
	})
	if err != nil { ... }

In production mode, you typically embed the whole generated dist directory
generated by "vite build" into the Go binary, using go:embed. In that case,
your first parameter needs to be the embedded "dist" file system. The second
parameter must be false to enable production mode. The last parameter can be
blank, as it is not used in production mode.

Example:

	//go:embed all:dist
	var distFS embed.FS

	func DistFS() fs.FS {
		// Remove the "dist" prefix
		efs, err := fs.Sub(distFS, "dist")
		if err != nil {
			panic(fmt.Sprintf("unable to serve frontend: %v", err))
		}
		return efs
	}

	// Serve in production mode
	v, err := vite.NewHandler(vite.Config{
		FS:    DistFS(),
		IsDev: false,
	})
	if err != nil { ... }
*/
package vite
