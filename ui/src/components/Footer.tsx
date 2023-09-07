import logo from "../assets/icons/logo-41.svg";
import { Box, Divider, Stack, Typography } from "@mui/material";
import { FC } from "react";
const FooterList: FC<{ title: string; items: string[] }> = ({ title, items }) => (
  <Stack spacing={1}>
    <Typography variant="h5" color="red">
      {title}
    </Typography>
    {items.map((item, index) => (
      <Typography key={index}>{item}</Typography>
    ))}
  </Stack>
);

const Footer = () => {
  const footerList = ["lorem", "ipsum", "sit", "dolor"];
  const address = {
    companyName: "DK Tech Company",
    street: "uma street, lost city, aincard",
  };

  return (
    <Box p="172px">
      <Divider color="red" sx={{ mb: "2rem" }} />
      <Stack maxHeight="258px" direction="row" justifyContent="center" alignItems="flex-start">
        <Stack spacing={5} flexGrow={3} height="100%">
          <Stack spacing={2} direction="row" alignItems="center">
            <img src={logo} alt="logo-icon" height={60} />
            <Typography variant="h6">
              Lets change your habit <br /> join with million people
            </Typography>
          </Stack>
          <Typography variant="h6">
            <b>{address.companyName}</b>
            <br />
            {address.street}
          </Typography>
        </Stack>
        <Stack direction="row" flexGrow={1} justifyContent="space-between">
          <FooterList title="Features" items={footerList} />
          <FooterList title="Pricing" items={footerList} />
        </Stack>
      </Stack>
    </Box>
  );
};

export default Footer;
