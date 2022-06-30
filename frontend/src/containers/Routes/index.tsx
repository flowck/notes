import { Route, Routes as RouteList } from "react-router-dom";
import { Login } from "../../modules/Auth/Login";

export function Routes() {
  return (
    <RouteList>
      <Route path="/" element={<h1>Home</h1>} />
      <Route path="/my" element={<h1>My profile</h1>} />
      <Route path="/login" element={<Login />} />
    </RouteList>
  );
}
