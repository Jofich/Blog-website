document.addEventListener('DOMContentLoaded', () => {
    
  
    
    const form = document.getElementById('signupForm');

form.addEventListener('submit', async (event) => {
  event.preventDefault();

  const usernameLabel = document.getElementById('username');
  const passwordLabel = document.getElementById('password'); 
  const emailLabel = document.getElementById('email'); 
  console.log(username,password)
  if (!usernameLabel.value || !passwordLabel.value || !emailLabel.value) {
    return;
  }

  try {
    const response = await fetch('/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({username : usernameLabel.value, email: emailLabel.value, password : passwordLabel.value })
    });
    
    if (!response.ok) {
      const Data = await response.json();
      if(response.status === 301){
        newUrl = Data.redirect_url;
        window.location.href = newUrl;
        return;
      }
        
        const err = Data.error;
        if (err == "user with this mail already exists"){
          emailLabel.classList.add("error-input");
          usernameLabel.classList.remove("error-input");
        }
        if(err == "this username is already taken"){
          usernameLabel.classList.add("error-input");
          emailLabel.classList.remove("error-input");
        }
    }
    
  } catch(error){
    console.log(response);
    console.error(error, error.stack);
  }
  });
});