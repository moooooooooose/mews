import React from "react";
import styled, { keyframes } from "styled-components";
// import logo from './logo.svg';
import Form from "./components/Form.js";
import "./App.css";

function App() {
  return (
    <Wrapper>
      <Form />
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

export default App;
