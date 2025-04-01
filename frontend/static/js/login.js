import { MainPage, AddNavBar } from "./mainPage.js";

export function UserLogin() {
  const username_email_login = document.getElementById("username_email_login");
  const password_login = document.getElementById("password_login");
  const loginError = document.getElementById("login_error");
  let logUser;

  if (username_email_login.value.includes("@")) {
    logUser = {
      email: username_email_login.value,
      password: password_login.value,
    };
  } else {
    logUser = {
      username: username_email_login.value,
      password: password_login.value,
    };
  }
  if (username_email_login.value === "" || password_login.value === "") {
    loginError.innerHTML = "One or more field are empty !";
  } else {
    fetch("http://localhost:5656/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(logUser),
    })
      .then((response) => {
        return response.json();
      })
      .then((data) => {
        const ws = new WebSocket("ws://" + document.location.host + "/ws");
        ws.onopen = function () {
          ws.send(JSON.stringify(logUser.username));
        };
        AddNavBar(data, ws);
        MainPage(data, ws);
      })
      .catch((error) => {
        loginError.innerHTML = "wrong Username/Email or Password";
        error.console.error("fetch issue:", error);
      });
  }
}
