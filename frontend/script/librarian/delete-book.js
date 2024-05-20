        const form = document.querySelector('.form');
        const error = document.getElementById('error-message');

        form.addEventListener('submit', event => {
            event.preventDefault();

            const formData = new FormData(form);
            const bookId = formData.get('book-id');

            const bookIdValue = parseInt(bookId);

            fetch('http://localhost:9000/deleteBook', {
                method: 'DELETE',
                headers: { 'Content-Type': 'application/json'},
                body: JSON.stringify ({
                    bookId: bookIdValue
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