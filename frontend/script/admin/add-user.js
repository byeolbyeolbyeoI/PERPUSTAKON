        const form = document.querySelector('.form');
        const div = document.querySelector('.div-response');

        form.addEventListener('submit', event => {
            event.preventDefault();

            const formData = new FormData(form);
            const username = formData.get('username');
            const password = formData.get('password');
            const role = formData.get('role');

            fetch('http://localhost:9000/addUser', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json'},
                body: JSON.stringify ({
                    username : username,
                    password : password, 
                    role : role
                }),
            })
                .then(res => res.json())
                .then(data => console.log(data))
                .catch(error => console.log(error))
        });