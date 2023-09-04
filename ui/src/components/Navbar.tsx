import { NavLink } from "react-router-dom";
import styled from "@emotion/styled";
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
  background-color: ${(props) => (props.variant === "transparent" ? "transparent" : "#FF4F5A")};
`;

const Nav = styled.nav`
  display: flex;
  justify-content: center;
  align-items: center;

  :first-child {
    color: #ff4f5a;
  }
`;

const NavigationLink = styled(NavLink)`
  text-decoration: none;
  color: #000000;
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
