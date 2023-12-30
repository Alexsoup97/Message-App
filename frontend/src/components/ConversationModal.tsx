import {
  ModalDialog,
  Modal,
  DialogTitle,
  Stack,
  Input,
  FormLabel,
  Avatar,
  Button,
  ListItem,
  ListItemDecorator,
  ListItemContent,
  Checkbox,
  List,
} from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import { Search } from "@mui/icons-material";
import { APIConstants } from "../utils/Constants";
import axios from "axios";
import { useState } from "react";

export function ConversationModal({ open, setOpen, setNewConvo }: any) {
  const [conversationName, setConversationName] = useState("");
  const [usersToAdd, setUsersToAdd] = useState<string[]>([]);

  async function getUsers() {
    return (
      await axios.get(APIConstants.getAllUsers, { withCredentials: true })
    ).data;
  }

  const { isPending, isError, data } = useQuery({
    queryKey: ["users"],
    queryFn: getUsers,
  });

  if (isPending) {
    return <h1> </h1>;
  }

  if (isError) {
    return <h1> </h1>;
  }

  async function createConversation() {
    const response = await axios(APIConstants.CreateConversation, {
      method: "POST",
      withCredentials: true,
      data: {
        name: conversationName,
        participants: usersToAdd,
      },
    });
    await setNewConvo(response.data.id);
  }

  function addUser(username: string) {
    return (e: any) => {
      if (e.target.checked) {
        setUsersToAdd([...usersToAdd, username]);
      } else {
        setUsersToAdd(usersToAdd.filter((user) => user !== username));
      }
    };
  }

  const users = data ? (
    data.map((name: string) => (
      <ListItem key={name}>
        <ListItemDecorator>
          <Avatar size="sm" />
        </ListItemDecorator>
        <ListItemContent>{name}</ListItemContent>
        <Checkbox checked={usersToAdd.includes(name)} onClick={addUser(name)} />
      </ListItem>
    ))
  ) : (
    <br />
  );

  function closeModal() {
    setUsersToAdd([]);
    setConversationName("");
    setOpen(false);
  }

  return (
    <Modal open={open} onClose={closeModal}>
      <ModalDialog sx={{ overflow: "scroll" }}>
        <DialogTitle>New message</DialogTitle>
        <FormLabel>Group Chat Name</FormLabel>
        <Input
          placeholder="Name"
          onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
            setConversationName(e.target.value);
          }}
        ></Input>

        <FormLabel>To:</FormLabel>
        <Input startDecorator={<Search />} placeholder="Search"></Input>
        <form
          onSubmit={async (event: React.FormEvent<HTMLFormElement>) => {
            event.preventDefault();
            if (usersToAdd.length == 0 || conversationName == "") {
              return;
            }
            await createConversation();
            closeModal();
          }}
        >
          <Stack spacing={2}>
            <List>{users}</List>
            <Button type="submit">Create</Button>
          </Stack>
        </form>
      </ModalDialog>
    </Modal>
  );
}
