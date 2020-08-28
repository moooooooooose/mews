import React from "react";
import styled from "styled-components";
import { useGoogleLogin } from "react-google-login";
import { navigate } from "hookrouter";

const clientId =
  "500525628386-evab4nci76397rpec2lpfdooq4jd0tue.apps.googleusercontent.com";

function Login() {
  const onSuccess = res => {
    console.log("success", res);
    navigate("/logout");
  };
  const onFailure = res => {
    console.log("fails", res);
  };

  const { signIn } = useGoogleLogin({
    onSuccess,
    onFailure,
    clientId,
    isSignedIn: true,
    accessType: "offline"
  });

  return (
    <Wrapper>
      <Button onClick={signIn}>sign in</Button>
    </Wrapper>
  );
}

const Wrapper = styled.section`
  width: 30%;
`;

const Button = styled.button`
  border-radius: 50px;
  padding: 2em;
  background: transparent;
  color: white;
  border: 3px solid white;
  width: 100%;
`;

export default Login;
