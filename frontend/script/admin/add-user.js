        const form = document.querySelector('.form');
        const errorDiv = document.querySelector('.form__input-error-message');

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
                .then(function(data) {
                    if(data.success == false) {
                        const form = document.getElementById('form');
                        form.reset();
                        const error = document.getElementById('error-message');
                        error.style.color = "#cc3333";
                        errorDiv.innerHTML = `${data.message}`;
                    } else {
                        const form = document.getElementById('form');
                        form.reset();
                        const error = document.getElementById('error-message');
                        error.style.color = "#4bb544";
                        errorDiv.innerHTML = `${data.message}`;
                    }
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