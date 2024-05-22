        fetch("http://localhost:9000/getBooks")
        .then(function(res) {
            return res.json();
        })
        .then(function(books) {
            let placeholder = document.querySelector("#content");
            let out = "";
            for(const element of books.data) {
                out += `
                <tr>
                    <td>${element.book.id}</td>
                    <td>${element.book.title}</td>
                    <td>${element.book.author}</td>
                    <td>${element.book.genre}</td>
                    <td>${element.book.synopsis}</td>
                    <td>${element.book.releaseYear}</td>
                    <td>${element.available}</td>
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