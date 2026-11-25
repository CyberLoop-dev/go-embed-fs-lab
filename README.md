<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:111827,100:0f766e&height=170&section=header&text=Go%20Embed%20FS%20Lab&fontSize=38&fontColor=ffffff" alt="Go Embed FS Lab header"/></p>
<p align="center"><img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=19&duration=2800&pause=900&center=true&vCenter=true&width=820&lines=Go%20Embed%20FS%20Lab%20%E2%80%A2%20Learn%20how%20Go%20embeds%20local%20assets%20into%20a%20program%20and%20how%20the%20io%2Ffs%20helpers%20browse%20those%20embedded%20files." alt="Animated project title"/></p>
<p align="center">
<img src="https://img.shields.io/github/license/CyberLoop-dev/go-embed-fs-lab?style=flat-square" alt="License"/>
<img src="https://img.shields.io/github/languages/top/CyberLoop-dev/go-embed-fs-lab?style=flat-square" alt="Language"/>
<img src="https://img.shields.io/github/stars/CyberLoop-dev/go-embed-fs-lab?style=flat-square" alt="Stars"/>
<img src="https://img.shields.io/github/last-commit/CyberLoop-dev/go-embed-fs-lab?style=flat-square" alt="Last commit"/>
</p>

![Project illustration](assets/mark.svg)

## 🔎 Overview
Learn how Go embeds local assets into a program and how the io/fs helpers browse those embedded files.

## ✨ Features
- 📦 Embeds sample files at compile time with the standard library
- 🗂️ Lists embedded entries with fs.ReadDir and fs.WalkDir
- 🔎 Matches files with fs.Glob and reads a sub-filesystem with fs.Sub
- 🧩 Uses no third-party dependencies

## 🎬 Demo
No hosted demo is published. Run the project locally using the Quick Start command below.

## 🚀 Quick Start
    go run .

## 🧰 Tech Stack
- Go 1.22+
- Go standard library
- embed and io/fs

## 🗂️ Folder Structure
    main.go                 Embed and filesystem demonstrations
    assets/                 Embedded sample files
    assets/docs/            Nested documentation examples
    go.mod                  Module metadata

## 🛣️ Roadmap
- [ ] Add a flag to select a subdirectory
- [ ] Show file sizes and MIME hints

## 🤝 Contributing
Issues and pull requests are welcome. Keep changes focused, document behavior, and add a small reproducible example for new features.

## 📄 License
Released under the MIT License. See LICENSE.

<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:0f766e,100:111827&height=110&section=footer" alt="Animated footer wave"/></p>
