import { Box, Divider, Input, useTheme } from "@mui/joy";
import { ChatPreview } from "./chatPreview";
import { Search, AddCircle } from "@mui/icons-material";
import { useState } from "react";
import { GetChatNav } from "../utils/services/ChatService";
import { useQuery } from "@tanstack/react-query";

export function ChatNavigation() {
  const theme = useTheme();
  const [selected, setSelected] = useState("");

  const { data } = useQuery(GetChatNav());

  const viewMessages = messages.map((message) => (
    <ChatPreview
      key={message.id}
      isSelected={message.id === selected}
      clickHandler={() => setSelected(message.id)}
    />
  ));
  return (
    <Box
      sx={{
        height: "100vh",
        width: "25vw",
        bgcolor: theme.vars.palette.neutral.softBg,
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "columns",
          my: 3,
          gap: 2,
          mx: 3,

          alignContent: "center",
        }}
      >
        <Input startDecorator={<Search />} placeholder="Search" />

        <AddCircle sx={{ position: "relative", top: 5 }} />
      </Box>
      <Divider orientation="horizontal" />
      <Box>{viewMessages}</Box>
    </Box>
  );
}
