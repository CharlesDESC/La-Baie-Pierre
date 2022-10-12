const saisie = document.getElementById("form");
saisie.addEventListener("submit", e => {
    e.preventDefault();
    var email = document.getElementById("email");
    var pass = document.getElementById("password");
    console.log(email.value);
    console.log(pass.value);

    fetch("http://localhost:55/api/login", {
      method: "POST",
      body: JSON.stringify({
        email: email.value,
        password: password.value,
      })
    })
    .then((response) => response.json())
    .then(data => {
      console.log(data.ID);
      if (data.ID != 0) {
        document.cookie = 'name=' + data.name;
        document.cookie = 'id=' + data.ID;
        document.cookie = 'firstname=' + data.firstname;
        location.href = "./index.html";
      }
      else {
        let error = document.createElement('div');
        error.textContent = "mot de passe ou email incorecte";
        pouet.appendChild(error)
      }
    })
})