import {
  Sheet,
  FormControl,
  Input,
  FormLabel,
  Button,
  Typography,
  Link,
  Box,
} from "@mui/joy";
import { APIConstants } from "../../utils/Constants";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { sheetStyle, inputBox, gridBox } from "./LoginStyles";
import axios from "axios";
import { useAuth } from "../../utils/AuthProvider";

export function SignIn() {
  const [errorMessage, setErrorMessage] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const auth = useAuth();

  const nav = useNavigate();
  async function handleFormSubmit(e: React.ChangeEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.target);

    await axios(APIConstants.BackendUrl + "/api/user/", {
      method: "POST",
      data: formData,
      withCredentials: true,
    })
      .then(() => {
        auth?.Signin("fdsjk");
        nav("/dashboard/");
      })
      .catch(() => {
        setErrorMessage(
          "We didn't recognize the username or password you entered. Please try again.",
        );
      });
  }
  function updateUser(e: React.ChangeEvent<HTMLInputElement>) {
    setUsername(e.target.value);
  }

  function updatePassword(e: React.ChangeEvent<HTMLInputElement>) {
    setPassword(e.target.value);
  }

  return (
    <form onSubmit={handleFormSubmit}>
      <Box sx={gridBox}>
        <Sheet sx={sheetStyle} variant="soft">
          <Typography sx={{ my: 2 }} level="h2">
            Sign In
          </Typography>
          {errorMessage && (
            <Typography sx={{ mx: 2 }} color="danger" level="title-sm">
              {errorMessage}
            </Typography>
          )}
          <FormControl>
            <FormLabel>Username</FormLabel>
            <Input name="username" sx={inputBox} onChange={updateUser}>
              {" "}
            </Input>
          </FormControl>
          <FormControl>
            <FormLabel>Password</FormLabel>
            <Input
              name="password"
              onChange={updatePassword}
              type="password"
              sx={inputBox}
            ></Input>
          </FormControl>
          <Typography level="body-sm">
            Don't have an account?{" "}
            <Link href="signup" level="title-sm">
              Sign up!
            </Link>
          </Typography>
          <Button
            disabled={password === "" || username === ""}
            type="submit"
            variant="outlined"
            sx={{ my: 2, width: 125 }}
          >
            Sign In{" "}
          </Button>
        </Sheet>
      </Box>
    </form>
  );
}
