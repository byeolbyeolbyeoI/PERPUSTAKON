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

    document.addEventListener('DOMContentLoaded', function() {
    const logout = document.getElementById('logout')
    logout.addEventListener('click', function(event) {
        event.preventDefault(); // Prevent the default action
        
        fetch("http://localhost:9000/logoutHandler")

        setTimeout(() => {
            window.location.href = "http://localhost:9000/login";
        }, 3000);
    });
});