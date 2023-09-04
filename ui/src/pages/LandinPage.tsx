import Navbar from "../components/Navbar.tsx";
import heroImage from "../assets/images/hero-page-image.png";
import getStartedImg from "../assets/images/get-starter-image.png";
import smallTaskIcon from "../assets/icons/small-task-icon.svg";
import writeItIcon from "../assets/icons/write-it-icon.svg";
import doItIcon from "../assets/icons/do-it-icon.svg";
import repeatIcon from "../assets/icons/repeat-icon.svg";
import styled from "@emotion/styled";

const Container = styled.main`
  width: 100%;
  max-width: 1728px;
  margin: auto 0;

  header {
    max-width: 70rem;
    margin: 0 auto;
  }
`;

const HeroSection = styled.section`
  min-height: calc(100vh - 5.375rem);
`;

const GetStartedSection = styled.section`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;
  height: 100vh;
  max-width: 67rem;
  margin: auto;
  outline: 1px solid red;

  ul {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    outline: 1px solid red;

    li {
      outline: 1px solid red;
    }
  }
`;

const LandingPage = () => {
  return (
    <Container>
      <Navbar variant={"transparent"} />
      <HeroSection>
        <h3>Organize your day activity with Task Master</h3>
        <button>Get started</button>
        <img src={heroImage} alt="hero-page-image" />
      </HeroSection>
      <GetStartedSection>
        <h3>Don't let your day doing nothing</h3>
        <ul>
          <li>
            <img src={smallTaskIcon} alt="small-task-icon" />
          </li>
          <li>
            <img src={writeItIcon} alt="write-it-icon" />
          </li>
          <li>
            <img src={doItIcon} alt="do-it-icon" />
          </li>
          <li>
            <img src={repeatIcon} alt="repeat-icon" />
          </li>
        </ul>
      </GetStartedSection>
      <section>
        <div>
          <img src={getStartedImg} alt="get-started-image" />
        </div>
      </section>
    </Container>
  );
};

export default LandingPage;
