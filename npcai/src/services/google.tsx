import React from "react";
import GoogleLogin from "react-google-login";

interface Props {
  onSuccess: (response: any) => void;
  onFailure: (response: any) => void;
}

export const GoogleLoginButton: React.FC<Props> = ({ onSuccess, onFailure }) => {
  return (
    <GoogleLogin
      clientId="YOUR_CLIENT_ID"
      buttonText="Login with Gmail"
      onSuccess={onSuccess}
      onFailure={onFailure}
      cookiePolicy={"single_host_origin"}
    />
  );
};

