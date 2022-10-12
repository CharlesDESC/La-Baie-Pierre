const saisie = document.getElementById("form");
saisie.addEventListener("submit", e => {
    e.preventDefault();
    var name = document.getElementById("name");
    var firstname = document.getElementById("firstname");
    var date = document.getElementById("date");
    var email = document.getElementById("email");
    var pass = document.getElementById("password");
    console.log(name.value)
    console.log(pass.value)
    console.log(email.value)


    fetch('http://localhost:55/api/register', {
        method: 'POST',
        body: JSON.stringify({
            name : name.value,
            firstname : firstname.value,
            birthday : date.value,
            email: email.value,
            password: pass.value

        })
    })
    .then((response) => response.json())
    .then(data => {
        document.cookie = 'name=' + data.name;
        document.cookie = 'id=' + data.ID;
        document.cookie = 'firstname=' + data.firstname;
        location.href = "./index.html";
    })
})
