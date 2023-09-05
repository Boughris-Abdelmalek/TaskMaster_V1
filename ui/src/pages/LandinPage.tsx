import Navbar from "../components/Navbar";
import {
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Container,
  Stack,
  Typography,
} from "@mui/material";
import heroImage from "../assets/images/hero-page-image.png";
import getStartedImg from "../assets/images/get-starter-image.png";
import smallTaskIcon from "../assets/icons/small-task-icon.svg";
import writeItIcon from "../assets/icons/write-it-icon.svg";
import doItIcon from "../assets/icons/do-it-icon.svg";
import repeatIcon from "../assets/icons/repeat-icon.svg";
import { FC } from "react";

const steps = [
  { icon: smallTaskIcon, label: "small task" },
  { icon: writeItIcon, label: "write it" },
  { icon: doItIcon, label: "do it" },
  { icon: repeatIcon, label: "repeat" },
];

const StepItem: FC<{ icon: string; label: string }> = ({ icon, label }) => (
  <Box textAlign="center" width="100%">
    <img src={icon} alt={`${label}-icon`} />
    <Typography variant="h4" mt={5} textTransform="capitalize">
      {label}
    </Typography>
  </Box>
);

const LandingPage = () => {
  return (
    <Container maxWidth={"xl"} disableGutters>
      <Navbar variant="transparent" />
      <Card elevation={0}>
        <CardContent>
          <Typography variant="h3" textAlign="center">
            Organize your day's activities
            <br />
            with Task Master
          </Typography>
        </CardContent>
        <CardActions sx={{ display: "flex", justifyContent: "center" }}>
          <Button size="large" sx={{ fontSize: "1.75rem", mt: 5, px: 6 }}>
            Get started
          </Button>
        </CardActions>
        <CardMedia component="img" alt="hero-image" height="100%" image={heroImage} />
      </Card>
      <Box sx={{ height: "100vh" }}>
        <Stack justifyContent="space-evenly" alignItems="center" height="100%">
          <Typography variant="h3">Don't let your day go to waste</Typography>
          <Stack direction="row" width="100%">
            {steps.map((step) => {
              return <StepItem icon={step.icon} label={step.label} />;
            })}
          </Stack>
        </Stack>
      </Box>
      <Card elevation={0} sx={{ display: "flex", maxWidth: "1031px", m: "auto" }}>
        <CardMedia
          component="img"
          alt="get-started-image"
          image={getStartedImg}
          sx={{ width: "583px" }}
        />
        <Box sx={{ textAlign: "left" }}>
          <CardContent>
            <Typography variant="h3">
              Achieve your target <br /> and won your life
            </Typography>
          </CardContent>
          <CardActions>
            <Button size="large" sx={{ fontSize: "1.75rem", mt: 2, px: 6 }}>
              Get started
            </Button>
          </CardActions>
        </Box>
      </Card>
    </Container>
  );
};

export default LandingPage;
