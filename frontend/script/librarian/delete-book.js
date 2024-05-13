        const form = document.querySelector('.form');
        const div = document.querySelector('.div-response');

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
                .then(data => function() {
                    if(data.success === false) {
                        console.log(data.message);
                    };
                })
                .catch(error => console.log(error))
        });