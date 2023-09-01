import Logo from "../components/Logo.tsx";
import styled from "styled-components";

const Container = styled.div`
  width: 512px;
  height: 533px;
  background-color: ${(props) => props.theme.colors};
`;

const SignIn = () => {
  return (
    <Container>
      <Logo />
      <h3>Sign In</h3>
      <input />
      <input />
      <br />
      <p>Not have an account? Sign Up</p> {/* Fixed the grammar here */}
    </Container>
  );
};

export default SignIn;
