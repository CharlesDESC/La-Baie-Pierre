fetch('http://localhost:55/api/pierre', {
    method: 'GET'
    })
    .then((response) => response.json())
    .then(data => {
        for (let i = 0; i < data.length; i++) {
            console.log(data[i])
            var createDiv = document.createElement('div');
            createDiv.id = i;
            createDiv.textContent = data[i].pierre_name;

            card.appendChild(createDiv)
        }
    })














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
