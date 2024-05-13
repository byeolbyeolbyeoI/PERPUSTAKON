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