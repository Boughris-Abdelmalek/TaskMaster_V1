import AuthBox from "../components/AuthBox.tsx";
import { Box } from "@mui/material";
import styled from "@emotion/styled";
import theme from "../material-theme.ts";

const Container = styled(Box)`
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: ${theme.palette.secondary.main};
`;

const SignIn = () => {
  return (
    <Container>
      <AuthBox />
    </Container>
  );
};

export default SignIn;
