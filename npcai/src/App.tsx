import React, { useState } from 'react';
import {GoogleLoginButton} from './services';

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [userName, setUserName] = useState('');

  const onSuccess = (response: any) => {
    setIsLoggedIn(true);
    setUserName(response.profileObj.name);
  };

  const onFailure = (response: any) => {
    setIsLoggedIn(false);
    setUserName('');
  };

  return (
    <div>
      {isLoggedIn ? (
        <p>Welcome, {userName}!</p>
      ) : (
        <GoogleLoginButton onSuccess={onSuccess} onFailure={onFailure} />
      )}
    </div>
  );
};

export default App;
