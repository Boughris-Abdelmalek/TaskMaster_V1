import { createTheme } from "@mui/material";

declare module "@mui/material/styles" {
  interface Theme {}
  // allow configuration using `createTheme`
  interface ThemeOptions {}
}

const palette = {
  primary: {
    main: "#FF4F5A",
    contrastText: "#fff",
  },
  secondary: {
    main: "#F9F9F9",
    contrastText: "#fff",
  },
  text: {
    primary: "#FF4F5A",
    secondary: "#6B6B6B",
    inputPlaceholder: "#ABABAB",
    darkGrayText: "#787878",
    darkText: "#040404",
    black: "#000000",
    white: "#ffffff",
    disabled: "#bda0B2",
  },
};

const typography = {
  fontFamily: "Rowdies",
  h1: {
    fontFamily: "Rowdies",
    color: palette.primary.main,
    fontStyle: "normal",
  },
  h2: {
    fontFamily: "Rowdies",
    color: palette.primary.main,
    fontStyle: "normal",
  },
  h3: {
    fontFamily: "Rowdies",
    color: palette.primary.main,
    fontStyle: "normal",
  },
  h4: {
    fontFamily: "Rowdies",
    color: palette.primary.main,
  },
  h5: {
    fontFamily: "Rowdies",
    color: palette.primary.main,
  },
  h6: {
    fontFamily: "Rowdies",
    color: palette.primary.main,
  },
  subtitle1: {},
  subtitle2: {},
  body1: {},
  body2: {},
  caption: {},
  button: {},
  overline: {},
};

let theme = createTheme({
  palette,
  typography,
});

theme = createTheme(theme, {
  components: {},
});

export default theme;
