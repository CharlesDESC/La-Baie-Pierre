const response = await fetch('http://localhost:55/api/login', {
method: 'POST',
body: JSON.stringify({
email: "utilisateur7@local.com",
password: "user"
})
})
const json = await response.json()
console.log(json)