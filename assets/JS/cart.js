let allCookies = document.cookie;
allCookies = allCookies.replace("name=", "");
allCookies = allCookies.replace("firstname=", "");
allCookies = allCookies.replace("id=", "");
allCookies = allCookies.replace(";", "");
allCookies = allCookies.split(" ");
console.log("coockie ", allCookies[2]);
let pierre1 = 1;
let pierre2 = 2;
let pierre3 = 3;
let pierre4 = 4;
let pierre5 = 5;
let price = 0;

if (allCookies != "") {
  let member = document.getElementById("member");
  member.textContent = "bonjour " + allCookies[2] + " " + allCookies[0] + ", Voici votre panier :";
  let connection = document.getElementById('connection');
  connection.textContent = "Salut " + allCookies[2] + " ! (clique pour te deconnecter)"
                connection.style.marginLeft = "5px";
                connection.style.textAlign = "right";
                connection.style.cursor = "pointer";

                connection.onclick = function () {
                    document.cookie = "name=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    document.cookie = "id=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    document.cookie = "firstname=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    connection.href = "./index.html"
                }


  fetch("http://localhost:55/api/cart", {
    method: "POST",
    body: JSON.stringify({
      user_ID: parseInt(allCookies[1]),
    }),
  })
    .then((response) => response.json())
    .then((data) => {
      let pouet = [];
      console.log(data);
      if (data != null) {
        for (let i = 0; i < data.length; i++) {
          if (pierre1 == data[i].pierre_ID) {
            price += 5;
          }
          if (pierre2 == data[i].pierre_ID) {
            price += 50;
          }
          if (pierre3 == data[i].pierre_ID) {
            price += 20;
          }
          if (pierre4 == data[i].pierre_ID) {
            price += 60;
          }
          if (pierre5 == data[i].pierre_ID) {
            price += 65;
          }
        }
      }

      let test = document.getElementById("STotal")
      let createBr = document.createElement('br');
      test.textContent = price + " €";
      test.appendChild(createBr)
      test = document.getElementById("Total")
      createBr = document.createElement('br');
      test.textContent = price + " €";
      test.appendChild(createBr)
    });
}
