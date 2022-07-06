import { Box, Input, Typography, Alert } from "@mui/material";
import Button from "@mui/lab/LoadingButton";
import { FormEvent, useCallback, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Head } from "../../common/components/Head";
import { http } from "../../common/utils/http";
import { handleApiErrors } from "../../common/utils/api";

interface LoginData {
  email: string;
  password: string;
}

interface LoginSuccess {
  token: string;
}

export function Login() {
  const navigate = useNavigate();
  const [error, setError] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [formData, setFormData] = useState<LoginData>({ email: "", password: "" });

  const onLogin = useCallback(
    (event: FormEvent) => {
      event.preventDefault();

      setError("");

      if (!formData.email || !formData.password) {
        return;
      }

      setIsLoading(true);

      http
        .post<LoginSuccess>("/auth/signin", formData)
        .then(({ data }) => {
          window.localStorage.setItem("notes_user_token", data.token);
          navigate("/app/my");
        })
        .catch((err) => {
          setError(handleApiErrors(err));
        })
        .finally(() => setIsLoading(false));
    },
    [formData, navigate]
  );

  return (
    <Box sx={{ display: "flex", alignItems: "center", justifyContent: "center", height: "100vh" }}>
      <Head title="Login" />

      <Box component="form" onSubmit={onLogin} sx={{ display: "flex", flexDirection: "column", width: "500px" }}>
        <Typography variant="h4" component="h1" sx={{ mb: 2 }}>
          Login
        </Typography>
        <Input
          sx={{ mb: 1 }}
          placeholder="E-mai"
          onInput={({ target }) => setFormData((v) => ({ ...v, email: (target as HTMLInputElement).value }))}
        />
        <Input
          sx={{ mb: 4 }}
          type="password"
          placeholder="****"
          onInput={({ target }) => setFormData((v) => ({ ...v, password: (target as HTMLInputElement).value }))}
        />

        {error && (
          <Alert severity="error" sx={{ mb: 2 }}>
            {error}
          </Alert>
        )}

        <Box sx={{ display: "grid", gridTemplateColumns: "repeat(2, minmax(100px, 1fr))", gridGap: "10px" }}>
          <Button type="submit" loading={isLoading} variant="outlined" fullWidth>
            Login
          </Button>
          <Button variant="outlined" fullWidth>
            Signup
          </Button>
        </Box>
      </Box>
    </Box>
  );
}
