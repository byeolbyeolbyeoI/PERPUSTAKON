        const form = document.querySelector('.form');
        const error = document.getElementById('error-message');
        const denda = document.getElementById('denda-message');

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
                .then(function(data) {
                    form.reset();                
                    if(data.success == false) {
                        error.style.color = "#cc3333"; 
                    } else {
                        error.style.color = "#4bb544";
                        if(data.denda != 0) {
                            denda.innerHTML = `Dengan denda sebesar ${data.denda}`;
                        }
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