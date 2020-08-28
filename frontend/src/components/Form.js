import React, { useCallback } from "react";
import styled from "styled-components";

function Form() {
  const handleSubmit = useCallback(() => {});

  return (
    <StyledForm onSubmit={handleSubmit}>
      <StyledInput placeholder="photo url" name="input"></StyledInput>
      <StyledInput placeholder="spreadsheet" name="input"></StyledInput>
      <Button type="submit" value="Mews me">
        mews me
      </Button>
    </StyledForm>
  );
}

const StyledForm = styled.form`
  width: 30%;
`;
const StyledInput = styled.input`
  border-radius: 50px;
  padding: 2em;
  background: transparent;
  color: white;
  border: 3px solid white;
  width: 100%;
  margin-bottom: 1em;
`;
const Button = styled.button`
  border-radius: 50px;
  padding: 2em;
  background: transparent;
  color: white;
  border: 3px solid white;
  width: 100%;
`;

export default Form;
