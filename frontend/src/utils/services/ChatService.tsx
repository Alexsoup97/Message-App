import { QueryClient } from "@tanstack/react-query";
import axios from "axios";
import { APIConstants } from "../Constants";

async function getUserMessages() {
  await axios.get(`${APIConstants.BackendUrl}/api/messages/`);
}

export const GetChatNav = () => ({
  queryKey: ["message_list"],
  queryFn: async () => {
    await getUserMessages;
  },
});

export const loadChatData = (queryClient: QueryClient) => {
  return async ({ params }: any) => {
    const query = GetChatNav();
    return (
      queryClient.getQueryData(["message_list"]) ??
      (await queryClient.fetchQuery(query))
    );
  };
};
