import { Authentification } from "./authentification.js";

export function UserRegister() {
  const registerContent1 = document.getElementById("register_content_1");
  const registerContent2 = document.getElementById("register_content_2");
  const usernameRegister = document.getElementById("username_register");
  const emailRegister = document.getElementById("email_register");
  const passwordRegister = document.getElementById("password_register");
  const confirmPasswordRegister = document.getElementById(
    "confirm_password_register"
  );
  const firstNameRegister = document.getElementById("first_name_register");
  const lastNameRegister = document.getElementById("last_name_register");
  const genderRegister = document.getElementById("gender_register");
  const dateOfBirthRegister = document.getElementById("date_of_birth_register");
  const registerError1 = document.getElementById("register_error_1");
  const registerError2 = document.getElementById("register_error_2");

  const userRegisterObject = {
    username: usernameRegister.value,
    email: emailRegister.value,
    password: passwordRegister.value,
    confirm_password: confirmPasswordRegister.value,
    first_name: firstNameRegister.value,
    last_name: lastNameRegister.value,
    gender: genderRegister.value,
    date_of_birth: dateOfBirthRegister.value,
  };

  document
    .getElementById("register_button_Back")
    .addEventListener("click", () => {
      registerContent1.style.display = "block";
      registerContent2.style.display = "none";
      document.getElementById("register_button").innerHTML = "Next";
      document.getElementById("register_button_Back").style.display = "none";
      registerError2.innerHTML = "";
    });

  if (document.getElementById("register_button").innerHTML === "Next") {
    if (
      usernameRegister.value === "" ||
      emailRegister.value === "" ||
      passwordRegister.value === "" ||
      confirmPasswordRegister.value === ""
    ) {
      registerError1.innerHTML = "One or more fields are empty!";
    } else if (!isValidEmail(emailRegister.value)) {
      registerError1.innerHTML = "The Email is not valid";
    } else {
      if (passwordRegister.value !== confirmPasswordRegister.value) {
        registerError1.innerHTML = "Password and confirm password don't match!";
      } else {
        registerContent1.style.display = "none";
        registerContent2.style.display = "block";
        document.getElementById("register_button").innerHTML = "Register";
        document.getElementById("register_button_Back").style.display = "block";
      }
    }
  } else if (
    document.getElementById("register_button").innerHTML === "Register"
  ) {
    if (
      firstNameRegister.value === "" ||
      lastNameRegister.value === "" ||
      genderRegister.value === "" ||
      dateOfBirthRegister.value === ""
    ) {
      registerError2.innerHTML = "One or more fields are empty!";
    } else {
      try {
        fetch("http://localhost:5656/register", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(userRegisterObject),
        })
          .then((response) => {
            return response.json();
          })
          .then((data) => {
            console.log(data);
            if (data.message) {
              registerError2.innerHTML = "";
              registerError2.innerHTML = data.message;
              console.log(data.message);
            } else {
              registerError2.innerHTML = "";
              Authentification();
            }
          });
      } catch (error) {
        console.error("Error during registration:", error);
        registerError2.innerHTML = "An error occurred during registration.";
      }
    }
  }
}

function isValidEmail(email) {
  const emailPattern = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  return emailPattern.test(email);
}
