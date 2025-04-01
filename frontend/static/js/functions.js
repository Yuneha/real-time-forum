function Throttle(func, delay) {
  var last;
  var timer;
  return function () {
    var context = this;
    var now = +new Date();
    var args = arguments;
    if (last && now < last + delay) {
      clearTimeout(timer);
      timer = setTimeout(function () {
        last = now;
        func.apply(context, args);
      }, delay);
    } else {
      last = now;
      func.apply(context, args);
    }
  };
}

function FormatDate(date) {
  let [datePart, timePart] = date.split(" ");
  let [day, month, year] = datePart.split("/");
  let formattedDate = `${year}-${month}-${day} ${timePart}`;
  return formattedDate;
}

function TimeAgo(date) {
  const now = new Date();
  const postDate = new Date(date);
  const secondsAgo = Math.floor((now - postDate) / 1000);
  let interval = Math.floor(secondsAgo / 31536000);
  if (interval > 1) return interval + " years ago";
  interval = Math.floor(secondsAgo / 2592000);
  if (interval > 1) return interval + " months ago";
  interval = Math.floor(secondsAgo / 86400);
  if (interval > 1) return interval + " days ago";
  interval = Math.floor(secondsAgo / 3600);
  if (interval > 1) return interval + " hours ago";
  interval = Math.floor(secondsAgo / 60);
  if (interval > 1) return interval + " minutes ago";
  return Math.floor(secondsAgo) + " seconds ago";
}

export { FormatDate, TimeAgo, Throttle };
