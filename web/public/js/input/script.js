document.addEventListener('DOMContentLoaded', () => {
    
  
    
    const form = document.getElementById('registrationForm');

form.addEventListener('submit', async (event) => {
  event.preventDefault();

  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value; 
  console.log(username,password)
  if (!username || !password) {
    
    return;
  }

  try {
    const response = await fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ username, password })
    });

    if (!response.ok) {
      const Data = await response.json();
      if(response.status === 301){
        newUrl = Data.redirect_url;
        window.location.href = newUrl;
        return;
      }
        
      const usernameLabel = document.getElementById("username");
      const passwordLabel = document.getElementById("password");
      const errorDiv = document.getElementById("errorMessage");
      errorDiv.style.display = 'inline';
      usernameLabel.classList.add("error-input");
      passwordLabel.classList.add("error-input");
    }
  } catch (error) {
    console.error('Error:', error);
    alert('Произошла ошибка при отправке формы');
  }
  });
});