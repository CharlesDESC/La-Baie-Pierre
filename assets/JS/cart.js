let allCookies = document.cookie;
allCookies = allCookies.replace("name=", "");
allCookies = allCookies.replace("firstname=", "");
allCookies = allCookies.replace("id=", "");
allCookies = allCookies.replace(";", "");
allCookies = allCookies.split(" ");
console.log("coockie ", allCookies[2]);

if (allCookies != "") {
  let member = document.getElementById("member");
  member.textContent =
    "bonjour " + allCookies[2] + " " + allCookies[0] + ", Voici votre panier :";

  fetch("http://localhost:55/api/cart", {
    method: "POST",
    body: JSON.stringify({
      user_ID : parseInt(allCookies[1]),
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
    });
}
