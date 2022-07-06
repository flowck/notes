import { useEffect } from "react";
import { Route, Routes as RouteList, useLocation, useNavigate } from "react-router-dom";
import { Login } from "../../modules/Auth/Login";
import { My } from "../../modules/My";

export function Routes() {
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const isProtectedRoute = location.pathname.includes("/app");

    if (!isProtectedRoute) {
      return;
    }

    const authToken = localStorage.getItem("notes_user_token");

    if (!authToken) {
      navigate("/login");
    }
  }, [location, navigate]);

  return (
    <RouteList>
      <Route path="/" element={<h1>Home</h1>} />
      <Route path="/app/my" element={<My />} />
      <Route path="/login" element={<Login />} />
    </RouteList>
  );
}
