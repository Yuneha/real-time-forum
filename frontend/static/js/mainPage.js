import { CreatePost } from "./createPost.js";
import { Post } from "./post.js";
import { Profile } from "./profile.js";
import { FormatDate, TimeAgo } from "./functions.js";
import { Chat } from "./chat.js";
import { Authentification } from "./authentification.js";

let eventListenerSet = false;

export function AddNavBar(user, ws) {
  const navBar = document.getElementById("nav_bar");
  navBar.innerHTML = `<ul>\
      <li class="profile" id="profile">${user.username}</li>
      <li class="dm" id="dm">DM</li>
      <li class="notification" id="notification">Notifs()</li>
      <li class="logout" id="logout">LOGOUT</li>
    </ul>`;

  const logo = document.querySelector(".logo");
  const profile = document.querySelector(".profile");
  const dm = document.querySelector(".dm");
  const logout = document.getElementById("logout");
  const notifications = document.getElementById("notification");

  async function fetchNotifications() {
    try {
      const response = await fetch("/notification", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(user),
      });
      const unreadMessageCount = response.json();

      return unreadMessageCount;
    } catch (error) {
      console.log("Error fetching users:", error);
    }
  }
  fetchNotifications().then((data) => {
    const notif = data;

    if (notif.Count) {
      notifications.innerHTML = `Notifs(${notif.Count})`;
    } else {
      notifications.innerHTML = `Notifs(0)`;
    }
  });

  function toMainPage() {
    MainPage(user, ws);
  }

  function handleLogoClick() {
    toMainPage();
  }

  if (!eventListenerSet) {
    logo.addEventListener("click", handleLogoClick);
    eventListenerSet = true;
  }

  profile.addEventListener("click", () => {
    Profile(user);
  });

  dm.addEventListener("click", () => {
    Chat(user, ws);
  });

  logout.addEventListener("click", () => {
    eventListenerSet = false;
    fetch("/logout", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(user),
    });
    logo.removeEventListener("click", handleLogoClick);
    ws.close();
    Authentification();
  });
}

export function MainPage(user, ws) {
  fetch("/mainPage", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });

  const notifications = document.getElementById("notification");
  ws.onmessage = function (event) {
    const response = JSON.parse(event.data);

    if (response.Type && response.Type === "notif") {
      notifications.innerHTML = `Notifs(${response.Count})`;
    }
  };

  const content = document.querySelector(".content");
  content.id = "main_page";
  content.innerHTML = `<div class="main_page_content">
                <h2>Recent Posts</h2>
            <div class="post_grid" id="post_grid">
                <div class="create_post" id="create_post">
                  <h3 class="create_post_title">Create a Post</h3>
                  <span class="add_post">+</span>
                </div>
                <template id="postcard_template">
                  <div class="postcard">
                    <div class="postcard_content">
                        <h3 class="title"></h3>
                        <div class="information_post">
                            <span class="username"></span>
                            <span class="date"></span>
                        </div>
                        <div class="post_message">
                            <textarea name="message" class="message" id="message" cols="30" rows="6" disabled></textarea>
                        </div>
                        <span class="posted_ago"></span>
                    </div>
                  </div>
                </template>
            </div>
        </div>`;

  const postGrid = document.getElementById("post_grid");
  const postcardTemplate = document.getElementById("postcard_template").content;

  async function fetchPostcards() {
    try {
      const response = await fetch("/posts", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(user),
      });
      const postcards = await response.json();
      return postcards;
    } catch (error) {
      console.error("Error fetching postcards:", error);
    }
  }
  fetchPostcards().then((data) => {
    const posts = data.posts;
    const users = data.users;

    renderPostcards(posts, users);
  });

  function createPostcard(postcard, user) {
    const postedDate = FormatDate(postcard.publication_date);
    const postcardElement = document.importNode(postcardTemplate, true);
    postcardElement.querySelector(".title").textContent = postcard.title;
    postcardElement.querySelector(".username").textContent = user.username;
    postcardElement.querySelector(".date").textContent =
      postcard.publication_date.split(" ")[0];
    postcardElement.querySelector(".message").textContent = postcard.message;
    postcardElement.querySelector(".posted_ago").textContent =
      "Posted " + TimeAgo(postedDate);
    postcardElement.querySelector(".postcard").id =
      "postcard" + postcard.post_id;
    postGrid.appendChild(postcardElement);
  }

  function renderPostcards(posts, users) {
    postGrid.innerHTML = `<div class="create_post" id="create_post">
                  <h3 class="create_post_title">Create a Post</h3>
                  <span class="add_post">+</span>
                </div>`;
    document.getElementById("create_post").addEventListener("click", () => {
      CreatePost(user, ws);
    });
    if (posts && posts.length > 0) {
      for (let index = posts.length - 1; index >= 0; index--) {
        const postcard = posts[index];
        const postcardUser = users[index];
        createPostcard(postcard, postcardUser);
        const postClick = document.getElementById(
          "postcard" + postcard.post_id
        );
        postClick.addEventListener("click", () => {
          Post(postcard, user);
        });
      }
    }
  }
}
