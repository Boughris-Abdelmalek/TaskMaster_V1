import { NavLink } from "react-router-dom";
import styled from "styled-components";
import Logo from "./Logo.tsx";
import { FC } from "react";

type NavBarProps = {
  variant?: "transparent" | "";
};

const Header = styled.header<NavBarProps>`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  height: 5.375rem;
  padding-inline: 1.375rem 8.5rem;
  background-color: ${(props) =>
    props.variant === "transparent" ? "transparent" : props.theme.colors.primary};
`;

const Nav = styled.nav`
  display: flex;
  align-items: center;
  gap: 2rem;
`;

const NavigationLink = styled(NavLink)`
  width: 100%;
  white-space: nowrap;
`;

const Navbar: FC<NavBarProps> = ({ variant = "" }) => {
  return (
    <Header variant={variant}>
      <Logo variant={variant === "transparent" ? "" : "white"} />
      <Nav>
        <NavigationLink to={"/sign-in"}>Login</NavigationLink>
        <NavigationLink to={"/sign-up"}>Sign Up</NavigationLink>
      </Nav>
    </Header>
  );
};

export default Navbar;
