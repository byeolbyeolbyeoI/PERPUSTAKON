        fetch("http://localhost:9000/getUsers")
        .then(function(res) {
            return res.json();
        })
        .then(function(users) {
            let placeholder = document.querySelector("#content");
            let out = "";
            for(const element of users.data) {
                out += `
                <tr>
                    <td>${element.id}</td>
                    <td>${element.username}</td>
                    <td>${element.role}</td>
                </tr>
                `;
            }

            placeholder.innerHTML = out;
        });