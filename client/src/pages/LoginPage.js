import React from 'react';
import {LoginButton} from 'lib/auth/LoginButton.js';

import './LoginPage.css';

const LoginPage = () => (
  <div className="LoginPage">
    <div className="LoginPage_center">
      <h1>Terse</h1>
      <p>Link shortener</p>
      <LoginButton />
    </div>
  </div>
);

export default LoginPage;
