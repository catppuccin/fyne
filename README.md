<h3 align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/logos/exports/1544x1544_circle.png" width="100" alt="Logo"/><br/>
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/misc/transparent.png" height="30" width="0px"/>
	Catppuccin for <a href="https://fyne.io/">Fyne</a>
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/misc/transparent.png" height="30" width="0px"/>
</h3>

<p align="center">
	<a href="https://github.com/catppuccin/fyne/stargazers"><img src="https://img.shields.io/github/stars/catppuccin/fyne?colorA=363a4f&colorB=b7bdf8&style=for-the-badge"></a>
	<a href="https://github.com/catppuccin/fyne/issues"><img src="https://img.shields.io/github/issues/catppuccin/fyne?colorA=363a4f&colorB=f5a97f&style=for-the-badge"></a>
	<a href="https://github.com/catppuccin/fyne/contributors"><img src="https://img.shields.io/github/contributors/catppuccin/fyne?colorA=363a4f&colorB=a6da95&style=for-the-badge"></a>
</p>

<p align="center">
	<img src="assets/preview.webp"/>
</p>

## Previews

<details>
<summary>🌻 Latte</summary>
<img src="assets/latte.webp"/>
</details>
<details>
<summary>🪴 Frappé</summary>
<img src="assets/frappe.webp"/>
</details>
<details>
<summary>🌺 Macchiato</summary>
<img src="assets/macchiato.webp"/>
</details>
<details>
<summary>🌿 Mocha</summary>
<img src="assets/mocha.webp"/>
</details>

## Usage

1. Run `go get github.com/catppuccin/fyne`
2. Add an import of `github.com/catppuccin/fyne` to your code
3. Call the app's `Settings.().SetTheme()` method with an instantiated theme struct:

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

## 💝 Thanks to

- [Michael Baklor](https://github.com/mbaklor)

&nbsp;

<p align="center">
	<img src="https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/footers/gray0_ctp_on_line.svg?sanitize=true" />
</p>

<p align="center">
	Copyright &copy; 2021-present <a href="https://github.com/catppuccin" target="_blank">Catppuccin Org</a>
</p>

<p align="center">
	<a href="https://github.com/catppuccin/catppuccin/blob/main/LICENSE"><img src="https://img.shields.io/static/v1.svg?style=for-the-badge&label=License&message=MIT&logoColor=d9e0ee&colorA=363a4f&colorB=b7bdf8"/></a>
</p>
