import { QueryClient } from "@tanstack/react-query";
import axios from "axios";
import { APIConstants } from "../Constants";

async function getUserMessages() {
  return await axios(`${APIConstants.BackendUrl}/api/messages/conversations`, {
    method: "GET",
    withCredentials: true,
  });
}

export const GetChatNav = () => ({
  queryKey: ["message_list"],
  queryFn: async () => {
    return await getUserMessages();
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
