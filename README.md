<div align="center">
  <h1> wyd 👀 </h1>
  <em>⚠️W.I.P.⚠️</em>
</div>

  <p align="center">
    <img width="250" src="https://images.metmuseum.org/CRDImages/ad/original/DT4436.jpg" alt="Lover's Eyes" title="Lover's Eyes">
  </p>

<p align="center">
  <strong>A real time website logger. </strong>
 </p>

<p align="center"> Chrome extension + API that displays what website I am currently on. 
</p>

## 📚 Table of Contents

- [Stack](#-stack)
- [How Does It Work](#-how-does-it-work)
- [Todo](#-todo)

## 🥞 Stack

- [Go](https://go.dev)
- Vanilla HTML, CSS and JavaScript

## 🔧 How Does It Work

The chrome extension sends updates to the API with what website I'm currently on. The API is consumed through [nimatullo.com](https://nimatullo.com) where it receives these updates through [server sent events](https://en.wikipedia.org/wiki/Server-sent_events).

## ✅ Todo

- [ ] only allow POST requests from the extension
- [x] use server sent events to update status
- [x] connect database to persist last accessed website
- [x] add toggle for website logging
