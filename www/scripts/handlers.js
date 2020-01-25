// Reset class 'fault'
function resetFault(item) {

  if (item.classList.contains('fault')) {
    item.classList.remove('fault');
  }
  
}

// Reset class 'fault' email and phone together
function resetFaultEmailPhone() {
  resetFault(document.getElementById('emailInput'));
  resetFault(document.getElementById('phoneInput'));
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
  if (data.email == '' && data.phone == '') {
    document.getElementById('emailInput').classList.add('fault');
    document.getElementById('phoneInput').classList.add('fault');
    valid = false;
  }
  if (!valid) {
    return;
  }

  axios.post('/signup', data)
  .then(response => {})
  .catch(error => console.log(error));
  
}