import React from "react";
import styled, { keyframes } from "styled-components";
import "./App.css";

import { useRoutes, A } from "hookrouter";
import routes from "./Routes";

function App() {
  const routeResult = useRoutes(routes);
  return (
    <Wrapper>
      <A href="/">Home</A>
      <A href="/login">Log in</A>
      <A href="/logout">Log out</A>
      {routeResult || <Section />}
    </Wrapper>
  );
}

const HeaderKeyFrame = keyframes`
  0% {
    background-position: 0% 50%;
  }
  100% {
    background-position: 100% 50%;
  }
`;

const Wrapper = styled.section`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
  margin: 0 auto;
  background: linear-gradient(270deg, #46e8be, #e8ba46, #e846d4);
  background-size: 600% 600%;
  animation: ${HeaderKeyFrame} 120s ease infinite;
`;

const Section = styled.section`
  background: #eef;
`;

export default App;
