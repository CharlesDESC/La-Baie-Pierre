fetch('http://localhost:55/api/pierre', {
    method: 'GET'
})
    .then((response) => response.json())
    .then(data => {
        let nb_pierre = [];
        for (let i = 0; i < data.length; i++) {
            nb_pierre.push(1)
            var createDiv = document.createElement('div');
            createDiv.id = i;
            createDiv.className = "grid-item";

            card.appendChild(createDiv)

            let pouet = document.getElementById(i);
            createDiv = document.createElement('div');
            createDiv.className = "cssCard";
            createDiv.id = "cssCard" + i;

            pouet.appendChild(createDiv)

            pouet = document.getElementById('cssCard' + i);
            createDiv = document.createElement('img');
            createDiv.className = "img";
            createDiv.src = "../assets/IMG/cailoux.png";

            pouet.appendChild(createDiv)

            createDiv = document.createElement('div');
            createDiv.className = "block";
            createDiv.id = "block" + i;

            pouet.appendChild(createDiv)

            pouet = document.getElementById('block' + i);
            createDiv = document.createElement('div');
            createDiv.className = "namepriece";
            createDiv.id = "namepriece" + i;

            pouet.appendChild(createDiv)

            pouet = document.getElementById('namepriece' + i);
            createDiv = document.createElement('span');
            createDiv.className = "nameCaillou";
            createDiv.textContent = data[i].pierre_name;
            createDiv.id = "nameCaillou" + i;

            pouet.appendChild(createDiv)

            pouet = document.getElementById('namepriece' + i);
            createDiv = document.createElement('span');
            createDiv.textContent = data[i].pierre_price;

            pouet.appendChild(createDiv)

            pouet = document.getElementById('block' + i);
            createDiv = document.createElement('div');
            createDiv.className = "padding";
            createDiv.id = "padding" + i;
            pouet.appendChild(createDiv)

            pouet = document.getElementById('padding' + i);
            createDiv = document.createElement('div');
            createDiv.className = "mainDiv";
            createDiv.id = "mainDiv" + i;
            pouet.appendChild(createDiv)

            pouet = document.getElementById('mainDiv' + i);
            createDiv = document.createElement('button');
            createDiv.className = "minus";
            createDiv.id = "minus" + i;
            createDiv.textContent = "-";

            pouet.appendChild(createDiv)

            let min = document.getElementById('minus' + i);

            min.onclick = function () {
                if (nb_pierre[i] > 0) {
                    nb_pierre[i] -= 1;
                    let update = document.getElementById("numberPlace" + i)
                    update.textContent = nb_pierre[i];
                    console.log(nb_pierre);
                }
            }


            createDiv = document.createElement('span');
            createDiv.id = "numberPlace" + i;
            createDiv.textContent = nb_pierre[i];

            pouet.appendChild(createDiv)

            createDiv = document.createElement('button');
            createDiv.className = "plus";
            createDiv.id = "plus" + i;
            createDiv.textContent = "+";

            pouet.appendChild(createDiv)


            let plus = document.getElementById('plus' + i);

            plus.onclick = function () {
                if (nb_pierre[i] < 10) {
                    nb_pierre[i] += 1;
                    let update = document.getElementById("numberPlace" + i)
                    update.textContent = nb_pierre[i];
                }
            }


            pouet = document.getElementById('padding' + i);
            createDiv = document.createElement('div');
            createDiv.className = "btn";
            createDiv.id = "btn" + i;

            pouet.appendChild(createDiv)

            pouet = document.getElementById('btn' + i);
            createDiv = document.createElement('button');
            createDiv.id = "submit";







            let connection = document.getElementById('connection');

            if (document.cookie != "") {
                createDiv = document.createElement('button');
                createDiv.textContent = "ajouter au panier";
                createDiv.style.cursor = "pointer";
                createDiv.onclick = function () {
                    for (let z = 0; z < nb_pierre[i]; z++) {
                        fetch('http://localhost:55/api/cart', {
                            method: 'POST',
                            body: JSON.stringify({
                                user_ID: parseInt(allCookies[1]),
                                pierre_ID: data[i].ID
                            })
                        })
                        console.log(z)
                    }
                }
                createDiv.id = "submit";
                pouet.appendChild(createDiv)

                createDiv = document.createElement('button');
                createDiv.textContent = "ajouter et consulter le panier";
                createDiv.onclick = function () {
                    fetch('http://localhost:55/api/cart', {
                        method: 'POST',
                        body: JSON.stringify({
                            user_ID: parseInt(allCookies[1]),
                            pierre_ID: data[i].ID
                        })
                    })
                    window.location = "./panier.html";
                }

                createDiv.id = "submit";
                pouet.appendChild(createDiv)

                connection.textContent = "Salut " + allCookies[2] + " ! (clique pour te deconnecter)"
                connection.style.marginLeft = "5px";
                connection.style.cursor = "pointer";

                connection.onclick = function () {
                    document.cookie = "name=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    document.cookie = "id=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    document.cookie = "firstname=; expires=Mon, 02 Oct 2000 01:00:00 GMT"
                    connection.href = "./index.html"
                }


            } else {
                connection.href = "./login.html";
                pouet.onclick = function () {
                    window.location = "./register.html";
                }
                createDiv.textContent = "crée vous un compte pour ajouter au panier";

                pouet.appendChild(createDiv)
            }
        }
    })

let allCookies = document.cookie;
allCookies = allCookies.replace('name=', '')
allCookies = allCookies.replace('firstname=', '')
allCookies = allCookies.replace('id=', '')
allCookies = allCookies.replace(';', '')
allCookies = allCookies.split(' ')
console.log("cookie ", allCookies);




