import logo from "../assets/icons/logo-41.svg";
import logoWhite from "../assets/icons/logo-white-41.svg";
import { FC } from "react";
import { Box, Typography } from "@mui/material";

type LogoVariantProps = {
  variant?: "white" | "";
};

const Logo: FC<LogoVariantProps> = ({ variant = "" }) => {
  return (
    <Box sx={{ display: "flex", flexGrow: 1, columnGap: 2 }}>
      <img src={variant === "white" ? logoWhite : logo} alt="logo" />
      <Typography
        variant="h4"
        fontFamily="Rowdies"
        color={variant === "white" ? "#FFFFFF" : "#FF4F5A"}
      >
        Task Master
      </Typography>
    </Box>
  );
};

export default Logo;
