part of 'chat_bloc.dart';

@immutable
sealed class ChatState {}

final class ChatInitial extends ChatState {}

// final class LoadingSignedUpUsers extends ChatState {}
// final class SignedUpUsersLoaded extends ChatState{
//   final List<User> users;
//   SignedUpUsersLoaded(this.users);
// }

final class LoadingCurrentChats extends ChatState{}
final class CurrentChatsLoaded extends ChatState{
  final List<ChatEntity> chats;
  CurrentChatsLoaded(this.chats);
}

final class LoadingChatWithUser extends ChatState{}
final class ChatWithUserLoaded extends ChatState{
  final List<MessageEntity> messages;
  ChatWithUserLoaded(this.messages);
}

final class CreateNewChat extends ChatState{
  final String sellerId;
  CreateNewChat(this.sellerId);
}


// final class MessageSent extends ChatState{
//   // final String confirmation;
//   // MessageSent(this.confirmation);
// }

// final class TypingStatus extends ChatState{
//   // final bool isTyping;
//   // TypingStatus(this.isTyping);
// }
// final class MessageRecieved extends ChatState{
//   // final String message;
//   // MessageSent(this.message);
// }

// final class LoadingUserProfile extends ChatState{}
// final class UserProfileLoaded extends ChatState{
//   // final User user;
//   // UserProfileLoaded(this.user);
// }


final class ChatDeleted extends ChatState{
  final String chatId;
  ChatDeleted(this.chatId);
}


final class ChatError extends ChatState{
  final String error;
  ChatError(this.error);
}
