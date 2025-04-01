export function Index() {
  const loginForm = document.querySelector(".login_form");
  const registerForm = document.querySelector(".register_form");
  const switchButton = document.getElementById("switch");
  const switchButton2 = document.getElementById("to_register");

  switchButton.addEventListener("click", toggleForms);
  switchButton2.addEventListener("click", toggleForms);

  function toggleForms() {
    loginForm.classList.toggle("hidden");
    loginForm.classList.toggle("translateA");
    registerForm.classList.toggle("hidden");
    registerForm.classList.toggle("translateB");
    switchButton.classList.toggle("switch_left");
    switchButton.classList.toggle("switch_right");

    if (switchButton.classList.contains("switch_left")) {
      switchButton.innerHTML = "<";
      switchButton2.innerHTML = "To Register";
      loginForm.style.zIndex = 1;
      registerForm.style.zIndex = 0;
    } else {
      switchButton.innerHTML = ">";
      switchButton2.innerHTML = "To Login";
      loginForm.style.zIndex = 0;
      registerForm.style.zIndex = 1;
    }
  }
}
