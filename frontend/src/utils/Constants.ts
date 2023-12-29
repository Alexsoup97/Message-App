export class APIConstants {
  public static BackendUrl: string = "http://localhost:3000";

  public static getAllUsers: string =
    APIConstants.BackendUrl + "/api/user/users";
  
  public static CreateConversation: string = APIConstants.BackendUrl + "/api/messages/conversations"
}
