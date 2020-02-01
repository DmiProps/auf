// Reset class 'fault'
function resetFault(item) {

  if (item.classList.contains('fault')) {
    item.classList.remove('fault');
  }
  
}

// Validate and send registration data
function signUp() {

  data = {
    user: document.getElementById('userInput').value,
    pass: document.getElementById('passInput').value,
    email: document.getElementById('emailInput').value,
    phone: document.getElementById('phoneInput').value
  }

  valid = true;
  if (data.user == '') {
    document.getElementById('userInput').classList.add('fault');
    valid = false;
  }
  if (data.pass == '') {
    document.getElementById('passInput').classList.add('fault');
    valid = false;
  }
  if (data.email == '') {
    document.getElementById('emailInput').classList.add('fault');
    valid = false;
  }
  if (!valid) {
    return;
  }

  axios.post('/signup', data)
  .then(response => {
    if (response.data.Ok == true) {
      window.location = '/www/signin.html';
    } else {
      el = document.getElementById('userMsg');
      el.innerText = response.data.UserMsg;
      el.hidden = (response.data.UserMsg == '');
      el = document.getElementById('emailMsg');
      el.innerText = response.data.EmailMsg;
      el.hidden = (response.data.EmailMsg == '');
      el = document.getElementById('phoneMsg');
      el.innerText = response.data.PhoneMsg;
      el.hidden = (response.data.PhoneMsg == '');
    }
  })
  .catch(error => console.log(error));
  
}