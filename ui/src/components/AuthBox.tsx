import { Box, Button, Divider, Link, Stack, SvgIcon, TextField, Typography } from "@mui/material";
import Logo from "./Logo.tsx";
import { ReactComponent as GoogleIcon } from "../assets/icons/icon-google.svg";
import { ReactComponent as GithubIcon } from "../assets/icons/icon-github.svg";
import styled from "@emotion/styled";
import theme from "../material-theme.ts";
import { useNavigate } from "react-router-dom";
import React, { FC } from "react";

const SignInBox = styled(Button)`
  color: black;
  background-color: transparent;
  font-size: 1.125rem;
  font-weight: normal;
  width: 100%;
  height: 2.625rem;
  border: 1px solid ${theme.palette.lightGray.main};

  &:hover {
    background-color: ${theme.palette.lightGray.main};
  }
`;

const SignInOption: FC<{ icon: React.ElementType; label: string; redirectUrl: string }> = ({
  icon,
  label,
  redirectUrl,
}) => {
  const navigate = useNavigate();
  const redirect = () => navigate(redirectUrl);

  return (
    <SignInBox
      onClick={redirect}
      startIcon={
        <SvgIcon
          component={icon}
          inheritViewBox
          style={{ width: "22px", height: "22px" }}
          color="action"
        />
      }
    >
      {label}
    </SignInBox>
  );
};

const AuthBox = () => {
  return (
    <Box
      sx={{
        backgroundColor: "white",
        height: "38rem",
        width: "25rem",
        p: 5,
      }}
    >
      <Logo />
      <Typography variant="h5" fontFamily="Rowdies" my={4}>
        Sign in
      </Typography>
      <Stack spacing={4} alignItems="center" my={5}>
        <TextField
          size="small"
          fullWidth
          id="outlined-basic"
          label="email"
          variant="outlined"
          type="email"
        />
        <TextField
          fullWidth
          size="small"
          id="outlined-basic"
          label="password"
          variant="outlined"
          type="password"
        />
        <Button fullWidth size="large">
          Login
        </Button>
        <Divider component="div" role="presentation" sx={{ width: "100%" }} />
        <SignInOption icon={GoogleIcon} label="Continue using Google" redirectUrl="#" />
        <SignInOption icon={GithubIcon} label="Continue using Github" redirectUrl="#" />
      </Stack>
      <Divider />
      <Typography my={2} textAlign="center">
        Don't have an account? <Link href="/signup">Sign up</Link>
      </Typography>
    </Box>
  );
};

export default AuthBox;
