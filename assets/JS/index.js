fetch('http://localhost:55/api/pierre', {
    method: 'GET'
    })
    .then((response) => response.json())
    .then(data => {
        for (let i = 0; i < data.length; i++) {
            var createDiv = document.createElement('div');
            createDiv.id = i;
            createDiv.className = "cssCard";
            createDiv.textContent = data[i].pierre_name;
            createDiv.style.backgroundColor = "grey";
            createDiv.style.margin = "10px";
            createDiv.style.borderRadius = "20px";
            createDiv.style.display =" flex";
            createDiv.style.flexDirection =  "column";
            createDiv.style.alignItems = "center";


            card.appendChild(createDiv)

            let pouet = document.getElementById(i);
            let createPouet = document.createElement('img');
            createPouet.src = "../assets/IMG/cailoux.png";
            createPouet.style.width = "80px";
            createPouet.style.height = "80px";
            pouet.appendChild(createPouet)

            let connection = document.getElementById('connection');

            if (document.cookie != "") {
                createPouet = document.createElement('a');
                createPouet.textContent = "ajouter au panier";
                createPouet.style.cursor = "pointer";
                createPouet.onclick = function() {
                    fetch('http://localhost:55/api/cart', {
                        method: 'POST',
                        body: JSON.stringify({
                            user_ID : parseInt(allCookies[1]),
                            pierre_ID : data[i].ID
                        })
                    })
                }
                a = document.createElement('a');
                a.textContent = "ajouter et consulter le panier";
                a.onclick = function() {
                    fetch('http://localhost:55/api/cart', {
                        method: 'POST',
                        body: JSON.stringify({
                            user_ID : parseInt(allCookies[1]),
                            pierre_ID : data[i].ID
                        })
                    })
                }
                a.href = "./panier.html";
                pouet.appendChild(a)

                connection.textContent = "Salut " + allCookies[2] + " ! (clique pour te deconnecter)"
                connection.style.marginLeft = "5px";
                connection.style.cursor = "pointer";

                connection.onclick = function() {
                    document.cookie = "name=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    document.cookie = "id=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    document.cookie = "firstname=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    connection.href="./index.html"
                }


            } else {
                createPouet = document.createElement('a');
                createPouet.textContent = "crée vous un compte pour ajouter au panier";
                createPouet.href = "./register.html";

                connection.href="./login.html";
            }
            pouet.appendChild(createPouet)
        }

    })

let allCookies = document.cookie;
allCookies = allCookies.replace('name=', '')
allCookies = allCookies.replace('firstname=', '')
allCookies = allCookies.replace('id=', '')
allCookies = allCookies.replace(';', '')
allCookies = allCookies.split(' ')
console.log("coockie ", allCookies);












// const response = await fetch('http://localhost:55/api/login', {
// method: 'POST',
// body: JSON.stringify({
// email: "utilisateur7@local.com",
// password: "user"
// })
// })
// const json = await response.json()
// console.log(json)


// document.body.onload = addElement;

// function addElement () {
//     var newDiv = document.createElement("div");
//     var newContent = document.createTextNode(data.ID);
//     newDiv.appendChild(newContent);
//     var currentDiv = document.getElementById('div1');
//     document.body.insertBefore(newDiv, currentDiv);
// }
