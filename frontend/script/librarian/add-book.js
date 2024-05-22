        const form = document.querySelector('.form');
        const error = document.getElementById('error-message');

        form.addEventListener('submit', event => {
            event.preventDefault();
            console.log("here");

            const formData = new FormData(form);
            const title = formData.get('title');
            const author = formData.get('author');
            const genre = formData.get('genre');
            const genreArray = genre.split(",");
            const synopsis = formData.get('synopsis');
            const releaseYear= formData.get('release-year');
            const available = true;

            const releaseYearValue = parseInt(releaseYear);

            console.log("sini");
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
                .then(function(data) {
                    form.reset();                
                    if(data.success == false) {
                        error.style.color = "#cc3333";
                    } else {
                        error.style.color = "#4bb544";
                    }
                    
                    error.innerHTML = `${data.message}`;
                })
                .catch(error => console.log(error))
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