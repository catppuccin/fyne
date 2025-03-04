<h3 align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/logos/exports/1544x1544_circle.png" width="100" alt="Logo"/><br/>
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/misc/transparent.png" height="30" width="0px"/>
	Catppuccin for <a href="https://go.dev">Go</a>
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/misc/transparent.png" height="30" width="0px"/>
</h3>

<p align="center">
	<a href="https://github.com/catppuccin/go/stargazers"><img src="https://img.shields.io/github/stars/mbaklor/fyne-catppuccin?colorA=363a4f&colorB=b7bdf8&style=for-the-badge"></a>
	<a href="https://github.com/catppuccin/go/issues"><img src="https://img.shields.io/github/issues/mbaklor/fyne-catppuccin?colorA=363a4f&colorB=f5a97f&style=for-the-badge"></a>
	<a href="https://github.com/catppuccin/go/contributors"><img src="https://img.shields.io/github/contributors/mbaklor/fyne-catppuccin?colorA=363a4f&colorB=a6da95&style=for-the-badge"></a>
</p>


![fyne_catppuccin](https://github.com/user-attachments/assets/d45a69e1-6194-4398-b295-d6886bfcda0f)

## Usage

1. Run `go get github.com/mbaklor/fyne-catppuccin`
2. Add an import of `github.com/mbaklor/fyne-catppuccin` to your code
3. Call the fyne app's `SetTheme` method with an instantiated theme struct:
   ```go
   func main() {
       a := app.New()

       ctp := catppuccin.New() // this creates a Theme struct with blue as the accent color.
                               //you can instead call NewWithAccent() for a different accent color
       ctp.SetFlavor(catppuccin.Frappe) // if you want to follow the system theme don't set flavor manually.
                                        // by default uses Latte for light and Mocha for dark
       a.Settings().SetTheme(ctp)
       ...
   }
   ```
