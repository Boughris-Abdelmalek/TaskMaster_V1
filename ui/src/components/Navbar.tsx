import { NavLink } from "react-router-dom";
import { MdOutlineViewTimeline } from "react-icons/md";
import styled from "styled-components";

const Header = styled.header`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  height: 5.375rem;
  padding-inline: 1.375rem 8.5rem;
  background-color: #ff4f5a;
`;

const Nav = styled.nav`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 2rem;
`;

const NavigationLink = styled(NavLink)``;

const Logo = styled(MdOutlineViewTimeline)`
  width: 41px;
  height: 41px;
`;

const Navbar = () => {
  return (
    <Header>
      <Logo />
      <Nav>
        <NavigationLink to={"/sign-in"}>Login</NavigationLink>
        <NavigationLink to={"sign-up"}>Sign Up</NavigationLink>
      </Nav>
    </Header>
  );
};

export default Navbar;
