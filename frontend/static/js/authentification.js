import { UserRegister } from "./register.js";
import { UserLogin } from "./login.js";
import { Index } from "./index.js";

export function Authentification() {
  document.getElementById("nav_bar").innerHTML = "";
  const content = document.querySelector(".content");
  content.id = "authentification";
  content.innerHTML = `<span id="switch" class="switch_left">&lt;</span>

    <form action="" method="post" class="login_form translateA" id="login_form">
        <h1>LOGIN</h1>
        <div id="login_error" class="error_message"></div>
        <input type="text" name="username_email_login" id="username_email_login" placeholder="Email or Username">
        <input type="password" name="password_login" id="password_login" placeholder="Password">
        <button type="button" id="login_button">login</button>
    </form>

    <form action="" method="post" class="register_form hidden" id="register_form">
        <h1>REGISTER</h1>
        <div id="register_content_1">
        <div id="register_error_1" class="error_message"></div>
        <input type="text" name="username_register" id="username_register" placeholder="Enter an Username">
            <input type="email" name="email_register" id="email_register" placeholder="Enter an Email">
            <input type="password" name="password_register" id="password_register" placeholder="Enter a Password">
            <input type="password" name="confirm_password_register" id="confirm_password_register"
                placeholder="Confirm password">
        </div>

        <div id="register_content_2">
        <div id="register_error_2" class="error_message"></div>
            <input type="text" name="first_name_register" id="first_name_register"
                placeholder="Enter your First Name">
            <input type="text" name="last_name_register" id="last_name_register" placeholder="Enter your Last Name">
            <select id="gender_register">
              <option value="H">H</option>
              <option value="F">F</option>
            </select>
            </br>
            BirthDate:<input type="date" name="date_of_birth_register" id="date_of_birth_register">
        </div>
        <div class="buttons">
        <button type="button" id="register_button_Back">Back</button>
        <button type="button" id="register_button">Next</button>
        </div>
    </form>
    
    <span id="to_register" class="switch_left">To register</span>`;

  const registerContent2 = document.getElementById("register_content_2");
  registerContent2.style.display = "none";

  document.getElementById("register_button").addEventListener("click", () => {
    UserRegister();
  });

  document.getElementById("login_button").addEventListener("click", () => {
    UserLogin();
  });

  Index();
}

Authentification();
