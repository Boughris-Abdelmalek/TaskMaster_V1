import logo from "../assets/icons/logo-144.svg";
import logoWhite from "../assets/icons/logo-white-144.svg";
import styled from "@emotion/styled";
import { FC } from "react";
import { Typography } from "@mui/material";

type LogoVariantProps = {
  variant?: "white" | "";
};

const Container = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 45px;
  width: 17rem;
`;

const ImageLogo = styled.img`
  height: 100%;
  width: 100%;
  flex: 1;
`;

const Logo: FC<LogoVariantProps> = ({ variant = "" }) => {
  return (
    <Container>
      <ImageLogo src={variant === "white" ? logoWhite : logo} alt={"logo"} />
      <Typography variant="h3">Task Master</Typography>
    </Container>
  );
};

export default Logo;
