// Reset class 'fault'
function resetFault(el) {

  if (el.classList.contains('fault')) {
    el.classList.remove('fault');
  }
  
}

// Validation sign up by client method
function validateSignUpData(data) {

  result = true;

  // Check user name
  inp = document.getElementById('userInput');
  msg = document.getElementById('userMsg');
  if (data.user == '') {
    inp.classList.add('fault');
    msg.innerText = 'Required.';
    msg.hidden = false;
    result = false;
  } else {
    msg.hidden = true;
  }

  // Check password
  inp = document.getElementById('passInput');
  msg = document.getElementById('passMsg');
  if (data.pass == '') {
    inp.classList.add('fault');
    msg.innerText = 'Required.';
    msg.hidden = false;
    result = false;
  } else if (!data.pass.match(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,}$/)) {
    msg.innerText = 'Must be 8 characters or more, needs at least one number, one lower and one upper case letters.';
    msg.hidden = false;
    result = false;
  } else {
    msg.hidden = true;
  }

  // Check e-mail
  inp = document.getElementById('emailInput');
  msg = document.getElementById('emailMsg');
  if (data.email == '') {
    inp.classList.add('fault');
    msg.innerText = 'Required.';
    msg.hidden = false;
    result = false;
  } else if (!data.email.match(/^[0-9a-z-\.]+\@[0-9a-z-]{2,}\.[a-z]{2,}$/i)) {
    msg.innerText = 'Invalid email address.';
    msg.hidden = false;
    result = false;
  } else {
    msg.hidden = true;
  }

  return result;

}

// Validate and send registration data
function signUp() {

  // 1. Get sign up data
  data = {
    user: document.getElementById('userInput').value,
    pass: document.getElementById('passInput').value,
    email: document.getElementById('emailInput').value,
    phone: document.getElementById('phoneInput').value
  }

  // 2. Validation by client method
  if (!validateSignUpData(data)) {
    return
  }

  // 3. Blocking elements
  document.getElementById('userInput').disabled = true;
  document.getElementById('passInput').disabled = true;
  document.getElementById('emailInput').disabled = true;
  document.getElementById('phoneInput').disabled = true;

  document.getElementById('userMsg').hidden = true;
  document.getElementById('passMsg').hidden = true;
  document.getElementById('emailMsg').hidden = true;
  document.getElementById('phoneMsg').hidden = true;

  el = document.getElementById('signupBtn');
  el.classList.add('busy');
  if (el.classList.contains('button-primary')) {
    el.classList.remove('button-primary');
  }

  // 4. Send data to server
  axios.post('/signup', data)
  .then(response => {

    // 5. Unblocking elements
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

      // 6. Go to activate-link
      window.location = '/www/activate-link.html';

    } else {

      // 7. Show validation messages
      el = document.getElementById('userMsg');
      el.innerText = response.data.UserMsg;
      el.hidden = (response.data.UserMsg == '');

      el = document.getElementById('passMsg');
      el.innerText = response.data.PassMsg;
      el.hidden = (response.data.PassMsg == '');

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

// Resend activation link
function resendLink() {

  // 1. Get link from params
  var params = window
    .location
    .search
    .replace('?','')
    .split('&')
    .reduce(
        function(p,e){
            var a = e.split('=');
            p[ decodeURIComponent(a[0])] = decodeURIComponent(a[1]);
            return p;
        },
        {}
    );
  data = {
    ActivationLink: params['link']
  }

  // 2. Blocking elements
  el = document.getElementById('resendBtn');
  el.classList.add('busy');
  if (el.classList.contains('button-primary')) {
    el.classList.remove('button-primary');
  }

  // 3. Send data to server
  axios.post('/resend-link', data)
  .then(response => {

    // 4. Unblocking elements
    el = document.getElementById('resendBtn');
    el.classList.add('button-primary');
    if (el.classList.contains('busy')) {
      el.classList.remove('busy');
    }

    if (response.data.Ok == true) {

      // 5. Go to activate-link
      window.location = '/www/activate-link.html';

    } else {

      // 6. Show validation messages
      el = document.getElementById('infoMessage');
      el.innerText = response.data.Message;
      
    }
  })
  .catch(error => console.log(error));

}

// Validation sign in by client method
function validateSignInData(data) {

  result = true;

  // Check user name
  inp = document.getElementById('userInput');
  msg = document.getElementById('userMsg');
  if (data.user == '') {
    inp.classList.add('fault');
    msg.innerText = 'Required.';
    msg.hidden = false;
    result = false;
  } else {
    msg.hidden = true;
  }

  // Check password
  inp = document.getElementById('passInput');
  msg = document.getElementById('passMsg');
  if (data.pass == '') {
    inp.classList.add('fault');
    msg.innerText = 'Required.';
    msg.hidden = false;
    result = false;
  } else {
    msg.hidden = true;
  }

  return result;

}

// Validate and send authentification data
function signIn() {

  // 1. Get sign up data
  data = {
    user: document.getElementById('userInput').value,
    pass: document.getElementById('passInput').value
  }

  // 2. Validation by client method
  if (!validateSignInData(data)) {
    return
  }

  // 3. Blocking elements
  document.getElementById('userInput').disabled = true;
  document.getElementById('passInput').disabled = true;

  document.getElementById('userMsg').hidden = true;
  document.getElementById('passMsg').hidden = true;

  el = document.getElementById('signinBtn');
  el.classList.add('busy');
  if (el.classList.contains('button-primary')) {
    el.classList.remove('button-primary');
  }

  // 4. Send data to server
  axios.post('/signin', data)
  .then(response => {

    // 5. Unblocking elements
    document.getElementById('userInput').disabled = false;
    document.getElementById('passInput').disabled = false;
    el = document.getElementById('signinBtn');
    el.classList.add('button-primary');
    if (el.classList.contains('busy')) {
      el.classList.remove('busy');
    }

    if (response.data.Ok == true) {

      // 6. Go to activate-link
      window.location = '/www/account-info.html';

    } else {

      // 7. Show validation messages
      el = document.getElementById('userMsg');
      el.innerText = response.data.UserMsg;
      el.hidden = (response.data.UserMsg == '');

      el = document.getElementById('passMsg');
      el.innerText = response.data.PassMsg;
      el.hidden = (response.data.PassMsg == '');
      
    }
  })
  .catch(error => console.log(error));
  
}