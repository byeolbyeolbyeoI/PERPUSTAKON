        const form = document.querySelector('.form');
        const div = document.querySelector('.div-response');

        form.addEventListener('submit', event => {
            event.preventDefault();

            const formData = new FormData(form);
            const userId = formData.get('user-id');

            const userIdValue = parseInt(userId);

            fetch('http://localhost:9000/returnBook', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json'},
                body: JSON.stringify ({
                    userId : userIdValue,
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