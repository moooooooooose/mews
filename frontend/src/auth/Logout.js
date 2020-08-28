import React from "react";
import styled from "styled-components";
import { useGoogleLogout } from "react-google-login";

const clientId =
  "500525628386-evab4nci76397rpec2lpfdooq4jd0tue.apps.googleusercontent.com";

function Logout() {
  const onSuccess = res => {
    console.log("success", res);
  };
  const onFailure = res => {
    console.log("fails", res);
  };

  const { signOut } = useGoogleLogout({
    clientId,
    onSuccess,
    onFailure
  });

  return (
    <Wrapper>
      <Button onClick={signOut}>sign out</Button>
    </Wrapper>
  );
}

const Wrapper = styled.section`
  width: 40%;
`;

const Button = styled.button`
  border-radius: 50px;
  padding: 2em;
  background: transparent;
  color: white;
  border: 3px solid white;
  width: 100%;
`;

export default Logout;
