import { Routes } from "../Routes";
import { BrowserRouter } from "react-router-dom";
import React from "react";
import { GlobalStyles } from "../../common/styles/globalStyles";

import { HelmetProvider } from "react-helmet-async";

function App() {
  return (
    <>
      <GlobalStyles />
      <HelmetProvider>
        <BrowserRouter>
          <Routes></Routes>
        </BrowserRouter>
      </HelmetProvider>
    </>
  );
}

export default App;
