        const form = document.querySelector('.form');
        const div = document.querySelector('.div-response');

        form.addEventListener('submit', event => {
            event.preventDefault();

            const formData = new FormData(form);
            const title = formData.get('title');
            const author = formData.get('author');
            const genre = formData.get('genre');
            const genreArray = genre.split(",");
            const synopsis = formData.get('synopsis');
            const releaseYear= formData.get('release-year');
            const available = true;

            const releaseYearValue = parseInt(releaseYear);

            fetch('http://localhost:9000/addBook', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json'},
                body: JSON.stringify ({
                    book : {
                        title: title,
                        author: author,
                        genre: genreArray,
                        synopsis: synopsis,
                        releaseYear: releaseYearValue 
                    },
                    available: available
                }),
            })
                .then(res => res.json())
                .then(data => console.log(data))
                .catch(error => console.log(error))
        });