import { createGlobalStyle } from "styled-components";

export const GlobalStyles = createGlobalStyle`
  @font-face {
    font-family: "Roboto-Light";
    src: url('/assets/fonts/Roboto-Light.ttf');
  }

  @font-face {
    font-family: "Roboto-Regular";
    src: url('/assets/fonts/Roboto-Regular.ttf');
  }

  @font-face {
    font-family: "Roboto-Medium";
    src: url('/assets/fonts/Roboto-Medium.ttf');
  }

  @font-face {
    font-family: "Roboto-Bold";
    src: url('/assets/fonts/Roboto-Bold.ttf');
  }

  body {
    font-size: 16px;
    font-family: 'Roboto-Regular', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  }
`;
