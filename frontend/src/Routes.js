import React from "react";
import Logout from "./auth/Logout";
import Form from "./components/Form";
import Login from "./auth/Login";

const routes = {
  "/": () => <Form />,
  "/login": () => <Login />,
  "/logout": () => <Logout />
};

export default routes;
