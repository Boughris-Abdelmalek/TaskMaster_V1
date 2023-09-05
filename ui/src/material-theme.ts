import { createTheme } from "@mui/material";

declare module "@mui/material/styles" {
  interface Palette {
    tertiary: {
      main: string;
      contrastText: string;
    };
  }

  interface PaletteOptions {
    tertiary?: {
      main: string;
      contrastText: string;
    };
  }
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
  tertiary: {
    main: "#FFFFFF",
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
  fontFamily: "Rosario Variable",
  h1: {
    fontFamily: "Rosario Variable",
    color: palette.text.black,
    fontStyle: "normal",
  },
  h2: {
    fontFamily: "Rosario Variable",
    color: palette.text.black,
    fontStyle: "normal",
  },
  h3: {
    fontFamily: "Rosario Variable",
    color: palette.text.black,
    fontStyle: "normal",
  },
  h4: {
    fontFamily: "Rosario Variable",
    color: palette.text.black,
  },
  h5: {
    fontFamily: "Rosario Variable",
    color: palette.text.black,
  },
  h6: {
    fontFamily: "Rosario Variable",
    color: palette.text.black,
  },
  subtitle1: {},
  subtitle2: {},
  body1: {},
  body2: {},
  caption: {},
  button: {
    fontFamily: "Rosario Variable",
  },
  overline: {},
};

const breakpoints = {
  values: {
    xs: 0,
    sm: 600,
    md: 900,
    lg: 1200,
    xl: 1363,
  },
};

let theme = createTheme({
  palette,
  typography,
  breakpoints,
});

theme = createTheme(theme, {
  components: {
    MuiAppBar: {
      styleOverrides: {
        root: {
          backgroundColor: theme.palette.tertiary.main,
        },
      },
      defaultProps: {
        variant: "elevation",
        elevation: 0,
      },
    },
    MuiToolbar: {
      defaultProps: {
        variant: "dense",
      },
    },
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: 8,
          transition: "unset",
          textTransform: "capitalize",
          fontWeight: "700",
          fontStyle: theme.typography.button,
        },
      },
      defaultProps: {
        size: "small",
        variant: "contained",
        disableElevation: true,
      },
    },
    MuiContainer: {
      styleOverrides: {
        root: {
          maxWidth: "1363px",
        },
      },
    },
  },
});

export default theme;
