// Reset class 'fault'
function resetFault(el) {

  if (el.classList.contains('fault')) {
    el.classList.remove('fault');
  }
  
}

// Validate and send registration data
function signUp() {

  document.getElementById('userMsg').hidden = true;
  document.getElementById('emailMsg').hidden = true;
  document.getElementById('phoneMsg').hidden = true;

  document.getElementById('userInput').disabled = true;
  document.getElementById('passInput').disabled = true;
  document.getElementById('emailInput').disabled = true;
  document.getElementById('phoneInput').disabled = true;

  el = document.getElementById('signupBtn');
  el.classList.add('busy');
  if (el.classList.contains('button-primary')) {
    el.classList.remove('button-primary');
  }

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
    document.getElementById('userInput').disabled = false;
    document.getElementById('passInput').disabled = false;
    document.getElementById('emailInput').disabled = false;
    document.getElementById('phoneInput').disabled = false;
    el = document.getElementById('signupBtn');
    el.classList.add('button-primary');
    if (el.classList.contains('busy')) {
      el.classList.remove('busy');
    }
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