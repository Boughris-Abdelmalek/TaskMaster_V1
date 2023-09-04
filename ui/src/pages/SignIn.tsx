import Logo from "../components/Logo.tsx";
import styled from "@emotion/styled";

const Container = styled.div`
  width: 512px;
  height: 533px;
  background-color: white;
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
