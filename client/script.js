console.log("LALALALA");

email = document.getElementById("email");
password = document.getElementById("password");
number = document.getElementById("number");
btn = document.getElementById("button");

btn.addEventListener("click", function () {
  fetch("http://localhost:8080/", {
    method: "POST",
    body: JSON.stringify({
      email: email.value,
      password: password.value,
      number: number.value,
    }),
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  })
    .then((response) => response.json())
    .then((data) => {
      var result = document.getElementById("result");

      var plus = document.createElement("p");
      plus.textContent = "Посещений = " + data.plus;

      var minus = document.createElement("p");
      minus.textContent = "Пропусков = " + data.minus;

      var n = document.createElement("p");
      n.textContent = "Пропусков по уважительной = " + data.n;

      result.appendChild(plus);
      result.appendChild(minus);
      result.appendChild(n);
    });
});
