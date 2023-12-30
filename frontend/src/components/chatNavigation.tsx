import {
  Box,
  Divider,
  Input,
  useTheme,
  IconButton,
  Typography,
} from "@mui/joy";
import { ChatPreview } from "./chatPreview";
import { Search, AddCircle } from "@mui/icons-material";
import { useState } from "react";
import { GetChatNav } from "../utils/services/ChatService";
import { useQuery } from "@tanstack/react-query";
import { ConversationModal } from "./ConversationModal";

export function ChatNavigation() {
  const theme = useTheme();
  const [selected, setSelected] = useState("");
  const [isModalOpen, setModal] = useState(false);
  const { data, refetch } = useQuery(GetChatNav());

  async function setNewConvo(convoId: string) {
    await refetch();
    setSelected(convoId);
  }

  const viewMessages = data!.data ? (
    data!.data.map((message:any) => (
      <ChatPreview
        key={message.ConversationId}
        previewMessage={message.LastMessage}
        name={message.Name}
        isSelected={message.ConversationId === selected}
        clickHandler={() => setSelected(message.ConversationId)}
      />
    ))
  ) : (
    <br />
  );

  return (
    <Box
      sx={{
        height: "100vh",
        width: "25vw",
        maxWidth:400,
        bgcolor: theme.vars.palette.neutral.softBg,
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "columns",
          my:3,
          mx:3,
          gap: 2,
          alignContent: "center",
          overflowX: "hidden"
        }}
      >
        <Input sx={{minWidth:0.85}} startDecorator={<Search />} placeholder="Search" />

        <IconButton onClick={() => setModal(true)}>
          <AddCircle />
        </IconButton>
      </Box>
      <Divider orientation="horizontal" />
      <Box marginTop={1} marginLeft={1}>
        <Typography
          level="body-xs"
          textTransform="uppercase"
          sx={{ letterSpacing: "0.15rem" }}
        >
          Inbox
        </Typography>
      </Box>
      <Box
        sx={{
          overflowY: "scroll",
          height: 1,
          bgcolor: theme.vars.palette.neutral.softBg,
        }}
      >
        {viewMessages}
      </Box>

      <ConversationModal
        setNewConvo={setNewConvo}
        open={isModalOpen}
        setOpen={setModal}
      />
    </Box>
  );
}
