import logo from "../assets/icons/logo-144.svg";
import logoWhite from "../assets/icons/logo-white-144.svg";
import styled from "styled-components";
import { FC } from "react";

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

const TitleLogo = styled.h1<LogoVariantProps>`
  font-size: 2.25rem;
  width: 100%;
  text-align: center;
  white-space: nowrap;
  color: ${(props) =>
    props.variant === "white" ? props.theme.colors.white : props.theme.colors.primaryTextColor};
`;

const Logo: FC<LogoVariantProps> = ({ variant = "" }) => {
  return (
    <Container>
      <ImageLogo src={variant === "white" ? logoWhite : logo} alt={"logo"} />
      <TitleLogo variant={variant}>Task Master</TitleLogo>
    </Container>
  );
};

export default Logo;
