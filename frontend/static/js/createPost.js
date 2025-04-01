import { MainPage } from "./mainPage.js";

export function CreatePost(user, ws) {
  const content = document.querySelector(".content");
  content.id = "create_post";
  content.innerHTML = `<div class="create_post_content">
    <div class="create_post_container">
      <h1>Create Post</h1>
          <form class="create_post_form">
              <input type="text" name="title" id="title" placeholder="Enter a title">
              <span>select up to 3 categories</span>
              <select>
                <option>test</option>
              </select>
              <textarea name="message" id="message" cols="50" rows="10" placeholder="Enter a message"></textarea>
              <button type="button" id="create_post_submit">submit</button>
          </form>
    </div>
  </div>`;

  const createPostSubmit = document.getElementById("create_post_submit");

  createPostSubmit.addEventListener("click", () => {
    const title = document.getElementById("title");
    const message = document.getElementById("message");
    const createPostObject = {
      title: title.value,
      message: message.value,
    };
    fetch("http://localhost:5656/createPost", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ user: user, post: createPostObject }),
    })
      .then((response) => {
        return response.json();
      })
      .then(() => {
        MainPage(user, ws);
      })
      .catch((error) => {
        console.error("fetch issue:", error);
      });
  });
}
